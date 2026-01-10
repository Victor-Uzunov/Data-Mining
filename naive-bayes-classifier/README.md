# Naive Bayes Classifier — Probabilistic Modeling and Independence Assumption

## Problem
Classify categorical records (e.g., House Votes 84 dataset) into classes (e.g., Democrat vs Republican) using feature likelihoods and class priors.

## Model
Naive Bayes assumes conditional independence of features given the class.
- Goal: predict class y ∈ {C1, C2, ...} for input x = (x1, x2, ..., xd)
- Decision rule (MAP):
  \[ \hat{y} = \arg\max_y \; P(y) \prod_{j=1}^d P(x_j \mid y) \]
- In log-space (numerically stable):
  \[ \log P(y) + \sum_j \log P(x_j \mid y) \]

## Variants
- Multinomial NB: counts over discrete tokens/features
- Bernoulli NB: binary features (present/absent)
- Gaussian NB: continuous features modeled with normal distributions (mean/variance per class)

Given House Votes features are categorical (y/n/?), Bernoulli or Multinomial NB is appropriate after encoding.

## Estimation (Training)
- Class prior: \( P(y=c) = \frac{\text{count}(y=c)}{N} \)
- Likelihoods: \( P(x_j = v \mid y=c) = \frac{\text{count}(x_j=v, y=c) + \alpha}{\text{count}(y=c) + \alpha K_j} \)
  - \( \alpha \) is Laplace smoothing to avoid zero probabilities
  - \( K_j \) is number of possible values for feature j

## Handling Missing Values
- House Votes has "?" (unknown); options:
  - Treat as its own category
  - Impute (e.g., mode per class)
  - Skip in likelihood product for that feature

## Pseudocode (Bernoulli NB)
1) Encode features to {0,1} (e.g., y=1, n=0)
2) For each class c:
   - prior[c] = count(y=c) / N
   - For each feature j:
     - p_j1[c] = (count(x_j=1, y=c) + α) / (count(y=c) + 2α)
3) Predict(x):
   - For each class c: score[c] = log prior[c] + Σ_j [ x_j·log p_j1[c] + (1−x_j)·log(1−p_j1[c]) ]
   - return argmax_c score[c]

## Complexity
- Training: O(N·d) to count
- Prediction: O(d·|C|) per sample
- Memory: O(d·|C|)

## Data Mining Angle
- Probabilistic generative model; strong independence assumption
- Works well with high-dimensional sparse data (text)
- Smoothing and encoding choices affect performance; cross-validation for hyperparameters (α)

## Evaluation
- Train/test split or Stratified k-fold CV (preserve class ratios)
- Metrics: accuracy, precision/recall/F1, confusion matrix

## How to Run (Python)
- Temporary venv:
  - make py naive-bayes-classifier
- Persistent venv:
  - make venv naive-bayes-classifier
  - make run naive-bayes-classifier

## Exam Tips
- Derive MAP rule and explain why log-space is used
- Explain Laplace smoothing and independence assumption
- Discuss handling of missing values and choice of NB variant (Bernoulli vs Multinomial)

