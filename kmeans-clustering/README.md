# K-Means Clustering — Initialization, Metrics, and Practical Tips

## Problem
Cluster unlabeled points into `k` groups such that points in the same cluster are similar. The algorithm optimizes compactness by minimizing Within-Cluster Sum of Squares (WCSS).

## Algorithms

### K-Means (Lloyd's Algorithm)
- Objective: minimize WCSS = \(\sum_{i=1}^k \sum_{x \in C_i} \|x - \mu_i\|^2\)
- Steps:
  1. Initialize centroids \(\mu_i\)
  2. Assignment: assign each point to nearest centroid (Euclidean distance)
  3. Update: recompute each centroid as the mean of its assigned points
  4. Repeat 2–3 until convergence (centroids stop changing) or max iterations
- Sensitive to initialization → multiple restarts improve robustness

### K-Means++ (Better Initialization)
- Intuition: spread initial centroids apart proportional to squared distance from existing centroids
- Procedure:
  1. Pick 1st centroid uniformly at random
  2. For each next centroid, sample a point with probability proportional to its squared distance to nearest chosen centroid
- Benefits: reduces chance of poor local minima; often fewer iterations to converge

## Metrics
- WCSS (Within-Cluster Sum of Squares): lower is better; measures compactness of clusters
- Silhouette Score (metric_id=1): \(s \in [-1,1]\), higher is better
  - For each point x: \( s = \frac{b-a}{\max(a,b)} \), where a = intra-cluster distance; b = nearest-cluster distance
- Davies–Bouldin Index (metric_id=2): lower is better
  - Average similarity between each cluster and its most similar other cluster using intra-cluster scatter and centroid distances

## Convergence and Initialization Effects
- K-Means may converge to local minima; restarts with different seeds mitigate this
- K-Means++ improves the expected solution quality with minimal overhead
- Convergence criteria: \(\mu\) unchanged (or below tolerance) or max_iters reached

## Complexity
- Per iteration: O(n · k · d) (n samples, k clusters, d dimensions)
- Number of iterations typically small; overall O(n · k · d · t)
- Restarts multiply cost by the number of runs

## Data Mining Angle
- Unsupervised learning; cluster compactness vs separation trade-off
- Initialization, scaling, and k choice critically impact results
- Use metrics (Silhouette/DBI) and WCSS elbow to choose k

## Implementation Details (This Task)
- File: `kmeans-clustering/python/kmeans-clustering.py`
- Algorithms supported: `kmeans` (with restarts) and `kmeans++` (single run)
- Distance: Euclidean; centroids updated as mean of assigned points
- Visualization: Matplotlib/Seaborn (scatter plot + centroid markers)

## How to Run
Requires a data file of 2D points (e.g., `unbalance.txt`) placed at the repo root (referenced as `../unbalance.txt` from the python folder).

- Temporary venv (auto-install, cleans after):
```bash
make run kmeans-clustering unbalance.txt kmeans 1 8
```

- Persistent venv (recommended for repeated runs):
```bash
make venv kmeans-clustering
make run kmeans-clustering unbalance.txt kmeans 1 8
```

Arguments:
- `<data_file>`: path relative to repo root (e.g., `unbalance.txt`)
- `<algorithm>`: `kmeans` or `kmeans++`
- `<metric_id>`: `1` = Silhouette (higher better), `2` = Davies–Bouldin (lower better)
- `<num_clusters>`: integer k (e.g., `8`)

Examples:
```bash
# Standard K-Means with restarts
make run kmeans-clustering unbalance.txt kmeans 1 8

# K-Means++ initialization
make run kmeans-clustering unbalance.txt kmeans++ 2 8
```

## Practical Tips
- Feature scaling: standardize features if dimensions differ in scale
- Random seeds: set numpy RNG for reproducibility if needed
- k selection: try multiple k and compare metrics; elbow plot with WCSS
- Outliers: can distort centroids; consider robust preprocessing

## Exam Tips
- Derive K-Means steps and explain the WCSS objective
- Describe K-Means++ initialization and why it helps
- Explain Silhouette and Davies–Bouldin metrics and interpretation
- Discuss complexity and the impact of restarts and initialization
