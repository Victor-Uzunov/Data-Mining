# Traveling Salesman Problem (TSP) — Search and Optimization

## Problem
Given a set of cities and distances between each pair, find the shortest possible route that visits each city exactly once and returns to the origin city.

## Formulation
- Input: complete weighted graph G=(V,E), weights w(u,v)
- Objective: minimize tour cost ∑ w(pi, p(i+1)) over a permutation p of V (returning to start)
- Constraints: each city visited exactly once, start=end

## Algorithms

### Exact Methods
- Brute Force: try all permutations; time O(n!) — infeasible beyond small n
- Dynamic Programming (Held–Karp): O(n^2·2^n) time, O(n·2^n) space

### Heuristics (Approximate)
- Nearest Neighbor: greedily visit the closest unvisited city (fast, non-optimal)
- 2-Opt / k-Opt: local search that improves tour by swapping edges to reduce cost
- Simulated Annealing / Genetic Algorithms: metaheuristics for large instances

## Implementation Notes (Go)
- Represent cities and distances with adjacency matrix or coordinate distances (Euclidean)
- A simple baseline: Nearest Neighbor to build a tour, then apply 2-Opt to refine

## Pseudocode (Nearest Neighbor + 2-Opt)
- Start from a city; while unvisited exists: pick nearest; append
- Then loop over pairs of edges; if swapping reduces total cost → perform swap; repeat until no improvement

## Complexity
- Nearest Neighbor: O(n^2)
- 2-Opt: O(n^2) per pass; number of passes varies; good practical results

## Data Mining Angle
- Feature engineering on distances (e.g., learned metrics) impacts tour quality
- Heuristics and metaheuristics mirror optimization in clustering and model selection

## How to Run (Go version)
- make run tsp
- make test tsp

## Exam Tips
- Distinguish exact (DP O(n^2·2^n)) vs heuristic methods
- Explain local search (2-Opt) and why it improves tours
- Discuss trade-offs: solution quality vs runtime
