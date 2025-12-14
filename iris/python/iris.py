#!/usr/bin/env python3

import random
import numpy as np
import pandas as pd
from ucimlrepo import fetch_ucirepo
from collections import Counter

SEED = 43
np.random.seed(SEED)
random.seed(SEED)

def load_iris_data():
    print("Loading Iris dataset from UCIML Repo...")
    iris = fetch_ucirepo(id=53)

    X_raw = iris.data.features.values
    y_raw = iris.data.targets.values.flatten()

    unique_labels = np.unique(y_raw)
    label_map = {label: i for i, label in enumerate(unique_labels)}
    labels = np.array([label_map[label] for label in y_raw])

    print(f"Data loaded. Total records: {len(labels)}\n")
    return X_raw, labels

def stratified_train_test_split(X, y, test_size=0.2):
    train_idx, test_idx = [], []
    unique_classes = np.unique(y)
    
    for cls in unique_classes:
        cls_indices = np.where(y == cls)[0].tolist()
        random.shuffle(cls_indices)
        
        split_count = int(len(cls_indices) * (1 - test_size))
        
        train_idx.extend(cls_indices[:split_count])
        test_idx.extend(cls_indices[split_count:])
    
    random.shuffle(train_idx)
    random.shuffle(test_idx)
    
    return X[train_idx], X[test_idx], y[train_idx], y[test_idx]

def stratified_k_fold_indices(y, k_folds=10):
    fold_indices = [[] for _ in range(k_folds)]
    unique_classes = np.unique(y)
    
    for cls in unique_classes:
        cls_indices = np.where(y == cls)[0].tolist()
        random.shuffle(cls_indices)
        
        for i, idx in enumerate(cls_indices):
            fold_indices[i % k_folds].append(idx)
            
    return fold_indices

def calculate_accuracy(y_true, y_pred):
    return np.mean(y_true == y_pred) * 100.0

class StandardScaler:
    def __init__(self):
        self.mean_ = None
        self.std_ = None

    def fit(self, X):
        self.mean_ = X.mean(axis=0)
        self.std_ = X.std(axis=0)
        self.std_[self.std_ == 0] = 1 
        return self

    def transform(self, X):
        if self.mean_ is None or self.std_ is None:
            raise RuntimeError("Scaler must be fitted before transforming.")
        return (X - self.mean_) / self.std_

    def fit_transform(self, X):
        return self.fit(X).transform(X)

class KNNClassifier:
    def __init__(self, k=3):
        self.k = k
        self.X_train = None
        self.y_train = None

    def fit(self, X, y):
        self.X_train = X
        self.y_train = y

    def predict(self, X_test):
        if self.X_train is None:
            raise RuntimeError("Model must be fitted before predicting.")
            
        diff = X_test[:, np.newaxis, :] - self.X_train[np.newaxis, :, :]
        sq_distances = np.sum(diff**2, axis=2)
        distances = np.sqrt(sq_distances)
        
        k_indices = np.argsort(distances, axis=1)[:, :self.k]
        k_nearest_labels = self.y_train[k_indices]
        
        predictions = [
            Counter(labels).most_common(1)[0][0]
            for labels in k_nearest_labels
        ]
        
        return np.array(predictions)

def run_k_fold_experiment(X, y, k_values, k_folds=10):
    fold_indices = stratified_k_fold_indices(y, k_folds)
    results = {}
    best_k = -1
    best_mean_acc = -1.0
    
    print(f"--- Running {k_folds}-Fold Cross-Validation ---\n")

    for k in k_values:
        print(f"K = {k}:")
        fold_accuracies = []
        
        for i in range(k_folds):
            val_idx = fold_indices[i]
            
            train_idx = [idx for j, indices in enumerate(fold_indices) if i != j for idx in indices]
            
            X_train_fold, y_train_fold = X[train_idx], y[train_idx]
            X_val_fold, y_val_fold = X[val_idx], y[val_idx]
            
            scaler = StandardScaler().fit(X_train_fold)
            X_train_scaled = scaler.transform(X_train_fold)
            X_val_scaled = scaler.transform(X_val_fold)
            
            knn = KNNClassifier(k=k)
            knn.fit(X_train_scaled, y_train_fold)
            preds = knn.predict(X_val_scaled)
            
            acc = calculate_accuracy(y_val_fold, preds)
            fold_accuracies.append(acc)
            print(f"    Accuracy Fold {i + 1}: {acc:.2f}%")
        
        cv_mean = np.mean(fold_accuracies)
        cv_std = np.std(fold_accuracies)
        
        results[k] = {
            'mean_cv_accuracy': cv_mean,
            'std_cv_accuracy': cv_std,
            'fold_accuracies': fold_accuracies
        }
        
        print(f"Mean CV Accuracy: {cv_mean:.2f}% (Std: {cv_std:.2f}%)")
        print("-" * 30)

        if cv_mean > best_mean_acc:
            best_mean_acc = cv_mean
            best_k = k
            
    return results, best_k

def final_evaluation(X_train, X_test, y_train, y_test, best_k):
    print(f"\n*** Final Evaluation with Best K = {best_k} ***")
    
    final_scaler = StandardScaler().fit(X_train)
    X_train_scaled = final_scaler.transform(X_train)
    X_test_scaled = final_scaler.transform(X_test)
    
    knn_final = KNNClassifier(k=best_k)
    knn_final.fit(X_train_scaled, y_train)
    
    test_preds = knn_final.predict(X_test_scaled)
    test_acc = calculate_accuracy(y_test, test_preds)
    
    train_preds = knn_final.predict(X_train_scaled)
    train_acc = calculate_accuracy(y_train, train_preds)
    
    print(f"Training Set Accuracy (Best K): {train_acc:.2f}%")
    print(f"Test Set Accuracy (Best K): {test_acc:.2f}%")
    print("------------------------------------------")
    
    return train_acc, test_acc

if __name__ == '__main__':
    X, y = load_iris_data()

    X_train_full, X_test, y_train_full, y_test = stratified_train_test_split(
        X, y, test_size=0.2
    )
    
    k_values = [1, 2, 3, 5, 7, 9, 11, 15, 20, 25]

    cv_results, best_k_found = run_k_fold_experiment(
        X_train_full, y_train_full, k_values, k_folds=10
    )
    
    final_evaluation(X_train_full, X_test, y_train_full, y_test, best_k_found)