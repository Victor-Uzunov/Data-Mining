# Algorithms Solutions

A comprehensive collection of classic algorithm implementations with automated testing, performance benchmarking, and CI/CD integration. This repository focuses on providing optimal solutions to well-known computational problems with detailed explanations and multiple language support.

## 🧩 Featured Problems

### [Frog Leap Puzzle](./frog-leap-puzzle/)
A classic puzzle where frogs must leap over each other to switch positions. The implementation uses an iterative approach with strategic move prioritization to find the optimal solution sequence.

**Algorithm**: Iterative state-space search with heuristic move ordering  
**Complexity**: O(n²) moves for n frogs per side  
**Languages**: Go

### [N-Puzzle Solver](./n-puzzle/)
The famous sliding puzzle (8-puzzle, 15-puzzle, etc.) solver using optimal search algorithms. Finds the shortest sequence of moves to reach the goal state.

**Algorithm**: A* search with Manhattan distance heuristic  
**Complexity**: O(b^d) where b ≈ 3 and d is solution depth  
**Languages**: Go

## 🏗️ Repository Structure

Each algorithm solution follows a consistent, professional structure for easy navigation and testing:

```
algorithms-solutions/
├── .github/workflows/
│   └── test-all.yml          # Automated CI/CD pipeline
├── .gitignore                # Global gitignore patterns
├── go.mod                    # Go module dependencies
├── Makefile                  # Root-level build automation
├── README.md                 # This overview document
├── requirements.txt          # Python testing dependencies
├── test/                     # Shared testing utilities
│   └── main.go
└── <problem-name>/           # Individual algorithm solutions
    ├── README.md             # Problem description & algorithm analysis
    ├── Makefile              # Problem-specific build targets
    └── go/                   # Language-specific implementations
        └── <solution>.go     # Main implementation file
```

## 🔧 Technology Stack

### Primary Languages & Frameworks

#### 🔥 **Go** (Primary Implementation Language)
- **Version**: Go 1.21+
- **Benefits**: High performance, concurrent execution, static typing
- **Build**: Native binary compilation
- **Usage**: Production-ready implementations with optimal performance

#### 🐍 **Python** (Testing & Prototyping)
- **Version**: Python 3.11+
- **Benefits**: Rapid development, extensive libraries, readable code
- **Usage**: Test harnesses, algorithm prototyping, data analysis

### Development Tools

- **Build System**: Make (cross-platform compatibility)
- **CI/CD**: GitHub Actions (automated testing across environments)
- **Performance**: Built-in timing and benchmarking utilities
- **Quality**: Automated linting, formatting, and error checking

## 🚀 Quick Start

### Prerequisites
```bash
# Go installation (required)
go version  # Should be 1.21+

# Python installation (for testing)
python3 --version  # Should be 3.11+

# Make utility
make --version
```

### Running Solutions

#### Individual Problem
```bash
# Navigate to specific problem
cd n-puzzle

# Build and test
make build
make test

# Run with custom input
echo "3
2 8 3
1 6 4
7 0 5" | ./go/n-puzzle
```

#### All Problems
```bash
# From repository root
make build-all    # Build all implementations
make test-all     # Run all test suites
make bench-all    # Performance benchmarking
```

### Performance Measurement

All implementations include built-in timing capabilities:

```bash
# Standard execution with solution output
./solution < input.txt

# Time-only mode (for benchmarking)
FMI_TIME_ONLY=1 ./solution < input.txt
```

Output format includes standardized timing:
```
# TIMES_MS: alg=150
```

## 📊 Performance Characteristics

| Problem | Algorithm | Time Complexity | Space Complexity | Typical Performance |
|---------|-----------|----------------|------------------|-------------------|
| Frog Leap | Iterative Search | O(n²) | O(1) | < 1ms for n ≤ 10 |
| N-Puzzle | A* Search | O(b^d) | O(b^d) | < 100ms for 8-puzzle |

## 🧪 Testing Framework

### Automated Testing
- **Unit Tests**: Algorithm correctness validation
- **Integration Tests**: End-to-end solution verification  
- **Performance Tests**: Timing and memory usage benchmarks
- **Regression Tests**: Ensure optimizations don't break functionality

### Test Data
- **Generated Cases**: Programmatically created test inputs
- **Edge Cases**: Boundary conditions and corner cases
- **Benchmark Suite**: Standard problem instances for comparison

### Continuous Integration
```yaml
# .github/workflows/test-all.yml
- Builds all implementations
- Runs comprehensive test suites  
- Validates performance benchmarks
- Checks code quality and formatting
```

## 🎯 Problem Categories

### Search & Optimization
- **N-Puzzle**: Optimal pathfinding in state space
- **Frog Leap**: Constraint satisfaction with move ordering

### Future Additions (Planned)
- **A* Pathfinding**: Grid-based optimal route finding
- **Sudoku Solver**: Constraint propagation + backtracking
- **Traveling Salesman**: Dynamic programming approaches
- **Knapsack Problem**: Multiple optimization variants

## 🤝 Contributing

### Adding New Problems
1. Create problem directory: `mkdir new-problem`
2. Add README with problem description and algorithm analysis
3. Implement solution following project conventions
4. Add comprehensive test cases
5. Update root README with problem overview

### Code Style Guidelines
- **Go**: Follow `gofmt` and `golint` standards
- **Documentation**: Comprehensive algorithm explanations in README
- **Testing**: Include both correctness and performance tests
- **Naming**: Use descriptive, consistent naming conventions

### Implementation Requirements
- **Input/Output**: Standardized format across all solutions
- **Timing**: Built-in performance measurement
- **Error Handling**: Graceful handling of invalid inputs
- **Documentation**: Algorithm complexity analysis and references

## 📚 Educational Value

### Algorithm Analysis
Each implementation includes:
- **Time/Space Complexity**: Big-O analysis with explanations
- **Algorithm Description**: Step-by-step methodology
- **Optimization Techniques**: Performance improvements and trade-offs
- **Comparative Analysis**: Alternative approaches and their merits

### Learning Outcomes
- **Problem Solving**: Breaking down complex problems into manageable components
- **Algorithm Design**: Choosing appropriate data structures and strategies
- **Performance Optimization**: Understanding computational trade-offs
- **Software Engineering**: Professional development practices

## 📖 References & Further Reading

### Academic Sources
- "Introduction to Algorithms" by Cormen, Leiserson, Rivest, and Stein
- "Algorithm Design Manual" by Steven S. Skiena
- "Artificial Intelligence: A Modern Approach" by Russell and Norvig

### Online Resources
- [Algorithm Visualizations](https://visualgo.net/)
- [Big-O Complexity Analysis](https://www.bigocheatsheet.com/)
- [Competitive Programming Resources](https://codeforces.com/)

---

**License**: MIT | **Maintainer**: AI Algorithms Team | **Last Updated**: November 2025
