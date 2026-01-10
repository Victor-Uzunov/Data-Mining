# Frog Leap Puzzle — Search Algorithms and State Space Modeling

## Problem
You have N frogs on the left (→) and N frogs on the right (←) separated by one empty lily pad. Frogs can:
- Move forward onto an adjacent empty pad
- Jump over exactly one frog into an empty pad
Goal: swap the positions of the two groups using valid moves.

## State Representation
- A linear array of length 2N + 1, e.g. for N=3: [→,→,→,_,←,←,←]
- Actions: for each frog, check if a 1-step move or 2-step jump is valid
- Terminal state: [←,←,←,_,→,→,→]

## Search Formulations
This is a classic single-agent search problem over a finite state space.

### Breadth-First Search (BFS)
- Explores states in layers (shortest number of moves)
- Optimal wrt number of moves (if all moves cost 1)
- Time/space: O(b^d), where b is branching factor, d is depth of optimal solution
- Good baseline to prove minimal steps

### Depth-First Search (DFS)
- Explores deeply first; may find long non-minimal solution
- Lower memory vs BFS; not guaranteed to find optimal path
- Time: O(b^m), m = maximum depth; Space: O(bm)

### A* (Optional)
- If you define a heuristic like: number of frogs not yet crossing, or distance from goal patterns
- Can reduce search vs BFS while still being optimal (with admissible heuristic)

## Valid Move Rules (Local Constraints)
For index i:
- → frog (left group): can move right if pad i+1 empty; can jump right if frog at i+1 and pad i+2 empty
- ← frog (right group): can move left if pad i−1 empty; can jump left if frog at i−1 and pad i−2 empty

## Pseudocode (BFS)
- Model state as string/list; compute neighbors by applying valid moves
- Push initial state to queue; keep visited set
- While queue not empty: pop; if goal → reconstruct path; else push neighbors not visited

## Complexity Considerations
- State count is finite: each position is a permutation of (N left frogs, N right frogs, 1 empty). Total states: C(2N+1, N, N, 1) = multinomial.
- BFS can be costly for larger N due to memory; iterative deepening or A* helps.

## Data Mining Angle
- This task demonstrates search strategies and state modeling, foundational for feature search, rule discovery, and heuristic design
- A* heuristic design mirrors evaluation functions in model selection

## How to Run (Go version)
- Build & run with the main Makefile:
  - make run frog-leap-puzzle
  - make test frog-leap-puzzle

## Exam Tips
- Clearly define state, actions, goal
- Explain BFS vs DFS trade-offs (optimality vs memory)
- If asked about optimal solutions: BFS with unit costs
- If asked about heuristics: design an admissible one and explain why it never overestimates
