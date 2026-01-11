import sys
import os
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns
from sklearn.metrics import silhouette_score, davies_bouldin_score

class KMeansClustering:
    def __init__(self, k, algorithm='kmeans', max_iters=300, restarts=10):
        self.k = k
        self.algorithm = algorithm
        self.max_iters = max_iters
        self.restarts = restarts
        self.centroids = None
        self.labels = None
        self.wcss = None

    def _initialize_centroids(self, X):
        n_samples, n_features = X.shape

        if self.algorithm == 'kmeans++':
            centroids = [X[np.random.randint(n_samples)]]
            for _ in range(1, self.k):
                dist_sq = np.array([min([np.inner(c-x, c-x) for c in centroids]) for x in X])
                probs = dist_sq / dist_sq.sum()
                cumprobs = probs.cumsum()
                r = np.random.rand()
                for i, p in enumerate(cumprobs):
                    if r < p:
                        centroids.append(X[i])
                        break
            return np.array(centroids)
        else:
            indices = np.random.choice(n_samples, self.k, replace=False)
            return X[indices]

    def _assign_clusters(self, X, centroids):
        distances = np.linalg.norm(X[:, np.newaxis] - centroids, axis=2)
        return np.argmin(distances, axis=1)

    def _update_centroids(self, X, labels):
        centroids = np.zeros((self.k, X.shape[1]))
        for i in range(self.k):
            points = X[labels == i]
            if len(points) > 0:
                centroids[i] = points.mean(axis=0)
        return centroids

    def _calculate_wcss(self, X, centroids, labels):
        wcss = 0
        for i in range(self.k):
            points = X[labels == i]
            if len(points) > 0:
                wcss += np.sum((points - centroids[i])**2)
        return wcss

    def fit(self, X):
        best_wcss = float('inf')
        best_centroids = None
        best_labels = None

        runs = self.restarts if self.algorithm == 'kmeans' else 1

        print(f"Running {self.algorithm} with {runs} restart(s)...")

        for i in range(runs):
            centroids = self._initialize_centroids(X)

            for _ in range(self.max_iters):
                old_centroids = centroids.copy()
                labels = self._assign_clusters(X, centroids)
                centroids = self._update_centroids(X, labels)

                if np.allclose(old_centroids, centroids):
                    break

            current_wcss = self._calculate_wcss(X, centroids, labels)

            if current_wcss < best_wcss:
                best_wcss = current_wcss
                best_centroids = centroids
                best_labels = labels

        self.centroids = best_centroids
        self.labels = best_labels
        self.wcss = best_wcss

def plot_results(X, centroids, labels, title):
    plt.figure(figsize=(10, 7))
    sns.scatterplot(x=X[:, 0], y=X[:, 1], hue=labels, palette='Set1', s=100, legend='full')
    plt.scatter(centroids[:, 0], centroids[:, 1], c='black', s=300, marker='X', label='Centroids', zorder=10)
    plt.title(title)
    plt.legend()
    plt.grid(True, alpha=0.3)
    plt.show()

def main():
    if len(sys.argv) != 5:
        print("Usage: python kmeans_clustering.py <data_file> <algorithm> <metric_id> <num_clusters>")
        sys.exit(1)

    filename = sys.argv[1]
    algo = sys.argv[2].lower()
    metric_id = sys.argv[3]
    k = int(sys.argv[4])

    file_path = os.path.join('..', filename)

    try:
        X = np.loadtxt(file_path)
        print(f"Loaded data shape: {X.shape}")
    except Exception as e:
        print(f"Error loading file '{file_path}': {e}")
        sys.exit(1)

    model = KMeansClustering(k=k, algorithm=algo, restarts=10)
    model.fit(X)

    wcss = model.wcss

    if metric_id == '1':
        metric_name = "Silhouette Score"
        metric_val = silhouette_score(X, model.labels)
        explanation = "(Higher is better)"
    elif metric_id == '2':
        metric_name = "Davies-Bouldin Index"
        metric_val = davies_bouldin_score(X, model.labels)
        explanation = "(Lower is better)"
    else:
        metric_name = "Silhouette Score"
        metric_val = silhouette_score(X, model.labels)
        explanation = "(Higher is better)"

    print("-" * 40)
    print(f"Results for {algo} (k={k})")
    print("-" * 40)
    print(f"1. WCSS: {wcss:.2f}")
    print(f"2. {metric_name}: {metric_val:.4f} {explanation}")
    print("-" * 40)

    plot_title = f"{algo.capitalize()} Clustering (k={k})\nWCSS: {wcss:.0f} | {metric_name}: {metric_val:.3f}"
    plot_results(X, model.centroids, model.labels, plot_title)

if __name__ == "__main__":
    main()