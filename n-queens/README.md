# N-Queens — Backtracking and Constraint Satisfaction

## Problem

Place N queens on an N×N chessboard so that no two queens attack each other (no shared row, column, or diagonal).

## CSP Formulation

- **Variables**: rows (r = 1..N)
- **Domain**: for each row r, choose column c where the queen is placed
- **Constraints**: no two queens share same column; |r1−r2| ≠ |c1−c2| for diagonals

## Algorithm: Backtracking with Pruning

This implementation uses a **backtracking algorithm** with constraint satisfaction principles and pruning for efficiency:

### Core Algorithm Components

1. **Row by Row Queen Placement**
   - Places queens one row at a time
   - For each row, attempts to place the queen in a column that does not lead to a conflict with already placed queens

2. **Conflict Pruning**
   - **Early Pruning**: If a conflict is detected (same column or diagonal), that position is skipped for the current row
   - **Used Sets**: Maintains sets of used columns and diagonals to enable fast conflict checks

3. **Solution Recording**
   - When a queen is successfully placed in all rows (row == N), the current board configuration is saved as a solution

### Pseudocode

```
function dfs(row):
  if row == N:
    record_solution(current_board)
  for c in 0 to N-1:
    if c is not in used_columns and (row-c) is not in diag1 and (row+c) is not in diag2:
      place_queen(row, c)
      mark_used(row, c)
      dfs(row + 1)
      remove_queen(row, c)
      unmark_used(row, c)
```

## Complexity

- **Worst-case**: Exponential, due to the backtracking tree
- **With Pruning**: The search space is dramatically reduced, making the algorithm feasible for larger N. The exact number of solutions for small N is known, which also helps in estimating the algorithm's performance.

## Data Mining Angle

- The CSP formulation mirrors feature assignment under constraints, common in data mining.
- Backtracking serves as a generic search strategy, applicable to various rule-based systems and pattern avoidance problems.

## How to Run (Go version)

To run the N-Queens solver:

```bash
make run n-queens
make test n-queens
```

## Exam Tips

- Clearly state the problem's constraints and how diagonal attacks are checked.
- Explain the pruning structures used (e.g., sets for used columns and diagonals).
- Be prepared to contrast brute-force search with the backtracking approach used in this algorithm.
