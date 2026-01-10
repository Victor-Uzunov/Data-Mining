import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split, StratifiedKFold

class VotingBayesModel:
    def __init__(self, alpha=1.0):
        self.alpha = alpha
        self.class_log_probs = {}
        self.feature_log_probs = {}
        self.classes = []
        self.feature_list = []

    def train(self, features, targets):
        self.classes = np.unique(targets)
        self.feature_list = features.columns
        n_total = len(targets)

        for c in self.classes:
            c_count = len(targets[targets == c])
            self.class_log_probs[c] = np.log(c_count / n_total)

        for col in self.feature_list:
            self.feature_log_probs[col] = {}
            unique_vals = features[col].unique()
            num_unique = len(unique_vals)

            for c in self.classes:
                self.feature_log_probs[col][c] = {}
                
                subset = features[targets == c]
                counts = subset[col].value_counts()
                
                denom = len(subset) + (self.alpha * num_unique)
                
                for val in unique_vals:
                    count = counts.get(val, 0)
                    prob = (count + self.alpha) / denom
                    self.feature_log_probs[col][c][val] = np.log(prob)

    def classify(self, features):
        predictions = []
        
        for _, row in features.iterrows():
            scores = {c: self.class_log_probs[c] for c in self.classes}
            
            for col in self.feature_list:
                val = row[col]
                
                for c in self.classes:
                    try:
                        log_prob = self.feature_log_probs[col][c][val]
                    except KeyError:
                        log_prob = np.log(self.alpha / 1e6)
                    
                    scores[c] += log_prob
            
            predictions.append(max(scores, key=scores.get))
            
        return np.array(predictions)

def get_processed_data(mode):
    col_names = [
        'party', 'infants', 'water-project', 'budget', 'physician-fee', 'el-salvador',
        'religious-groups', 'anti-satellite', 'nicaraguan-contras', 'mx-missile',
        'immigration', 'synfuels', 'education', 'superfund', 'crime', 'duty-free', 'south-africa'
    ]
    
    df = pd.read_csv('../house-votes-84.data', names=col_names, header=None)
    X = df.drop('party', axis=1)
    y = df['party']

    if mode == 1:
        for col in X.columns:
            for party in y.unique():
                is_party = (y == party)
                is_missing = (X[col] == '?')
                
                valid_votes = X.loc[is_party & ~is_missing, col]
                if not valid_votes.empty:
                    party_mode = valid_votes.mode()[0]
                    X.loc[is_party & is_missing, col] = party_mode
    
    return X, y

def execute_task():
    try:
        mode_input = int(input("Enter 0 to keep '?' as 3rd option or 1 to guess the missing values: "))
    except ValueError:
        print("Invalid input.")
        return

    X, y = get_processed_data(mode_input)

    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=0.20, stratify=y, random_state=42, shuffle=True
    )

    model = VotingBayesModel(alpha=1.0)
    model.train(X_train, y_train)

    train_preds = model.classify(X_train)
    train_acc = (train_preds == y_train).mean() * 100

    skf = StratifiedKFold(n_splits=10, shuffle=True, random_state=42)
    fold_accuracies = []
    
    for train_idx, val_idx in skf.split(X_train, y_train):
        X_f_train, X_f_val = X_train.iloc[train_idx], X_train.iloc[val_idx]
        y_f_train, y_f_val = y_train.iloc[train_idx], y_train.iloc[val_idx]
        
        fold_model = VotingBayesModel(alpha=1.0)
        fold_model.train(X_f_train, y_f_train)
        fold_acc = (fold_model.classify(X_f_val) == y_f_val).mean() * 100
        fold_accuracies.append(fold_acc)

    test_preds = model.classify(X_test)
    test_acc = (test_preds == y_test).mean() * 100

    print(f"\n1. Train Set Accuracy:\n    Accuracy: {train_acc:.2f}%")
    
    print("\n10-Fold Cross-Validation Results:")
    for i, acc in enumerate(fold_accuracies):
        print(f"    Accuracy Fold {i+1}: {acc:.2f}%")
    
    print(f"\n    Average Accuracy: {np.mean(fold_accuracies):.2f}%")
    print(f"    Standard Deviation: {np.std(fold_accuracies):.2f}%")
    
    print(f"\n2. Test Set Accuracy:\n    Accuracy: {test_acc:.2f}%")

if __name__ == "__main__":
    execute_task()