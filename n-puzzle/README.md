# N-Puzzle Solver

## Problem Description

The N-puzzle is a sliding puzzle that consists of a frame of numbered square tiles in random order with one tile missing. The puzzle is played on a grid that is usually 4×4 (15-puzzle), 3×3 (8-puzzle), or 5×5 (24-puzzle), with the total number of tiles being N.

The objective is to place the tiles in order by making sliding moves that use the empty space. The goal state typically has numbers arranged sequentially from 1 to N, with the empty space in the bottom-right corner.

### Example (3×3 / 8-Puzzle)

**Initial State:**
```
2 8 3
1 6 4
7 _ 5
```

**Goal State:**
```
1 2 3
4 5 6
7 8 _
```

## Algorithm Explanation

### A* Search Algorithm

This implementation uses the **A* (A-star) search algorithm**, which is one of the most effective algorithms for solving the N-puzzle optimally. A* is a best-first search algorithm that finds the least-cost path from a given initial node to a goal node.

#### Core Components

1. **State Representation**
   - Each puzzle state is represented as a 2D grid
   - The empty space is typically represented as 0 or '_'
   - Each state tracks the position of the blank tile for efficient move generation

2. **Search Strategy**
   - **Open Set**: Priority queue of nodes to be evaluated (frontier)
   - **Closed Set**: Set of nodes already evaluated (explored)
   - **Priority Function**: f(n) = g(n) + h(n)
     - g(n): Actual cost from start to node n
     - h(n): Heuristic estimate of cost from n to goal

3. **Heuristic Functions**
   
   The choice of heuristic is crucial for A* performance. Common heuristics include:
   
   **Manhattan Distance (Primary)**
   - Sum of Manhattan distances of all tiles from their goal positions
   - Manhattan distance = |x1 - x2| + |y1 - y2|
   - Admissible (never overestimates) and consistent
   - Example: If tile 5 is at position (1,1) but should be at (1,2), distance = 1
   
   **Linear Conflict (Enhanced)**
   - Manhattan distance + 2 × number of linear conflicts
   - Linear conflict: two tiles in same row/column, in reverse order relative to goal
   - More informed than Manhattan distance alone
   
   **Hamming Distance (Alternative)**
   - Number of tiles in wrong positions
   - Less informed than Manhattan distance

#### Algorithm Steps

1. **Initialization**
   - Add initial state to open set with f(n) = h(n) (g(n) = 0)
   - Initialize closed set as empty

2. **Main Loop**
   ```
   while open set is not empty:
       current = node with lowest f(n) from open set
       
       if current is goal state:
           return reconstruct_path(current)
       
       move current from open set to closed set
       
       for each neighbor of current:
           if neighbor in closed set:
               continue
           
           tentative_g = g(current) + 1
           
           if neighbor not in open set:
               add neighbor to open set
           elif tentative_g >= g(neighbor):
               continue
           
           update neighbor with better path:
               parent[neighbor] = current
               g[neighbor] = tentative_g
               f[neighbor] = g[neighbor] + h(neighbor)
   ```

3. **Move Generation**
   - From any state, generate valid moves by sliding tiles into empty space
   - Typically 2-4 possible moves (up, down, left, right)
   - Avoid generating the parent state to prevent cycles

4. **Path Reconstruction**
   - Once goal is found, backtrack through parent pointers
   - Reconstruct sequence of moves from initial to goal state

#### Optimizations

1. **Solvability Check**
   - Not all N-puzzle configurations are solvable
   - Use inversion count to determine solvability before search
   - For odd-width grids: solvable if inversion count is even
   - For even-width grids: more complex rules involving blank position

2. **Memory Management**
   - Use efficient data structures (hash maps for closed set)
   - Consider iterative deepening A* (IDA*) for memory-constrained environments

3. **Tie Breaking**
   - When f(n) values are equal, prefer states closer to goal (higher g(n))
   - Helps avoid exploring unnecessary branches

#### Complexity Analysis

- **Time Complexity**: O(b^d) where b is branching factor (~3) and d is solution depth
- **Space Complexity**: O(b^d) for storing open and closed sets
- **Optimality**: Guaranteed to find optimal solution with admissible heuristic
- **Completeness**: Always finds solution if one exists

#### Performance Characteristics

- **8-Puzzle**: Typically solved in milliseconds, max ~31 moves
- **15-Puzzle**: Can take seconds to minutes, max ~80 moves  
- **24-Puzzle**: Very challenging, may require specialized techniques

## Usage

### Input Format
The program expects the puzzle size followed by the initial configuration:
```
3
2 8 3
1 6 4  
7 0 5
```

### Output Format
Sequence of states from initial to goal, followed by timing information:
```
2 8 3
1 6 4
7 0 5

2 8 3
1 0 4
7 6 5

...

1 2 3
4 5 6
7 8 0

# TIMES_MS: alg=45
```

### Environment Variables
- `FMI_TIME_ONLY`: Set to only measure execution time without printing solution steps

## Building and Running

### Go Implementation
```bash
# Build
cd go && go build -o n-puzzle n-puzzle.go

# Run
echo "3
2 8 3
1 6 4
7 0 5" | ./n-puzzle

# Time only
FMI_TIME_ONLY=1 echo "3
2 8 3
1 6 4
7 0 5" | ./n-puzzle
```

### Using Makefile
```bash
# Build all implementations
make build

# Test with sample input
make test

# Run performance benchmark
make bench
```

## Implementation Notes

- The algorithm guarantees finding the optimal (shortest) solution
- Performance varies significantly based on initial configuration difficulty
- Unsolvable puzzles are detected before search begins
- Memory usage grows exponentially with puzzle size
- For puzzles larger than 4×4, consider using IDA* or other memory-efficient variants

## References

- Hart, P. E.; Nilsson, N. J.; Raphael, B. (1968). "A Formal Basis for the Heuristic Determination of Minimum Cost Paths"
- Korf, R. E. (1985). "Depth-first iterative-deepening: An optimal admissible tree search"
- Reinefeld, A. (1993). "Complete Solution of the Eight-Puzzle and the Benefit of Node Ordering in IDA*"
