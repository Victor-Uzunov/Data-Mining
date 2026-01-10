# Tic-Tac-Toe — Minimax and Game Search

## Problem
Play optimal Tic-Tac-Toe on a 3×3 board. The objective is to place three of your marks (X or O) in a row, column, or diagonal.

## Game Search Formulation
- Deterministic, perfect-information, zero-sum game
- Players alternate turns; terminal states: win/loss/draw
- Use Minimax to choose optimal moves assuming rational opponent

## Minimax Algorithm
- Value of a node = best achievable outcome assuming both players play optimally
- Max player tries to maximize; Min player tries to minimize
- On terminal states: return utility (+1 win, −1 loss, 0 draw)
- Recursively evaluate child states; choose max/min accordingly

### Alpha-Beta Pruning (Optional)
- Maintain bounds α (best for Max so far) and β (best for Min)
- If a branch cannot influence the final decision (β ≤ α), prune it
- Same optimal result as Minimax, fewer node evaluations

## Evaluation Function (for larger games)
- For non-terminal states, you can define heuristics (e.g., count of open 2-in-a-rows)
- Tic-Tac-Toe is small enough to search exhaustively

## Pseudocode
- minimax(state, is_max):
  - if terminal: return utility
  - if is_max: return max(minimax(child, False) for child in successors)
  - else: return min(minimax(child, True) for child in successors)

## Complexity
- Branching factor b ≈ up to 9 moves initially; depth d ≤ 9
- Minimax time O(b^d); alpha-beta reduces effective branching

## Data Mining Angle
- Search with evaluation functions parallels model scoring and selection
- Pruning is analogous to cutting unpromising hypotheses

## How to Run (Python)
- make py tic-tac-toe
- Or with persistent venv: make venv tic-tac-toe && make run tic-tac-toe

## Exam Tips
- Explain zero-sum assumption and rational opponent
- Describe terminal utilities and recursion
- Show how alpha-beta pruning reduces evaluations without altering optimality
