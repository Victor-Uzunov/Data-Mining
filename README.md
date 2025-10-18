# Algorithms Solutions

A collection of algorithm implementations with automated testing and CI/CD integration.

## Repository Structure

Each algorithm solution is organized in its own directory with a consistent structure:

```
algorithms-solutions/
├── .github/workflows/
│   └── test-all.yml          # Generic CI/CD pipeline for all tasks
├── .gitignore                # Global gitignore
├── go.mod                    # Go module for Go solutions
├── Makefile                  # Generic Makefile for all tasks
├── README.md                 # This file
├── requirements.txt          # Python dependencies for testing
└── <task-name>/              # Individual algorithm solutions
    ├── README.md             # Problem description and testing
    ├── go/                   # Go implementation (optional)
    │   └── solution.go
    ├── python/               # Python implementation (optional)
    │   └── solution.py
    ├── java/                 # Java implementation (optional)
    │   └── Solution.java
    └── cpp/                  # C++ implementation (optional)
        └── solution.cpp
```

## Supported Programming Languages

We support multiple programming languages for algorithm implementations. Contributors can choose their preferred language or implement solutions in multiple languages:

### 🔥 **Go** (Primary)
- **File**: `<task-name>/go/solution.go`
- **Build**: Compiles to native binary
- **Run**: `./solution`
- **Requirements**: Go 1.21+

### 🐍 **Python**
- **File**: `<task-name>/python/solution.py`
- **Build**: No compilation needed
- **Run**: `python3 solution.py`
- **Requirements**: Python 3.11+

### ☕ **Java**
- **File**: `<task-name>/java/Solution.java`
- **Build**: Compiles to `.class` files
- **Run**: `java Solution`
- **Requirements**: Java 17+

### ⚡ **C++**
- **File**: `<task-name>/cpp/solution.cpp`
- **Build**: Compiles with `g++`
- **Run**: `./solution`
- **Requirements**: GCC/G++ compiler

### Language Priority
When multiple implementations exist for a task, the system auto-detects in this order:
1. Go
2. Python
3. Java
4. C++

You can specify a language explicitly:
```bash
make run TASK=my-algorithm LANGUAGE=python N=5
```

## Quick Start

### Using the Generic System

The root-level Makefile provides commands that work with any task:

```bash
# List all available tasks
make list-tasks

# Build a specific task (auto-detects language)
make build TASK=frog-leap-puzzle

# Run a task with input (works with any supported language)
make run TASK=frog-leap-puzzle N=3 LANG=go
make run TASK=frog-leap-puzzle N=3 LANG=python
make run TASK=frog-leap-puzzle N=3 LANG=java

# Test a task (tests all available language implementations)
make test TASK=frog-leap-puzzle

# Format code for a task
make fmt TASK=frog-leap-puzzle
```

### Batch Operations

```bash
# Build all tasks
make build-all

# Test all tasks
make test-all

# Clean all tasks
make clean-all
```

### Creating a New Task

```bash
# Initialize a new task with basic structure
make init-task TASK=new-algorithm

# This creates:
# - new-algorithm/README.md (template)
# - new-algorithm/go/solution.go (template with proper structure)
```

## Development Workflow

### Setting Up Development Environment

```bash
# Install development dependencies
make deps

# This installs:
# - fmi-ai-judge (for testing)
# - goimports (for code formatting)
```

### Adding a New Solution

1. **Create the task structure:**
   ```bash
   make init-task TASK=my-algorithm
   ```

2. **Implement your solution in `my-algorithm/go/solution.go`:**
   ```go
   package main

   import (
       "bufio"
       "fmt"
       "os"
       "strconv"
       "time"
   )

   func solve(input int) []string {
       // Your algorithm implementation
       return []string{"result"}
   }

   func main() {
       timeOnly := os.Getenv("FMI_TIME_ONLY") == "1"
       
       scanner := bufio.NewScanner(os.Stdin)
       scanner.Scan()
       n, _ := strconv.Atoi(scanner.Text())
       
       start := time.Now()
       result := solve(n)
       duration := time.Since(start)
       
       if timeOnly {
           fmt.Printf("# TIMES_MS: alg=%d\n", duration.Nanoseconds()/1000000)
       } else {
           for _, line := range result {
               fmt.Println(line)
           }
       }
   }
   ```

3. **Test your solution:**
   ```bash
   make build TASK=my-algorithm
   make run TASK=my-algorithm N=5
   make test TASK=my-algorithm
   ```

## Contributing

Here's how to add your algorithm solution:

### 1. Fork and Clone

```bash
git clone https://github.com/YOUR_USERNAME/algorithms-solutions.git
cd algorithms-solutions
```

### 2. Set Up Development Environment

```bash
# Install dependencies
make deps

# Verify the existing tests pass
make test-all
```

### 3. Create Your Solution

```bash
# Create a new branch for your contribution
git checkout -b add-your-algorithm-name

# Initialize your new algorithm
make init-task TASK=your-algorithm-name

# Implement your solution in your preferred language:
# your-algorithm-name/go/solution.go (Go)
# your-algorithm-name/python/solution.py (Python) 
# your-algorithm-name/java/Solution.java (Java)
# your-algorithm-name/cpp/solution.cpp (C++)

# Update your-algorithm-name/README.md with problem description
```

### 4. Test Your Solution

```bash
# Build and test your solution
make build TASK=your-algorithm-name
make run TASK=your-algorithm-name N=10 LANG=go
make test TASK=your-algorithm-name

# Ensure all existing tests still pass
make test-all

```

### 5. Submit a Pull Request

1. **Commit your changes:**
   ```bash
   git add .
   git commit -m "Add: your-algorithm-name solution in [Go/Python/Java/C++]"
   git push origin add-your-algorithm-name
   ```

2. **Create a Pull Request** on GitHub from your branch to the main repository with:
   - Clear description of the algorithm and problem solved
   - Programming language(s) used
   - Performance characteristics (time/space complexity)
   - Test cases and expected outputs

### 6. Automated Testing

Once you submit a PR, our CI pipeline will automatically:
- ✅ Build your solution
- ✅ Run correctness tests using `fmi-ai-judge`
- ✅ Run performance benchmarks
- ✅ Validate code formatting
- ✅ Test timing mode functionality

The pipeline tests all solutions in parallel, so your changes won't break existing algorithms.

## Solution Requirements

All solutions should meet these standards:

- **Correctness**: Pass all test cases
- **Performance**: Handle specified input sizes efficiently
- **Code Quality**: Well-formatted, readable Go code
- **Timing Support**: Support `FMI_TIME_ONLY=1` environment variable
- **Input/Output**: Read from stdin, write to stdout (exact format as specified)

## Available Algorithms

Currently implemented algorithms:

- **frog-leap-puzzle** - Classic frog leap puzzle with optimal DFS solution

## Need Help?

- Check existing solutions for examples
- Read task-specific README files for problem descriptions
- Use `make help` for available commands
- Open an issue for questions or bug reports

## Requirements

- **Go 1.21+** - for building solutions
- **Python 3.11+** - for testing with fmi-ai-judge
- **Make** - for using the build system

## License

This project is for educational purposes.
