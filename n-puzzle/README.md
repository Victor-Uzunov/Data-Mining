# N-Puzzle — Heuristic Search (A*) and State Space

## Problem

Given an N×N sliding puzzle with tiles 1..(N^2−1) and one empty cell, reach the goal arrangement by sliding tiles into the empty space.

## State Representation

- Board as a 2D grid or flattened list of length N^2
- Empty cell index determines legal moves (up/down/left/right)
- Goal test: tiles in sorted order; empty at last cell (by convention)

## Heuristic Search (A*)

- A* chooses next node by f(n) = g(n) + h(n), where:
  - g(n): cost to reach n (number of moves)
  - h(n): heuristic estimate of remaining cost
- With admissible h (never overestimates), A* finds optimal paths

### Common Heuristics

- Misplaced tiles: h = number of tiles not in goal position
- Manhattan distance: sum over tiles of |x−x_goal| + |y−y_goal|
- Manhattan dominates misplaced tiles (more informed) and is admissible

## Algorithm Steps

- Use a priority queue ordered by f = g + h
- Start from initial board; push with g=0, h computed
- Pop best f; if goal, reconstruct path
- Generate neighbors by sliding the empty cell; update g/h; push unseen or better paths
- Maintain visited or use hash of board for closed set

## Complexity

- State space grows rapidly with N; 8-puzzle is solvable/unsolvable depending on parity
- A* time/space: exponential in depth in worst case; memory is major bottleneck

## Data Mining Angle

- Heuristic design is akin to feature engineering/evaluation functions
- Search strategies parallel model selection in large hypothesis spaces

## How to Run (Go version)

- make run n-puzzle
- make test n-puzzle

## Exam Tips

- Explain admissibility and consistency; why Manhattan is admissible
- Discuss solvability (inversions parity for 8-puzzle)
- Describe open/closed lists and f=g+h prioritization
