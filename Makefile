# Algorithms Solutions - Simplified Makefile
# Easy-to-use commands for building, running, and testing algorithm solutions
# Supports: Go and Python

# Colors for output
GREEN := \033[0;32m
YELLOW := \033[1;33m
RED := \033[0;31m
BLUE := \033[0;34m
NC := \033[0m # No Color

# Auto-detect all tasks (only Go and Python)
TASKS := $(shell find . -maxdepth 2 -type d \( -name "go" -o -name "python" \) | sed 's|^\./||g' | sed 's|/[^/]*$$||g' | sort -u)

.DEFAULT_GOAL := help

# ============================================================================
# HELP & DISCOVERY
# ============================================================================

help:
	@echo "$(BLUE)╔════════════════════════════════════════════════════════════╗$(NC)"
	@echo "$(BLUE)║  Algorithms Solutions - Quick Reference                   ║$(NC)"
	@echo "$(BLUE)╚════════════════════════════════════════════════════════════╝$(NC)"
	@echo ""
	@echo "$(GREEN)📋 Quick Start:$(NC)"
	@echo "  make ls                    - List all available tasks"
	@echo "  make run TASK=<name> [LANG=go|python] - Run a task"
	@echo "  make test TASK=<name> [LANG=go|python] - Build and test a task"
	@echo ""
	@echo "$(GREEN)🔧 Common Commands:$(NC)"
	@echo "  make run TASK=<name>       - Run task (prompts if both exist)"
	@echo "  make run TASK=<name> LANG=go       - Run Go implementation"
	@echo "  make run TASK=<name> LANG=python   - Run Python implementation"
	@echo "  make test TASK=<name>      - Test task with fmi-ai-judge"
	@echo "  make test TASK=<name> LANG=go      - Test Go implementation"
	@echo "  make test TASK=<name> LANG=python  - Test Python implementation"
	@echo "  make build TASK=<name>     - Build task (for Go)"
	@echo "  make clean TASK=<name>     - Clean build artifacts"
	@echo "  make clean-all             - Remove ALL binaries and artifacts"
	@echo ""
	@echo "$(GREEN)🐍 Python-Specific:$(NC)"
	@echo "  make py TASK=<name>        - Run Python task with temp venv (auto-cleanup)"
	@echo "  make venv TASK=<name>      - Create persistent venv for development"
	@echo "  make clean-venvs           - Remove all Python virtual environments"
	@echo ""
	@echo "$(GREEN)🆕 Create New Task:$(NC)"
	@echo "  make new TASK=<name>       - Create new task structure"
	@echo ""
	@echo "$(YELLOW)💡 Examples:$(NC)"
	@echo "  make ls"
	@echo "  make run TASK=frog-leap-puzzle"
	@echo "  make run TASK=knapsack LANG=go"
	@echo "  make run TASK=iris LANG=python"
	@echo "  make test TASK=n-queens LANG=go"
	@echo "  make py TASK=iris"
	@echo "  make venv TASK=naive-bayes-classifier"
	@echo ""

ls: list
list:
	@echo "$(GREEN)📦 Available Tasks:$(NC)"
	@echo ""
	@for task in $(TASKS); do \
		if [ -d "$$task" ]; then \
			langs=$$(find $$task -maxdepth 1 -type d \( -name "go" -o -name "python" \) 2>/dev/null | sed 's|.*/||g' | sort | tr '\n' ' ' | sed 's/ *$$//'); \
			if [ -n "$$langs" ]; then \
				printf "  $(BLUE)%-30s$(NC) [%s]\n" "$$task" "$$langs"; \
			fi; \
		fi; \
	done
	@echo ""

# ============================================================================
# MAIN COMMANDS
# ============================================================================

# Makefile
run: ## Run a task
	@if [ -z "$(TASK)" ]; then \
		echo "$(RED)Error: Please specify a task name$(NC)"; \
		echo "Usage: make run TASK=<task-name> [LANG=go|python]"; \
		exit 1; \
	fi
	@if [ ! -d "$(TASK)" ]; then \
		echo "$(RED)Error: Task '$(TASK)' not found$(NC)"; \
		echo "Run 'make ls' to see available tasks"; \
		exit 1; \
	fi
	@HAS_GO=0; \
	HAS_PYTHON=0; \
	if [ -d "$(TASK)/go" ]; then HAS_GO=1; fi; \
	if [ -d "$(TASK)/python" ]; then HAS_PYTHON=1; fi; \
	if [ $$HAS_GO -eq 0 ] && [ $$HAS_PYTHON -eq 0 ]; then \
		echo "$(RED)Error: No implementation found for '$(TASK)'$(NC)"; \
		exit 1; \
	fi; \
	SELECTED_LANG=""; \
	if [ -n "$(LANG)" ] && [ "$(LANG)" != "C.UTF-8" ] && [ "$(LANG)" != "en_US.UTF-8" ]; then \
		SELECTED_LANG="$(LANG)"; \
	fi; \
	if [ -n "$$SELECTED_LANG" ]; then \
		if [ "$$SELECTED_LANG" = "go" ] && [ $$HAS_GO -eq 0 ]; then \
			echo "$(RED)Error: Task '$(TASK)' has no Go implementation$(NC)"; \
			exit 1; \
		fi; \
		if [ "$$SELECTED_LANG" = "python" ] && [ $$HAS_PYTHON -eq 0 ]; then \
			echo "$(RED)Error: Task '$(TASK)' has no Python implementation$(NC)"; \
			exit 1; \
		fi; \
	elif [ $$HAS_GO -eq 1 ] && [ $$HAS_PYTHON -eq 1 ]; then \
		echo "$(YELLOW)Both Go and Python implementations available for '$(TASK)'$(NC)"; \
		echo "Please specify which to run:"; \
		echo "  make run TASK=$(TASK) LANG=go"; \
		echo "  make run TASK=$(TASK) LANG=python"; \
		exit 1; \
	elif [ $$HAS_GO -eq 1 ]; then \
		SELECTED_LANG="go"; \
	else \
		SELECTED_LANG="python"; \
	fi; \
	echo "$(GREEN)🚀 Running: $(TASK) ($$SELECTED_LANG)$(NC)"; \
	if [ "$$SELECTED_LANG" = "go" ]; then \
		(cd $(TASK)/go && go run .); \
	elif [ "$$SELECTED_LANG" = "python" ]; then \
		if [ -d "$(TASK)/python/.venv" ]; then \
			(cd $(TASK)/python && . .venv/bin/activate && python *.py); \
		else \
			(cd $(TASK)/python && ../../venv-run.sh python *.py); \
		fi; \
	fi


build: ## Build a task (for Go)
	@if [ -z "$(filter-out build,$@)" ] && [ -z "$(MAKECMDGOALS)" ]; then \
		echo "$(RED)Error: Please specify a task name$(NC)"; \
		echo "Usage: make build <task-name>"; \
		exit 1; \
	fi
	@TASK="$(word 2,$(MAKECMDGOALS))"; \
	if [ ! -d "$$TASK" ]; then \
		echo "$(RED)Error: Task '$$TASK' not found$(NC)"; \
		echo "Run 'make ls' to see available tasks"; \
		exit 1; \
	fi; \
	if [ -d "$$TASK/go" ]; then \
		LANG="go"; \
	elif [ -d "$$TASK/python" ]; then \
		LANG="python"; \
	else \
		echo "$(RED)Error: No implementation found for '$$TASK'$(NC)"; \
		exit 1; \
	fi; \
	echo "$(GREEN)🔨 Building: $$TASK ($$LANG)$(NC)"; \
	if [ "$$LANG" = "go" ]; then \
		GOFILE=$$(cd $$TASK/go && ls *.go 2>/dev/null | head -n 1); \
		BINARY=$${GOFILE%.go}; \
		(cd $$TASK/go && go build -o $$BINARY .); \
	elif [ "$$LANG" = "python" ]; then \
		echo "$(YELLOW)Python doesn't need compilation$(NC)"; \
	fi

test: ## Test a task with fmi-ai-judge
	@if [ -z "$(TASK)" ]; then \
		echo "$(RED)Error: Please specify a task name$(NC)"; \
		echo "Usage: make test TASK=<task-name> [LANG=go|python]"; \
		exit 1; \
	fi
	@if [ ! -d "$(TASK)" ]; then \
		echo "$(RED)Error: Task '$(TASK)' not found$(NC)"; \
		echo "Run 'make ls' to see available tasks"; \
		exit 1; \
	fi
	@HAS_GO=0; \
	HAS_PYTHON=0; \
	if [ -d "$(TASK)/go" ]; then HAS_GO=1; fi; \
	if [ -d "$(TASK)/python" ]; then HAS_PYTHON=1; fi; \
	if [ $$HAS_GO -eq 0 ] && [ $$HAS_PYTHON -eq 0 ]; then \
		echo "$(RED)Error: No implementation found for '$(TASK)'$(NC)"; \
		exit 1; \
	fi; \
	SELECTED_LANG=""; \
	if [ -n "$(LANG)" ] && [ "$(LANG)" != "C.UTF-8" ] && [ "$(LANG)" != "en_US.UTF-8" ]; then \
		SELECTED_LANG="$(LANG)"; \
	fi; \
	if [ -n "$$SELECTED_LANG" ]; then \
		if [ "$$SELECTED_LANG" = "go" ] && [ $$HAS_GO -eq 0 ]; then \
			echo "$(RED)Error: Task '$(TASK)' has no Go implementation$(NC)"; \
			exit 1; \
		fi; \
		if [ "$$SELECTED_LANG" = "python" ] && [ $$HAS_PYTHON -eq 0 ]; then \
			echo "$(RED)Error: Task '$(TASK)' has no Python implementation$(NC)"; \
			exit 1; \
		fi; \
	elif [ $$HAS_GO -eq 1 ] && [ $$HAS_PYTHON -eq 1 ]; then \
		echo "$(YELLOW)Both Go and Python implementations available for '$(TASK)'$(NC)"; \
		echo "Please specify which to test:"; \
		echo "  make test TASK=$(TASK) LANG=go"; \
		echo "  make test TASK=$(TASK) LANG=python"; \
		exit 1; \
	elif [ $$HAS_GO -eq 1 ]; then \
		SELECTED_LANG="go"; \
	else \
		SELECTED_LANG="python"; \
	fi; \
	echo "$(GREEN)🔨 Building: $(TASK) ($$SELECTED_LANG)$(NC)"; \
	if [ "$$SELECTED_LANG" = "go" ]; then \
		GOFILE=$$(cd $(TASK)/go && ls *.go 2>/dev/null | head -n 1); \
		BINARY=$${GOFILE%.go}; \
		(cd $(TASK)/go && go build -o $$BINARY .); \
	fi; \
	echo "$(GREEN)🧪 Testing: $(TASK) ($$SELECTED_LANG)$(NC)"; \
	if [ "$$SELECTED_LANG" = "go" ]; then \
		if command -v judge >/dev/null 2>&1; then \
			GOFILE=$$(cd $(TASK)/go && ls *.go 2>/dev/null | head -n 1); \
			BINARY=$${GOFILE%.go}; \
			(cd $(TASK)/go && judge run --bench $$BINARY); \
		else \
			echo "$(YELLOW)fmi-ai-judge not found. Install with: pip install fmi-ai-judge$(NC)"; \
		fi; \
	elif [ "$$SELECTED_LANG" = "python" ]; then \
		if command -v judge >/dev/null 2>&1; then \
			PYFILE=$$(cd $(TASK)/python && ls *.py 2>/dev/null | head -n 1); \
			(cd $(TASK)/python && judge run --bench python3 $$PYFILE); \
		else \
			echo "$(YELLOW)fmi-ai-judge not found. Install with: pip install fmi-ai-judge$(NC)"; \
		fi; \
	fi

clean: ## Clean build artifacts for a task
	@if [ -z "$(filter-out clean,$@)" ] && [ -z "$(MAKECMDGOALS)" ]; then \
		echo "$(RED)Error: Please specify a task name$(NC)"; \
		echo "Usage: make clean <task-name>"; \
		exit 1; \
	fi
	@TASK="$(word 2,$(MAKECMDGOALS))"; \
	echo "$(YELLOW)🧹 Cleaning: $$TASK$(NC)"; \
	if [ -d "$$TASK/go" ]; then \
		GOFILE=$$(cd $$TASK/go && ls *.go 2>/dev/null | head -n 1); \
		if [ -n "$$GOFILE" ]; then \
			BINARY=$${GOFILE%.go}; \
			rm -f $$TASK/go/$$BINARY; \
		fi; \
		rm -rf $$TASK/go/.judge; \
	fi; \
	if [ -d "$$TASK/python" ]; then \
		rm -rf $$TASK/python/.judge; \
	fi; \
	echo "$(GREEN)✓ Clean complete$(NC)"

clean-all: ## Remove ALL binaries and artifacts
	@echo "$(YELLOW)🧹 Cleaning all binaries and artifacts...$(NC)"
	@for task in $(TASKS); do \
		if [ -d "$$task/go" ]; then \
			GOFILE=$$(cd $$task/go && ls *.go 2>/dev/null | head -n 1); \
			if [ -n "$$GOFILE" ]; then \
				BINARY=$${GOFILE%.go}; \
				rm -f $$task/go/$$BINARY; \
			fi; \
			rm -rf $$task/go/.judge; \
		fi; \
		if [ -d "$$task/python" ]; then \
			rm -rf $$task/python/.judge; \
		fi; \
	done
	@echo "$(GREEN)✓ All binaries and artifacts removed$(NC)"

# ============================================================================
# PYTHON-SPECIFIC COMMANDS
# ============================================================================

py: ## Run Python task with temporary venv (auto-cleanup)
	@if [ -z "$(filter-out py,$@)" ] && [ -z "$(MAKECMDGOALS)" ]; then \
		echo "$(RED)Error: Please specify a task name$(NC)"; \
		echo "Usage: make py <task-name>"; \
		exit 1; \
	fi
	@TASK="$(word 2,$(MAKECMDGOALS))"; \
	if [ ! -d "$$TASK/python" ]; then \
		echo "$(RED)Error: Task '$$TASK' has no Python implementation$(NC)"; \
		exit 1; \
	fi; \
	echo "$(GREEN)🐍 Running Python: $$TASK (temporary venv)$(NC)"; \
	(cd $$TASK/python && ../../venv-run.sh python *.py)

venv: ## Create persistent venv for a Python task
	@if [ -z "$(filter-out venv,$@)" ] && [ -z "$(MAKECMDGOALS)" ]; then \
		echo "$(RED)Error: Please specify a task name$(NC)"; \
		echo "Usage: make venv <task-name>"; \
		exit 1; \
	fi
	@TASK="$(word 2,$(MAKECMDGOALS))"; \
	if [ ! -d "$$TASK/python" ]; then \
		echo "$(RED)Error: Task '$$TASK' has no Python implementation$(NC)"; \
		exit 1; \
	fi; \
	echo "$(GREEN)🔧 Creating venv for: $$TASK$(NC)"; \
	(cd $$TASK/python && python3 -m venv .venv && \
	. .venv/bin/activate && \
	pip install --upgrade pip --quiet && \
	if [ -f ../../requirements.txt ]; then \
		pip install -r ../../requirements.txt --quiet; \
	fi && \
	if [ -f requirements.txt ]; then \
		pip install -r requirements.txt --quiet; \
	fi && \
	echo "$(GREEN)✓ Virtual environment created at $$TASK/python/.venv$(NC)" && \
	echo "$(YELLOW)To activate: cd $$TASK/python && source .venv/bin/activate$(NC)")

clean-venvs: ## Remove all Python virtual environments
	@echo "$(YELLOW)🧹 Cleaning Python virtual environments...$(NC)"
	@find . -type d -name ".venv" -path "*/python/.venv" -exec rm -rf {} + 2>/dev/null || true
	@echo "$(GREEN)✓ All Python virtual environments removed$(NC)"

# ============================================================================
# UTILITIES
# ============================================================================

new: ## Create a new task structure
	@if [ -z "$(filter-out new,$@)" ] && [ -z "$(MAKECMDGOALS)" ]; then \
		echo "$(RED)Error: Please specify a task name$(NC)"; \
		echo "Usage: make new <task-name>"; \
		exit 1; \
	fi
	@TASK="$(word 2,$(MAKECMDGOALS))"; \
	if [ -d "$$TASK" ]; then \
		echo "$(YELLOW)Warning: Task '$$TASK' already exists$(NC)"; \
		exit 1; \
	fi; \
	echo "$(GREEN)🆕 Creating new task: $$TASK$(NC)"; \
	mkdir -p $$TASK/{go,python}; \
	echo "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello from $$TASK\")\n}" > $$TASK/go/main.go; \
	echo "def main():\n    print('Hello from $$TASK')\n\nif __name__ == '__main__':\n    main()" > $$TASK/python/solution.py; \
	echo "# $$TASK\n\nTask description here.\n\n## Usage\n\n\`\`\`bash\nmake run $$TASK\n\`\`\`" > $$TASK/README.md; \
	echo "$(GREEN)✓ Task '$$TASK' created with Go and Python templates$(NC)"; \
	echo "$(YELLOW)Edit the files in $$TASK/go/ and $$TASK/python/$(NC)"

# Catch-all target for task names (allows "make run frog-leap-puzzle" syntax)
%:
	@:

.PHONY: help ls list run build test clean py venv clean-venvs new

