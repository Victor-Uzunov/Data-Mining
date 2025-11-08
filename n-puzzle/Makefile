.PHONY: build test clean

# Default language detection
LANGUAGE ?= go

build:
	@echo "Building n-puzzle solution..."
	@case $(LANGUAGE) in \
		go) go build -o go/n-puzzle ./go ;; \
		python) echo "Python doesn't need compilation" ;; \
		java) cd java && javac *.java ;; \
		cpp) cd cpp && g++ -o n-puzzle *.cpp ;; \
		*) echo "Unsupported language: $(LANGUAGE)" && exit 1 ;; \
	esac

test: build
	@echo "Running fmi-ai-judge tests..."
	@case $(LANGUAGE) in \
		go) \
			if command -v judge >/dev/null 2>&1; then \
				judge run --bench go/n-puzzle; \
			elif [ -f "../.venv/bin/judge" ]; then \
				../.venv/bin/judge run --bench go/n-puzzle; \
			elif [ -f ".venv/bin/judge" ]; then \
				.venv/bin/judge run --bench go/n-puzzle; \
			else \
				echo "Error: fmi-ai-judge not found. Install with: pip install fmi-ai-judge"; \
				exit 1; \
			fi ;; \
		python) \
			if command -v judge >/dev/null 2>&1; then \
				judge run --bench python/solution.py; \
			elif [ -f "../.venv/bin/judge" ]; then \
				../.venv/bin/judge run --bench python/solution.py; \
			elif [ -f ".venv/bin/judge" ]; then \
				.venv/bin/judge run --bench python/solution.py; \
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
				judge run --bench cpp/n-puzzle; \
			elif [ -f "../.venv/bin/judge" ]; then \
				../.venv/bin/judge run --bench cpp/n-puzzle; \
			elif [ -f ".venv/bin/judge" ]; then \
				.venv/bin/judge run --bench cpp/n-puzzle; \
			else \
				echo "Error: fmi-ai-judge not found. Install with: pip install fmi-ai-judge"; \
				exit 1; \
			fi ;; \
		*) echo "Unsupported language: $(LANGUAGE)" && exit 1 ;; \
	esac

clean:
	@echo "Cleaning build artifacts..."
	@rm -f go/n-puzzle
	@rm -f java/*.class
	@rm -f cpp/n-puzzle
	@rm -rf .judge
