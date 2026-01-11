# 🧠 Algorithms Solutions

A collection of algorithm implementations in **Go** and **Python**. This repository contains solutions to various algorithmic problems including search algorithms, optimization problems, and AI/ML challenges.
You can find more details and theory for the course in this github - https://github.com/ElitsaY/AI_FMI_course

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

2. Install Python dependencies (handled automatically by the Makefile for temp venvs; for persistent venvs they are installed during `make venv <task>`). Consolidated dependencies live in the root `requirements.txt`.

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

```bash
make run <task-name>          # Auto-detects language and runs
make run frog-leap-puzzle     # Example: Go
make run iris                 # Example: Python
make run decision-tree        # Example: Python (ID3)
make run neural-networks      # Example: Python (MLP, interactive)
```

### Building Tasks (Go)

```bash
make build <task-name>
make build n-queens
```

### Testing Tasks

```bash
make test <task-name>
make test tsp
```

### Cleaning

```bash
make clean <task-name>        # Clean specific task
make clean-venvs              # Remove Python venvs
make clean-all                # Remove ALL binaries and .judge artifacts
```

## 🐍 Python Virtual Environments

- Temporary venv (auto): `make py <task>` — creates, installs dependencies from root `requirements.txt`, runs, then cleans up.
- Persistent venv: `make venv <task>` then `make run <task>` — installs from root `requirements.txt`, reuses environment.

## 🏗️ Creating New Tasks

```bash
make new <task-name>
make new my-algorithm
```

## 📁 Project Structure (high level)

```
algorithms-solutions/
├── Makefile
├── README.md
├── requirements.txt
├── venv-run.sh
├── python-with-venv.sh
│
├── frog-leap-puzzle/
│   └── go/
├── n-queens/
│   └── go/
├── n-puzzle/
│   └── go/
├── tsp/
│   └── go/
├── iris/
│   └── python/
├── tic-tac-toe/
│   └── python/
├── naive-bayes-classifier/
│   └── python/
├── decision-tree/
│   └── python/
└── algorithms/
    ├── beam-search/
    ├── dfs-and-bfs/
    ├── dijkstra/
    ├── genetic/
    └── minmax/
```

## 📚 Study Guide Coverage (Data Mining)

Task READMEs include theory, pseudocode, complexity, and exam tips:
- Frog Leap Puzzle — BFS/DFS/A* state modeling
- N-Queens — Backtracking and CSP
- N-Puzzle — A* and admissible heuristics (Manhattan)
- TSP — Exact vs heuristic (Nearest Neighbor, 2-Opt)
- Iris — KNN, scaling, cross-validation
- Naive Bayes Classifier — MAP rule, smoothing, variants
- Tic-Tac-Toe — Minimax and alpha-beta pruning
- Decision Tree — ID3 (entropy, information gain), pruning (REP)
- Neural Networks — MLP activations (sigmoid/tanh), backpropagation, XOR learnability

## 🔧 Troubleshooting

- Missing Python module: ensure dependencies are installed via `make venv <task>` or `pip install -r requirements.txt`.
- Judge artifacts: use `make clean <task>` or `make clean-all`.
- Go build issues: initialize/tidy modules inside `<task>/go/`.

## 📜 License

See individual task directories for specific licensing information.
