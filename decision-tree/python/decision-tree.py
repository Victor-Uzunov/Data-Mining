import pandas as pd
import numpy as np
import io
import sys
import os
from collections import Counter
from sklearn.model_selection import StratifiedKFold, train_test_split
from sklearn.metrics import accuracy_score

class Node:
    def __init__(self, attribute=None, label=None, is_leaf=False):
        self.attribute = attribute
        self.children = {}
        self.label = label
        self.is_leaf = is_leaf

class ID3Tree:
    def __init__(self, config):
        self.root = None
        self.config = config
        self.max_depth = 10 if 'N' in config['pre_flags'] else float('inf')
        self.min_samples = 5 if 'K' in config['pre_flags'] else 0
        self.min_gain = 0.01 if 'G' in config['pre_flags'] else 0

    def entropy(self, y):
        counts = np.unique(y, return_counts=True)[1]
        probabilities = counts / counts.sum()
        return -np.sum(probabilities * np.log2(probabilities))

    def info_gain(self, X, y, attribute):
        total_entropy = self.entropy(y)
        values, counts = np.unique(X[attribute], return_counts=True)
        weighted_entropy = 0

        for v, count in zip(values, counts):
            subset_y = y[X[attribute] == v]
            weighted_entropy += (count / len(y)) * self.entropy(subset_y)

        return total_entropy - weighted_entropy

    def fit(self, X, y):
        if 'E' in self.config['post_flags']:
            train_x, val_x, train_y, val_y = train_test_split(
                X, y, test_size=0.2, stratify=y, random_state=42
            )
            self.root = self._build_tree(train_x, train_y, depth=0)
            self._prune_rep(self.root, val_x, val_y)
        else:
            self.root = self._build_tree(X, y, depth=0)

    def _build_tree(self, X, y, depth):
        if len(y) < self.min_samples:
             return Node(label=self._most_common(y), is_leaf=True)

        if len(np.unique(y)) == 1:
            return Node(label=y.iloc[0], is_leaf=True)

        if depth >= self.max_depth:
            return Node(label=self._most_common(y), is_leaf=True)

        if X.shape[1] == 0:
            return Node(label=self._most_common(y), is_leaf=True)

        gains = {col: self.info_gain(X, y, col) for col in X.columns}
        best_attr = max(gains, key=gains.get)

        if gains[best_attr] < self.min_gain:
             return Node(label=self._most_common(y), is_leaf=True)

        node = Node(attribute=best_attr)

        unique_values = X[best_attr].unique()
        for value in unique_values:
            mask = X[best_attr] == value
            child_X = X[mask].drop(columns=[best_attr])
            child_y = y[mask]

            if len(child_y) == 0:
                node.children[value] = Node(label=self._most_common(y), is_leaf=True)
            else:
                node.children[value] = self._build_tree(child_X, child_y, depth + 1)

        return node

    def _most_common(self, y):
        return Counter(y).most_common(1)[0][0]

    def predict(self, X):
        predictions = []
        for _, row in X.iterrows():
            predictions.append(self._predict_row(row, self.root))
        return np.array(predictions)

    def _predict_row(self, row, node):
        if node.is_leaf:
            return node.label

        val = row.get(node.attribute)
        if val in node.children:
            return self._predict_row(row, node.children[val])
        else:
            counts = []
            for child in node.children.values():
                if child.is_leaf: counts.append(child.label)
            if not counts: return "no-recurrence-events"
            return Counter(counts).most_common(1)[0][0]

    def _prune_rep(self, node, val_X, val_y):
        if node.is_leaf:
            return

        for child in node.children.values():
            self._prune_rep(child, val_X, val_y)

        old_preds = self.predict(val_X)
        old_acc = accuracy_score(val_y, old_preds)

        original_attribute = node.attribute
        original_children = node.children
        original_is_leaf = node.is_leaf
        original_label = node.label

        subset_mask = [self._reaches_node(row, self.root, node) for _, row in val_X.iterrows()]
        subset_y = val_y[subset_mask]

        if len(subset_y) == 0:
            return

        majority_val_class = self._most_common(subset_y)

        node.is_leaf = True
        node.label = majority_val_class
        node.attribute = None
        node.children = {}

        new_preds = self.predict(val_X)
        new_acc = accuracy_score(val_y, new_preds)

        if new_acc >= old_acc:
            pass
        else:
            node.is_leaf = original_is_leaf
            node.label = original_label
            node.attribute = original_attribute
            node.children = original_children

    def _reaches_node(self, row, current_node, target_node):
        if current_node == target_node:
            return True
        if current_node.is_leaf:
            return False
        val = row.get(current_node.attribute)
        if val in current_node.children:
            return self._reaches_node(row, current_node.children[val], target_node)
        return False

def load_data():
    columns = [
        "age", "menopause", "tumor-size", "inv-nodes", "node-caps",
        "deg-malig", "breast", "breast-quad", "irradiat", "class"
    ]

    file_path = os.path.join("..", "breast-cancer.data")

    try:
        df = pd.read_csv(file_path, names=columns, header=None, dtype=str)
    except FileNotFoundError:
        print(f"Error: Could not find file at {file_path}")
        return None

    for col in df.columns:
        mode_val = df[df[col] != '?'][col].mode()[0]
        df[col] = df[col].replace('?', mode_val)

    return df

def parse_input(user_input):
    parts = user_input.split()
    mode = parts[0]
    flags = parts[1:] if len(parts) > 1 else []

    config = {
        'pre_flags': [],
        'post_flags': []
    }

    all_pre = ['N', 'K', 'G']
    all_post = ['E']

    if mode in ['0', '2']:
        if any(f in all_pre for f in flags):
            config['pre_flags'] = [f for f in flags if f in all_pre]
        else:
            config['pre_flags'] = all_pre

    if mode in ['1', '2']:
        if any(f in all_post for f in flags):
            config['post_flags'] = [f for f in flags if f in all_post]
        else:
            config['post_flags'] = all_post

    return config

def main():
    user_in = input("Enter configuration: ")
    config = parse_input(user_in)

    df = load_data()
    if df is None: return

    X = df.drop(columns=['class'])
    y = df['class']

    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=0.20, stratify=y, random_state=42
    )

    model = ID3Tree(config)
    model.fit(X_train, y_train)

    train_preds = model.predict(X_train)
    train_acc = accuracy_score(y_train, train_preds)

    print(f"1. Train Set Accuracy:")
    print(f"   Accuracy: {train_acc*100:.2f}%")

    skf = StratifiedKFold(n_splits=10, shuffle=True, random_state=42)
    fold_accuracies = []

    print(f"\n10-Fold Cross-Validation Results:")
    fold = 1
    X_train_cv = X_train.reset_index(drop=True)
    y_train_cv = y_train.reset_index(drop=True)

    for train_index, val_index in skf.split(X_train_cv, y_train_cv):
        X_fold_train, X_fold_val = X_train_cv.iloc[train_index], X_train_cv.iloc[val_index]
        y_fold_train, y_fold_val = y_train_cv.iloc[train_index], y_train_cv.iloc[val_index]

        fold_model = ID3Tree(config)
        fold_model.fit(X_fold_train, y_fold_train)
        fold_preds = fold_model.predict(X_fold_val)
        acc = accuracy_score(y_fold_val, fold_preds)
        fold_accuracies.append(acc)

        print(f"   Accuracy Fold {fold}: {acc*100:.2f}%")
        fold += 1

    avg_acc = np.mean(fold_accuracies)
    std_acc = np.std(fold_accuracies)

    print(f"\n   Average Accuracy: {avg_acc*100:.2f}%")
    print(f"   Standard Deviation: {std_acc*100:.2f}%")

    test_preds = model.predict(X_test)
    test_acc = accuracy_score(y_test, test_preds)

    print(f"\n2. Test Set Accuracy:")
    print(f"   Accuracy: {test_acc*100:.2f}%")

if __name__ == "__main__":
    main()