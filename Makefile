# Generic Makefile for algorithms-solutions repository
# Supports multiple tasks with multiple programming languages

# Default values
TASK ?=
N ?= 2
LANGUAGE ?= auto
GO_VERSION ?= 1.21

# Colors for output
GREEN := \033[0;32m
YELLOW := \033[1;33m
RED := \033[0;31m
NC := \033[0m # No Color

# Auto-detect tasks (directories containing language subdirectories)
TASKS := $(shell find . -maxdepth 2 -type d \( -name "go" -o -name "python" -o -name "java" -o -name "cpp" \) | sed 's|^\./||g' | sed 's|/[^/]*$$||g' | sort -u)

# Detect available languages for a task
define detect_languages
$(shell find $(1) -maxdepth 1 -type d \( -name "go" -o -name "python" -o -name "java" -o -name "cpp" \) | sed 's|.*/||g' | sort)
endef

# Default target
.DEFAULT_GOAL := help

help:
	@echo "$(GREEN)Algorithms Solutions - Generic Makefile$(NC)"
	@echo ""
	@echo "$(YELLOW)Available targets:$(NC)"
	@echo "  help              - Show this help message"
	@echo "  list-tasks        - List all available tasks"
	@echo "  build TASK=<name> [LANGUAGE=<lang>] - Build specific task"
	@echo "  run TASK=<name> [LANGUAGE=<lang>] N=<input> - Run specific task"
	@echo "  test TASK=<name> [LANGUAGE=<lang>] - Test specific task with fmi-ai-judge"
	@echo "  bench TASK=<name> [LANGUAGE=<lang>] - Benchmark specific task"
	@echo "  validate TASK=<name> [LANGUAGE=<lang>] - Validate task format without running tests"
	@echo "  fmt TASK=<name> [LANGUAGE=<lang>] - Format code for specific task"
	@echo "  clean TASK=<name> [LANGUAGE=<lang>] - Clean build artifacts"
	@echo "  build-all         - Build all tasks"
	@echo "  test-all          - Test all tasks"
	@echo "  fmt-all           - Format all tasks"
	@echo "  clean-all         - Clean all tasks"
	@echo ""
	@echo "$(YELLOW)Supported languages:$(NC)"
	@echo "  go, python, java, cpp"
	@echo "  Use LANGUAGE=auto (default) to auto-detect or specify explicitly"
	@echo ""
	@echo "$(YELLOW)Examples:$(NC)"
	@echo "  make run TASK=frog-leap-puzzle N=3 LANGUAGE=go"
	@echo "  make test TASK=frog-leap-puzzle LANGUAGE=python"
	@echo "  make validate TASK=frog-leap-puzzle"
	@echo "  make build-all"
	@echo ""
	@echo "$(YELLOW)Available tasks:$(NC)"
	@for task in $(TASKS); do \
		echo -n "  - $$task ("; \
		langs=$$(find $$task -maxdepth 1 -type d \( -name "go" -o -name "python" -o -name "java" -o -name "cpp" \) | sed 's|.*/||g' | sort | tr '\n' ',' | sed 's/,$$//'); \
		echo "$$langs)"; \
	done

list-tasks:
	@echo "$(GREEN)Available tasks:$(NC)"
	@for task in $(TASKS); do \
		if [ -d "$$task" ]; then \
			langs=$$(find $$task -maxdepth 1 -type d \( -name "go" -o -name "python" -o -name "java" -o -name "cpp" \) 2>/dev/null | sed 's|.*/||g' | sort | tr '\n' ' ' | sed 's/ *$$//'); \
			if [ -n "$$langs" ]; then \
				echo "  - $$task ($$langs)"; \
			fi; \
		fi; \
	done

# Validate task parameter and detect language
check-task:
	@if [ -z "$(TASK)" ]; then \
		echo "$(RED)Error: TASK parameter is required$(NC)"; \
		echo "Usage: make <target> TASK=<task-name> [LANGUAGE=<language>]"; \
		echo "Available tasks: $(TASKS)"; \
		exit 1; \
	fi
	@if [ ! -d "$(TASK)" ]; then \
		echo "$(RED)Error: Task '$(TASK)' not found$(NC)"; \
		echo "Available tasks: $(TASKS)"; \
		exit 1; \
	fi

# Detect the language to use for a task
detect-lang: check-task
	@if [ "$(LANGUAGE)" = "auto" ]; then \
		if [ -d "$(TASK)/go" ]; then \
			echo "go"; \
		elif [ -d "$(TASK)/python" ]; then \
			echo "python"; \
		elif [ -d "$(TASK)/java" ]; then \
			echo "java"; \
		elif [ -d "$(TASK)/cpp" ]; then \
			echo "cpp"; \
		else \
			echo "$(RED)Error: No supported language implementation found for task '$(TASK)'$(NC)" >&2; \
			exit 1; \
		fi \
	else \
		if [ ! -d "$(TASK)/$(LANGUAGE)" ]; then \
			echo "$(RED)Error: Language '$(LANGUAGE)' not available for task '$(TASK)'$(NC)" >&2; \
			exit 1; \
		fi; \
		echo "$(LANGUAGE)"; \
	fi

# Build specific task
build: check-task
	@DETECTED_LANG=$$(LANGUAGE=C $(MAKE) -s detect-lang TASK=$(TASK) LANGUAGE=$(LANGUAGE)); \
	echo "$(GREEN)Building task: $(TASK) ($$DETECTED_LANG)$(NC)"; \
	if [ -f "$(TASK)/Makefile" ]; then \
		cd $(TASK) && $(MAKE) build LANGUAGE=$$DETECTED_LANG; \
	else \
		case $$DETECTED_LANG in \
			go) go build -o $(TASK)/go/solution ./$(TASK)/go ;; \
			python) echo "Python doesn't need compilation" ;; \
			java) cd $(TASK)/java && javac Solution.java ;; \
			cpp) cd $(TASK)/cpp && g++ -o solution solution.cpp ;; \
			*) echo "$(RED)Unsupported language: $$DETECTED_LANG$(NC)" && exit 1 ;; \
		esac \
	fi

# Run specific task
run: check-task build
	@DETECTED_LANG=$$(LANGUAGE=C $(MAKE) -s detect-lang TASK=$(TASK) LANGUAGE=$(LANGUAGE)); \
	echo "$(GREEN)Running task: $(TASK) ($$DETECTED_LANG) with N=$(N)$(NC)"; \
	if [ -f "$(TASK)/Makefile" ]; then \
		cd $(TASK) && $(MAKE) run N=$(N) LANGUAGE=$$DETECTED_LANG; \
	else \
		case $$DETECTED_LANG in \
			go) echo $(N) | ./$(TASK)/go/solution ;; \
			python) echo $(N) | python3 $(TASK)/python/solution.py ;; \
			java) echo $(N) | cd $(TASK)/java && java Solution ;; \
			cpp) echo $(N) | ./$(TASK)/cpp/solution ;; \
			*) echo "$(RED)Unsupported language: $$DETECTED_LANG$(NC)" && exit 1 ;; \
		esac \
	fi

# Test specific task with enhanced judge options
test: check-task build
	@DETECTED_LANG=$$(LANGUAGE=C $(MAKE) -s detect-lang TASK=$(TASK) LANGUAGE=$(LANGUAGE)); \
	echo "$(GREEN)Testing task: $(TASK) ($$DETECTED_LANG)$(NC)"; \
	if [ -f "$(TASK)/Makefile" ]; then \
		cd $(TASK) && $(MAKE) test LANGUAGE=$$DETECTED_LANG; \
	else \
		case $$DETECTED_LANG in \
			go) fmi-ai-judge run $(TASK)/go/solution --timeout 30 ;; \
			python) fmi-ai-judge run "python3 $(TASK)/python/solution.py" --timeout 30 ;; \
			java) cd $(TASK)/java && fmi-ai-judge run "java Solution" --timeout 30 ;; \
			cpp) fmi-ai-judge run $(TASK)/cpp/solution --timeout 30 ;; \
			*) echo "$(RED)Unsupported language: $$DETECTED_LANG$(NC)" && exit 1 ;; \
		esac \
	fi

# Benchmark specific task with enhanced options
bench: check-task build
	@DETECTED_LANG=$$(LANGUAGE=C $(MAKE) -s detect-lang TASK=$(TASK) LANGUAGE=$(LANGUAGE)); \
	echo "$(GREEN)Benchmarking task: $(TASK) ($$DETECTED_LANG)$(NC)"; \
	if [ -f "$(TASK)/Makefile" ]; then \
		cd $(TASK) && $(MAKE) bench LANGUAGE=$$DETECTED_LANG; \
	else \
		case $$DETECTED_LANG in \
			go) fmi-ai-judge run $(TASK)/go/solution --benchmark --timeout 60 ;; \
			python) fmi-ai-judge run "python3 $(TASK)/python/solution.py" --benchmark --timeout 60 ;; \
			java) cd $(TASK)/java && fmi-ai-judge run "java Solution" --benchmark --timeout 60 ;; \
			cpp) fmi-ai-judge run $(TASK)/cpp/solution --benchmark --timeout 60 ;; \
			*) echo "$(RED)Unsupported language: $$DETECTED_LANG$(NC)" && exit 1 ;; \
		esac \
	fi

# Validate specific task (check format without running tests)
validate: check-task build
	@DETECTED_LANG=$$(LANGUAGE=C $(MAKE) -s detect-lang TASK=$(TASK) LANGUAGE=$(LANGUAGE)); \
	echo "$(GREEN)Validating task: $(TASK) ($$DETECTED_LANG)$(NC)"; \
	case $$DETECTED_LANG in \
		go) fmi-ai-judge validate $(TASK)/go/solution ;; \
		python) fmi-ai-judge validate "python3 $(TASK)/python/solution.py" ;; \
		java) cd $(TASK)/java && fmi-ai-judge validate "java Solution" ;; \
		cpp) fmi-ai-judge validate $(TASK)/cpp/solution ;; \
		*) echo "$(RED)Unsupported language: $$DETECTED_LANG$(NC)" && exit 1 ;; \
	esac

# Format specific task code
fmt: check-task
	@DETECTED_LANG=$$(LANGUAGE=C $(MAKE) -s detect-lang TASK=$(TASK) LANGUAGE=$(LANGUAGE)); \
	echo "$(GREEN)Formatting task: $(TASK) ($$DETECTED_LANG)$(NC)"; \
	if [ -f "$(TASK)/Makefile" ]; then \
		cd $(TASK) && $(MAKE) fmt LANGUAGE=$$DETECTED_LANG; \
	else \
		case $$DETECTED_LANG in \
			go) gofmt -w $(TASK)/go/solution.go ;; \
			python) autopep8 --in-place --aggressive --aggressive $(TASK)/python/solution.py ;; \
			java) cd $(TASK)/java && java -jar /path/to/google-java-format.jar --replace Solution.java ;; \
			cpp) clang-format -i $(TASK)/cpp/solution.cpp ;; \
			*) echo "$(RED)Unsupported language: $$DETECTED_LANG$(NC)" && exit 1 ;; \
		esac \
	fi

# Clean build artifacts for a specific task
clean: check-task
	@DETECTED_LANG=$$(LANGUAGE=C $(MAKE) -s detect-lang TASK=$(TASK) LANGUAGE=$(LANGUAGE)); \
	echo "$(GREEN)Cleaning task: $(TASK) ($$DETECTED_LANG)$(NC)"; \
	if [ -f "$(TASK)/Makefile" ]; then \
		cd $(TASK) && $(MAKE) clean LANGUAGE=$$DETECTED_LANG; \
	else \
		case $$DETECTED_LANG in \
			go) rm -f $(TASK)/go/solution ;; \
			python) echo "No build artifacts to clean for Python" ;; \
			java) cd $(TASK)/java && rm -f *.class ;; \
			cpp) cd $(TASK)/cpp && rm -f solution ;; \
			*) echo "$(RED)Unsupported language: $$DETECTED_LANG$(NC)" && exit 1 ;; \
		esac \
	fi

# Build all tasks
build-all:
	@echo "$(GREEN)Building all tasks...$(NC)"
	@for task in $(TASKS); do \
		$(MAKE) build TASK=$$task LANGUAGE=auto || exit 1; \
	done
	@echo "$(GREEN)All tasks built successfully!$(NC)"

# Test all tasks
test-all:
	@echo "$(GREEN)Testing all tasks...$(NC)"
	@for task in $(TASKS); do \
		$(MAKE) test TASK=$$task LANGUAGE=auto || exit 1; \
	done
	@echo "$(GREEN)All tasks tested successfully!$(NC)"

# Format all tasks
fmt-all:
	@echo "$(GREEN)Formatting all tasks...$(NC)"
	@for task in $(TASKS); do \
		$(MAKE) fmt TASK=$$task LANGUAGE=auto || exit 1; \
	done
	@echo "$(GREEN)All tasks formatted successfully!$(NC)"

# Clean all tasks
clean-all:
	@echo "$(GREEN)Cleaning all tasks...$(NC)"
	@for task in $(TASKS); do \
		$(MAKE) clean TASK=$$task LANGUAGE=auto || exit 1; \
	done
	@echo "$(GREEN)All tasks cleaned successfully!$(NC)"

# Install dependencies for all tasks
deps:
	@echo "$(GREEN)Installing dependencies for all tasks...$(NC)"
	@for task in $(TASKS); do \
		if [ -d "$$task/go" ]; then \
			cd $$task/go && go mod tidy && cd - || exit 1; \
		fi; \
		if [ -d "$$task/python" ]; then \
			cd $$task/python && pip install -r requirements.txt --quiet || exit 1; \
		fi; \
		if [ -d "$$task/java" ]; then \
			cd $$task/java && ./gradlew build -x test --quiet || exit 1; \
		fi; \
		if [ -d "$$task/cpp" ]; then \
			cd $$task/cpp && make deps --quiet || exit 1; \
		fi; \
	done
	@echo "$(GREEN)All task dependencies installed!$(NC)"

# Initialize task (create directories, files, etc. as needed)
init-task:
	@if [ -z "$(TASK)" ]; then \
		echo "$(RED)Error: TASK parameter is required for init-task$(NC)"; \
		exit 1; \
	fi
	@if [ -d "$(TASK)" ]; then \
		echo "$(YELLOW)Warning: Task '$(TASK)' already exists. Skipping creation.$(NC)"; \
	else \
		mkdir -p $(TASK)/{go,python,java,cpp} && \
		echo "package main" > $(TASK)/go/solution.go && \
		echo "print('Hello, World!')" > $(TASK)/python/solution.py && \
		echo "public class Solution { public static void main(String[] args) { System.out.println(\"Hello, World!\"); }}" > $(TASK)/java/Solution.java && \
		echo "#include <iostream>" > $(TASK)/cpp/solution.cpp && \
		echo "using namespace std; int main() { cout << \"Hello, World!\" << endl; return 0; }" >> $(TASK)/cpp/solution.cpp && \
		echo "$(GREEN)Task '$(TASK)' initialized with sample solutions.$(NC)"; \
	fi

.PHONY: help list-tasks check-task detect-lang build run test bench validate fmt clean build-all test-all fmt-all clean-all deps init-task
