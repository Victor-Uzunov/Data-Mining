# Decision Tree (ID3) — Entropy, Information Gain, and Pruning

## Problem
Build a decision tree to classify breast cancer recurrence using categorical features (Breast Cancer dataset). The goal is to learn splits that maximize class separation and evaluate with train/test and stratified k-fold CV.

## Model: ID3 Decision Tree
- Split criterion: Information Gain (IG) based on Entropy
- Entropy of labels y:
  - H(y) = −∑ p_i log2 p_i
- Information Gain for attribute A:
  - IG(A) = H(y) − ∑_v (|y_v|/|y|) · H(y_v), where y_v are labels where A = v
- Tree building:
  - Recursively choose the attribute with highest IG
  - Create children for each attribute value
  - Stop conditions (pure labels, no attributes, depth/min samples/min gain thresholds)

## Pruning (Reduced Error Pruning)
- Hold out a validation set
- Try converting a subtree into a leaf (majority class of samples reaching the node)
- Keep the change if validation accuracy doesn’t degrade (>= old accuracy)
- Helps reduce overfitting

## Configuration Flags (from input)
The script supports pre- and post-processing flags selected via the interactive input:
- Pre-processing flags (depth/regularization):
  - `N`: limit max depth to 10
  - `K`: require at least 5 samples to split (min_samples)
  - `G`: require minimum gain of 0.01 (min_gain)
- Post-processing flags:
  - `E`: enable Reduced Error Pruning (uses a validation split)
- Modes (first token entered):
  - `0`: apply only pre-flags
  - `1`: apply only post-flags
  - `2`: apply both (defaults if flags omitted)

Examples of input at runtime:
- `0 N K` → depth limit and min sample split only
- `1 E` → enable pruning only
- `2` → apply all: N, K, G and E

## Data Handling
- Dataset file: `decision-tree/breast-cancer.data`
- Missing values `?` are imputed using the mode (most frequent value) per column
- Categorical features are treated as discrete splits; no numeric binning is performed

## Evaluation
- Train/Test split: 80/20 with stratification
- Stratified 10-Fold CV: reports per-fold accuracy, mean and std
- Final test accuracy reported after training on the training set (possibly pruned)

## Complexity
- Tree induction: roughly O(n · d · V) to compute gains, where n samples, d attributes, V distinct values per attribute
- Prediction: O(tree_depth) per sample
- Pruning adds additional passes over validation set

## Data Mining Angle
- Decision trees are interpretable models: rules per path
- Entropy/IG relate to impurity reduction; similar to feature importance
- Pruning and regularization flags combat overfitting
- Stratified CV gives robust generalization estimate

## How to Run (Python)
Use the top-level Makefile commands:

- Temporary virtual environment:
  - `make py decision-tree`
- Persistent virtual environment:
  - `make venv decision-tree`
  - `make run decision-tree`

At runtime, the script will prompt:
- `Enter configuration:` — type a mode and optional flags (e.g., `2`, or `0 N K`, or `1 E`)

## Exam Tips
- Derive entropy and information gain; explain why IG chooses splits
- Discuss stopping criteria and their regularization effect (N/K/G)
- Explain Reduced Error Pruning and why validation accuracy guides pruning decisions
- Compare decision trees to naive Bayes/KNN: interpretability vs bias/variance, need for pruning/validation
