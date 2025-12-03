.PHONY: build test clean run

# Default language detection
LANGUAGE ?= python

build:
	@echo "Building tic-tac-toe solution..."
	@case $(LANGUAGE) in \
		go) go build -o go/tic-tac-toe ./go ;; \
		python) echo "Python doesn't need compilation" ;; \
		java) cd java && javac *.java ;; \
		cpp) cd cpp && g++ -o tic-tac-toe *.cpp ;; \
		*) echo "Unsupported language: $(LANGUAGE)" && exit 1 ;; \
	esac

test: build
	@echo "Running fmi-ai-judge tests..."
	@case $(LANGUAGE) in \
		go) \
			if command -v judge >/dev/null 2>&1; then \
				judge run --bench go/tic-tac-toe; \
			elif [ -f "../.venv/bin/judge" ]; then \
				../.venv/bin/judge run --bench go/tic-tac-toe; \
			elif [ -f ".venv/bin/judge" ]; then \
				.venv/bin/judge run --bench go/tic-tac-toe; \
			else \
				echo "Error: fmi-ai-judge not found. Install with: pip install fmi-ai-judge"; \
				exit 1; \
			fi ;; \
		python) \
			if command -v judge >/dev/null 2>&1; then \
				judge run --bench python/tictactoe.py; \
			elif [ -f "../.venv/bin/judge" ]; then \
				../.venv/bin/judge run --bench python/tictactoe.py; \
			elif [ -f ".venv/bin/judge" ]; then \
				.venv/bin/judge run --bench python/tictactoe.py; \
			else \
				echo "Error: fmi-ai-judge not found. Install with: pip install fmi-ai-judge"; \
				exit 1; \
			fi ;; \
		java) \
			if command -v judge >/dev/null 2>&1; then \
				judge run --bench java/Solution.class; \
			elif [ -f "../.venv/bin/judge" ]; then \
				../.venv/bin/judge run --bench java/Solution.class; \
			elif [ -f ".venv/bin/judge" ]; then \
				.venv/bin/judge run --bench java/Solution.class; \
			else \
				echo "Error: fmi-ai-judge not found. Install with: pip install fmi-ai-judge"; \
				exit 1; \
			fi ;; \
		cpp) \
			if command -v judge >/dev/null 2>&1; then \
				judge run --bench cpp/tic-tac-toe; \
			elif [ -f "../.venv/bin/judge" ]; then \
				../.venv/bin/judge run --bench cpp/tic-tac-toe; \
			elif [ -f ".venv/bin/judge" ]; then \
				.venv/bin/judge run --bench cpp/tic-tac-toe; \
			else \
				echo "Error: fmi-ai-judge not found. Install with: pip install fmi-ai-judge"; \
				exit 1; \
			fi ;; \
		*) echo "Unsupported language: $(LANGUAGE)" && exit 1 ;; \
	esac

run:
	@echo "Running tic-tac-toe in interactive mode..."
	@case $(LANGUAGE) in \
		go) ./go/tic-tac-toe ;; \
		python) python3 python/tictactoe.py ;; \
		java) cd java && java Solution ;; \
		cpp) ./cpp/tic-tac-toe ;; \
		*) echo "Unsupported language: $(LANGUAGE)" && exit 1 ;; \
	esac

clean:
	@echo "Cleaning build artifacts..."
	@rm -f go/tic-tac-toe
	@rm -f cpp/tic-tac-toe
	@rm -f java/*.class
	@rm -rf .judge
	@echo "Clean complete"

help:
	@echo "Tic-Tac-Toe Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build           - Build the solution (default: Python)"
	@echo "  make test            - Run fmi-ai-judge tests"
	@echo "  make run             - Run interactive game mode"
	@echo "  make clean           - Remove build artifacts"
	@echo ""
	@echo "Language Selection:"
	@echo "  make build LANGUAGE=python"
	@echo "  make build LANGUAGE=go"
	@echo "  make build LANGUAGE=java"
	@echo "  make build LANGUAGE=cpp"
	@echo ""
	@echo "Examples:"
	@echo "  make test LANGUAGE=python"
	@echo "  make run LANGUAGE=python"
