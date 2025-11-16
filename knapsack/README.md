# Knapsack Problem Solver

## Problem Description

The Knapsack Problem is a classic optimization problem where you have a knapsack with a limited weight capacity and a set of items, each with a weight and value. The goal is to determine the most valuable combination of items to include in the knapsack without exceeding the weight limit.

### Problem Variants

This implementation solves the **0/1 Knapsack Problem**, where each item can either be included (1) or excluded (0) from the knapsack - no fractional quantities are allowed.

### Example

**Input:**
- Knapsack capacity: 50
- Items: 
  - Item 1: Weight = 10, Value = 60
  - Item 2: Weight = 20, Value = 100  
  - Item 3: Weight = 30, Value = 120

**Optimal Solution:**
- Include Item 2 and Item 3
- Total weight: 50
- Total value: 220

## Algorithm Explanation

### Genetic Algorithm Approach

This implementation uses a **Genetic Algorithm (GA)**, a metaheuristic inspired by the process of natural selection. Genetic algorithms are particularly effective for complex optimization problems where finding the exact optimal solution is computationally expensive.

#### Core Components

1. **Chromosome Representation**
   - Each solution is represented as a binary array (chromosome)
   - `true` at position i means item i is included in the knapsack
   - `false` at position i means item i is excluded

2. **Population**
   - Collection of candidate solutions (chromosomes)
   - Size: 250 individuals
   - Initialized randomly while respecting weight constraints

3. **Fitness Function**
   - Evaluates the quality of each chromosome
   - Returns total value if weight ≤ capacity, otherwise returns 0
   - Invalid solutions (exceeding capacity) are heavily penalized

#### Genetic Algorithm Parameters

```go
MAXIMUM_GENERATIONS             = 1600   // Total evolution cycles
MUTATION_RATE                   = 0.01   // Probability of gene mutation
TOURNAMENT_SIZE                 = 15     // Competitors in selection
POPULATION_SIZE                 = 250    // Number of individuals
ELITE_COUNT                     = 7      // Best individuals preserved
```

#### Algorithm Steps

1. **Initialization**
   ```
   for each individual in population:
       randomly select items while weight <= capacity
       create binary chromosome representing selection
   ```

2. **Evolution Loop**
   ```
   for generation = 1 to MAXIMUM_GENERATIONS:
       // Selection
       parents = tournament_selection()
       
       // Reproduction
       offspring = crossover(parents)
       
       // Mutation
       mutate(offspring)
       
       // Replacement
       new_population = elitism + offspring
       
       // Track best solution
       update_best_solution()
   ```

3. **Selection - Tournament Selection**
   - Randomly select TOURNAMENT_SIZE individuals
   - Choose the one with highest fitness
   - Provides good selection pressure while maintaining diversity

4. **Crossover - Two-Point Crossover**
   - Select two random crossover points
   - Exchange genetic material between parents
   - Creates two offspring from two parents
   ```
   Parent1: [1,0,1,|1,0,1|,0,1]
   Parent2: [0,1,0,|0,1,0|,1,0]
   Child1:  [1,0,1,|0,1,0|,0,1]
   Child2:  [0,1,0,|1,0,1|,1,0]
   ```

5. **Mutation**
   - Each gene has MUTATION_RATE probability of flipping
   - Introduces new genetic material and prevents premature convergence

6. **Elitism**
   - Preserve ELITE_COUNT best individuals
   - Ensures the population doesn't lose its best solutions

#### Advantages of Genetic Algorithm

- **Global Search**: Explores multiple regions of solution space simultaneously
- **No Gradient Required**: Works with discrete, non-differentiable problems
- **Robust**: Handles noisy and complex fitness landscapes
- **Parallelizable**: Population-based approach allows parallel evaluation

#### Complexity Analysis

- **Time Complexity**: O(G × P × N) where:
  - G = number of generations (1600)
  - P = population size (250)  
  - N = number of items
- **Space Complexity**: O(P × N) for storing population
- **Solution Quality**: Near-optimal, not guaranteed optimal

#### Performance Characteristics

- **Small Problems** (N < 50): Very fast, often finds optimal solution
- **Medium Problems** (N < 200): Good solutions in reasonable time
- **Large Problems** (N > 500): Scales well, provides high-quality approximations

## Usage

### Input Format
The program expects the knapsack capacity, number of items, and item specifications:
```
Enter knapsack's capacity: 
50
Enter number of items:
3
10 60
20 100
30 120
```

### Output Format
The algorithm prints fitness values at random generations and final results:
```
100.0
150.0
200.0
220.0
Best solution is: 220
Items included in it are: 
Item 2
Item 3
```

## Building and Running

### Go Implementation
```bash
# Build
cd go && go build -o knapsack knapsack.go

# Run interactively
./knapsack

# Run with input file
./knapsack < input.txt
```

### Using Makefile
```bash
# Build the solution
make build

# Run tests (requires fmi-ai-judge)
make test

# Clean build artifacts
make clean
```

## Implementation Notes

- The algorithm uses elitism to preserve the best solutions across generations
- Random generation printing provides insight into convergence behavior
- Tournament selection balances exploration and exploitation effectively
- Two-point crossover preserves building blocks better than single-point
- Constraint handling: invalid solutions (exceeding capacity) get zero fitness

## Alternative Approaches

For comparison, other common approaches to the knapsack problem include:

- **Dynamic Programming**: O(nW) time, guarantees optimal solution
- **Greedy Algorithm**: O(n log n) time, approximation algorithm
- **Branch and Bound**: Exact algorithm, exponential worst case
- **Simulated Annealing**: Alternative metaheuristic approach

## References

- Goldberg, D. E. (1989). "Genetic Algorithms in Search, Optimization, and Machine Learning"
- Kellerer, H., Pferschy, U., & Pisinger, D. (2004). "Knapsack Problems"
- Mitchell, M. (1996). "An Introduction to Genetic Algorithms"
