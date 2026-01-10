# 🧠 Algorithms Solutions

A collection of algorithm implementations in **Go** and **Python**. This repository contains solutions to various algorithmic problems including search algorithms, optimization problems, and AI/ML challenges.

## 🚀 Quick Start

### Prerequisites

- **Go** (1.21+) - for Go solutions
- **Python** (3.8+) - for Python solutions

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd algorithms-solutions
```

2. Install Python dependencies (optional, handled automatically):
```bash
pip install -r requirements.txt
```

## 📋 Available Commands

### Discovery & Help

```bash
make              # Show help menu
make help         # Show detailed help
make ls           # List all available tasks
```

### Running Tasks

The simplest way to run any task:

```bash
make run <task-name>          # Auto-detects language and runs
make run frog-leap-puzzle     # Example: run frog-leap-puzzle
make run iris                 # Example: run iris dataset classifier
```

### Building Tasks

For compiled languages (Go):

```bash
make build <task-name>        # Build the task
make build n-queens           # Example
```

### Testing Tasks

Test with the fmi-ai-judge framework:

```bash
make test <task-name>         # Build and test
make test tsp                 # Example: test traveling salesman
```

### Cleaning

```bash
make clean <task-name>        # Clean specific task
make clean-venvs              # Clean all Python virtual environments
```

## 🐍 Python-Specific Commands

Python tasks automatically use **temporary virtual environments** that are cleaned up after execution. No manual venv management needed!

### Run with Temporary venv (Recommended)

```bash
make py <task-name>           # Run Python task with auto-cleanup venv
make py iris                  # Example
make py naive-bayes-classifier
```

### Create Persistent venv (For Development)

If you want to develop and need a persistent environment:

```bash
make venv <task-name>         # Create .venv in task directory
cd <task-name>/python
source .venv/bin/activate     # Activate the environment
```

### Clean All Python venvs

```bash
make clean-venvs              # Remove all Python virtual environments
```

## 🆕 Creating New Tasks

Create a new task structure quickly:

```bash
make new <task-name>          # Creates directories and templates
make new my-algorithm         # Example

# This creates:
# my-algorithm/
#   ├── go/main.go           (Go template)
#   ├── python/solution.py   (Python template)
#   └── README.md            (Task documentation)
```

## 📁 Project Structure

```
algorithms-solutions/
├── Makefile                  # Main build system (single source of truth)
├── README.md                 # This file
├── requirements.txt          # Python dependencies
├── venv-run.sh              # Auto-venv script (used internally)
├── python-with-venv.sh      # Venv wrapper (used internally)
│
├── algorithms/               # General algorithms
│   ├── beam-search/
│   ├── dfs-and-bfs/
│   ├── dijkstra/
│   ├── genetic/
│   └── minmax/
│
├── frog-leap-puzzle/        # Frog leap puzzle solver
│   ├── go/
│   └── README.md
│
├── n-queens/                # N-Queens problem
├── tsp/                     # Traveling Salesman Problem
├── knapsack/                # Knapsack problem
├── iris/                    # Iris dataset classifier
├── naive-bayes-classifier/  # Naive Bayes implementation
└── tic-tac-toe/            # Tic-tac-toe with AI
```

**Note:** Each task contains only the necessary files. No task-level Makefiles needed - the main Makefile handles everything!

Each task typically contains:
- `go/` - Go implementation (optional)
- `python/` - Python implementation (optional)
- `README.md` - Task-specific documentation

## 💡 Usage Examples

### Example 1: Run a Go program
```bash
make run frog-leap-puzzle
# Auto-detects Go, compiles and runs
```

### Example 2: Run a Python program with automatic venv
```bash
make py iris
# Creates temp venv, installs deps, runs, then cleans up
```

### Example 3: Develop a Python task
```bash
# Create persistent venv for development
make venv naive-bayes-classifier

# Activate it manually
cd naive-bayes-classifier/python
source .venv/bin/activate

# Now you can edit and run directly
python naive-bayes-classifier.py
```

### Example 4: Create and run a new task
```bash
make new bubble-sort
# Edit the generated files
# Then run it
make run bubble-sort
```

## 🔧 How Python Virtual Environments Work

### Automatic (Temporary) venv
When you run `make py <task>` or `make run <python-task>`:
1. A temporary venv is created in `/tmp/`
2. Dependencies from `requirements.txt` are installed
3. Your script runs
4. The venv is automatically deleted

**Pros:** No cleanup needed, always fresh environment  
**Cons:** Slower for repeated runs

### Manual (Persistent) venv
When you run `make venv <task>`:
1. A `.venv/` folder is created in `<task>/python/`
2. Dependencies are installed once
3. You manually activate/deactivate
4. Faster for development

**Pros:** Fast, great for development  
**Cons:** Need to manage manually

## 🎯 Task Categories

### Search & Traversal
- **dfs-and-bfs** - Depth-first and breadth-first search
- **beam-search** - Beam search algorithm
- **dijkstra** - Dijkstra's shortest path

### Optimization
- **genetic** - Genetic algorithms and crossovers
- **tsp** - Traveling Salesman Problem
- **knapsack** - 0/1 Knapsack problem

### Game Theory & AI
- **minmax** - Minimax algorithm
- **tic-tac-toe** - Tic-tac-toe with AI
- **frog-leap-puzzle** - Frog leap puzzle solver

### Constraint Satisfaction
- **n-queens** - N-Queens problem
- **n-puzzle** - N-puzzle solver

### Machine Learning
- **iris** - Iris dataset classifier
- **naive-bayes-classifier** - Naive Bayes from scratch

## 🛠️ Advanced Usage

### Language Priority
If a task has both Go and Python implementations, auto-detect prefers:
1. Go
2. Python

You can always explicitly run Python with `make py <task>`.

### Custom Requirements
To add task-specific Python dependencies:
1. Create `<task>/python/requirements.txt`
2. The venv system will automatically install them

### Testing with fmi-ai-judge
Ensure `fmi-ai-judge` is installed:
```bash
pip install fmi-ai-judge
```

Then run tests:
```bash
make test <task-name>
```

## 📝 Development Workflow

### Recommended workflow for new tasks:

1. **Create the task structure:**
   ```bash
   make new my-new-algorithm
   ```

2. **Implement your solution:**
   ```bash
   cd my-new-algorithm/python
   # Edit solution.py
   ```

3. **Test with temporary venv:**
   ```bash
   make py my-new-algorithm
   ```

4. **For intensive development, create persistent venv:**
   ```bash
   make venv my-new-algorithm
   cd my-new-algorithm/python
   source .venv/bin/activate
   # Now develop/test rapidly
   ```

5. **Clean up when done:**
   ```bash
   make clean my-new-algorithm
   ```

## 🧪 Testing & Validation

### Manual Testing
```bash
make run <task>              # Quick test
```

### Automated Testing
```bash
make test <task>             # Test with fmi-ai-judge
```

## 🤝 Contributing

1. Create a new task: `make new <task-name>`
2. Implement your solution in Go or Python (or both!)
3. Add a README.md describing the problem
4. Test it: `make test <task-name>`
5. Commit and push

## 📚 Additional Resources

- [fmi-ai-judge Documentation](https://pypi.org/project/fmi-ai-judge/)
- Go: https://go.dev/doc/
- Python: https://docs.python.org/3/

## 🐛 Troubleshooting

### "Task not found"
```bash
make ls                      # Check available tasks
```

### Python import errors
```bash
make clean-venvs             # Clean all venvs
make py <task>               # Try again with fresh venv
```

### Build errors for Go
```bash
cd <task>/go
go mod init <module-name>
go mod tidy
```

## ⚡ What's New?

### Simplified Build System
- **Single Makefile** - No more task-level Makefiles cluttering your project
- **Only Go & Python** - Removed Java and C++ complexity
- **Auto-detection** - Smart language detection for each task
- **Cleaner commands** - Simple, intuitive syntax like `make run <task>`
- **Focus on individual tasks** - Work on one task at a time for better clarity

## 📜 License

See individual task directories for specific licensing information.

---

**Happy Coding! 🚀**

For questions or issues, please open a GitHub issue.
