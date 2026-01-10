# Iris Classification — KNN, Preprocessing, and Evaluation

## Problem
Classify Iris flowers (Setosa, Versicolor, Virginica) using the classic Iris dataset. Features include sepal/petal length/width.

## Dataset
- Source: UCI ML Repository (id=53)
- Instances: 150; Classes: 3 (balanced, 50 each)
- Features: 4 continuous
- Typical preprocessing: shuffle, normalization/standardization optional (KNN benefits from scaling)

## Algorithm: k-Nearest Neighbors (KNN)
- Non-parametric, instance-based (lazy learning)
- Decision for a query x:
  1. Compute distances d(x, xi) to all training points (e.g., Euclidean)
  2. Take k smallest distances → neighborhood
  3. Predict majority class (optionally weighted by 1/d)
- Hyperparameter: k (odd values to avoid ties)
- Scaling matters: features with large variance dominate Euclidean distance

## Cross-Validation (Evaluation)
- 10-fold stratified CV used to estimate generalization
- Steps:
  1. Split data into 10 folds preserving class distribution
  2. For each fold: train on 9 folds, test on 1
  3. Aggregate metrics: accuracy mean and std
- Benefits: robust estimate, low variance vs single split

## Metrics
- Accuracy: correct predictions / total
- Confusion matrix: per-class errors; precision/recall (optional)

## Complexity
- Train time: O(1) (lazy); Memory: O(n)
- Predict time per query: O(n · d) to scan all points (n samples, d features)
- With KD-tree/ball tree (not here): average O(log n) for low d

## Data Mining Angle
- Instance-based learning; impact of feature scaling and distance metrics
- Model selection via k tuning; validation via cross-validation
- Bias-variance trade-off: small k → low bias/high variance; large k → high bias/low variance

## How to Run
- Create venv and run:
  - make venv iris
  - make run iris
- Or temporary venv:
  - make py iris

## Exam Tips
- Explain KNN mechanics, the role of k, and distance choice
- Discuss scaling; why standardization improves Euclidean-based methods
- Justify cross-validation and how it estimates performance
