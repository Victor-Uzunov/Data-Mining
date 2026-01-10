# 0/1 Knapsack — Dynamic Programming vs Greedy

## Problem
Given items with value `v_i` and weight `w_i` and capacity `W`, choose a subset to maximize total value subject to total weight ≤ `W`. Each item can be either taken or not (0/1).

## Formulations
- Decision version: Is there a selection with value ≥ `V` under capacity `W`?
- Optimization version: Maximize ∑ `v_i` `x_i` subject to ∑ `w_i` `x_i` ≤ `W`, `x_i` ∈ {0,1}

## Algorithms

### Greedy (Not optimal for 0/1)
- Sort by value/weight ratio `v_i` / `w_i` and take greedily
- Optimal only for fractional knapsack (where items can be split)

### Dynamic Programming (Optimal)
- DP table: `dp[i][c]` = maximum value using first `i` items with capacity `c`
- Recurrence:
  - `dp[i][c]` = `dp[i−1][c]` if `w_i` > `c` (cannot take item `i`)
  - `dp[i][c]` = max(`dp[i−1][c]`, `dp[i−1][c − w_i]` + `v_i`) otherwise
- Initialization: `dp[0][c]` = 0 for all `c`
- Answer: `dp[n][W]`
- Space optimization: 1D `dp[c]` iterating capacity descending

## Pseudocode (1D DP)
```
dp = array of size W+1 with zeros
for each item i:
  for c from W down to w_i:
    dp[c] = max(dp[c], dp[c − w_i] + v_i)
```

## Complexity
- Time: O(nW)
- Space: O(W) with 1D optimization
- Pseudo-polynomial: depends on `W` (capacity), not item values; NP-hard in general

## Data Mining Angle
- Feature selection under budget (capacity) constraints mirrors knapsack
- Regularization and resource allocation problems often reduce to knapsack-like trade-offs

## How to Run (Go version)
```
make run knapsack
make test knapsack
```

## Exam Tips
- Be able to derive the DP recurrence and justify it
- Explain why greedy by ratio fails for 0/1
- Discuss pseudo-polynomial nature and impact of large `W`
