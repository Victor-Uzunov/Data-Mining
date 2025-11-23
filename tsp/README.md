# Traveling Salesman Problem (TSP) Solver

## Problem Description

The Traveling Salesman Problem (TSP) is one of the most famous optimization problems in computer science. Given a set of cities and distances between them, the goal is to find the shortest possible route that visits each city exactly once and returns to the starting city.

### Problem Variants
- **Classical TSP**: Symmetric distances (this implementation)
- **Asymmetric TSP**: Different distances for opposite directions
- **Multiple TSP**: Multiple salesmen starting from different cities

### Example

**Input:**
```
dataset_name
4
CityA 0 0
CityB 100 0
CityC 100 100
CityD 0 100
```

**Output:**
```
CityA -> CityB -> CityC -> CityD -> CityA
Total distance: 400.0
```

## Algorithm Explanation

This implementation uses a **Hybrid Genetic Algorithm** combined with **2-opt local search** - a powerful metaheuristic approach that balances global exploration with local optimization for solving large-scale TSP instances.

### Why Genetic Algorithm + 2-opt for TSP?

**Advantages over Exact Algorithms:**
- **Scalability**: Handles hundreds of cities efficiently (exact algorithms become impractical beyond ~20 cities)
- **Time Complexity**: O(generations × population_size × n²) vs O(n! × n) for brute force
- **Quality Solutions**: Finds near-optimal solutions for large instances

**Advantages over Pure Heuristics:**
- **Global Search**: Avoids local optima through population diversity
- **Adaptive**: Evolves better solutions over time
- **Robustness**: Consistent performance across different problem instances

### Hybrid Algorithm Components

#### 1. **Route Representation**
```go
type Route []int  // Permutation of city indices: [0, 3, 1, 2] means visit cities in that order
```

#### 2. **Genetic Algorithm Core**

**Population Initialization**: Creates diverse random permutations
```go
func initPopulation(popSize, n int) [][]int
```

**Selection**: Tournament selection (k=5) for parent selection
```go
func tournamentSelection(scored []Scored, k int) [][]int
```

**Crossover**: Order Crossover (OX) - preserves city visiting order
```go
func orderCrossover(p1, p2 []int) []int
```

**Mutation**: Random swap mutation (3% rate)
```go
func mutate(route []int, rate float64)
```

#### 3. **2-opt Local Search Enhancement**
```go
func twoOpt(route []int, dist [][]float64) []int
```

The 2-opt algorithm improves routes by:
1. Taking two edges in the route
2. Removing them and reconnecting in the only other possible way
3. Keeping the improvement if it reduces total distance
4. Repeating until no improvements are found

#### 4. **Hybrid Strategy**

- **Small instances (≤20 cities)**: Apply 2-opt to top 10% of population each generation
- **Large instances (>20 cities)**: Apply 2-opt only to the best individual to balance speed vs. quality

### Algorithm Parameters

- **Population Size**: 250 individuals
- **Generations**: 200
- **Tournament Size**: 5
- **Mutation Rate**: 3%
- **Elitism**: 1 (best individual always survives)

## Input Formats

### 1. Named Dataset
```
dataset_name
n
city1 x1 y1
city2 x2 y2
...
cityN xN yN
```

### 2. Random Generation
```
n
```
Generates n random cities with coordinates in [0, 1000] × [0, 1000]

## Usage

### Build and Run
```bash
make build
echo "dataset_name
4
A 0 0
B 100 0
C 100 100
D 0 100" | ./go/tsp
```

### Expected Output
```
<initial_best_distance>
<progress_updates_during_evolution>
<final_best_distance>

A -> B -> C -> D
400.0
```

### Test with Judge
```bash
make test
```

## Implementation Details

### Distance Calculation
Uses Euclidean distance: `sqrt((x2-x1)² + (y2-y1)²)`

### Evolution Process
1. Initialize random population
2. For each generation:
   - Select parents via tournament selection
   - Create children via order crossover and mutation
   - Apply 2-opt local search to best individuals
   - Keep elite individuals for next generation
3. Return best solution found

### Performance Characteristics
- **Time**: ~1-5 seconds for 50 cities, ~10-30 seconds for 100 cities
- **Quality**: Typically within 2-5% of optimal for small instances
- **Memory**: O(population_size × cities) ≈ 250KB for 100 cities

## Algorithm Analysis

### Complexity
- **Time**: O(G × P × N²) where G=generations, P=population, N=cities
- **Space**: O(P × N) for population storage

### When to Use
- **Best for**: Medium to large TSP instances (20-500 cities)
- **Alternative to**: Exact algorithms when optimality isn't required
- **Compared to**: Simulated Annealing (better diversity), Ant Colony (faster convergence)

### Limitations
- Stochastic nature means results may vary between runs
- Not guaranteed to find optimal solution
- Performance depends on parameter tuning
