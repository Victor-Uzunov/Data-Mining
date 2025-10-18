# Task-specific Makefile for frog-leap-puzzle
# This file provides custom overrides for the generic system

# Variables
N ?= 2

# Custom build - uses root-level go.mod
build:
	@echo "Building frog-leap solution..."
	go build -o go/solution ./go

# Custom run - uses task-specific execution
run: build
	@echo $(N) | ./go/solution

# Custom test - uses enhanced fmi-ai-judge with timeout
test: build
	@echo "Running fmi-ai-judge tests..."
	fmi-ai-judge run go/solution --timeout 30

# Custom benchmark with enhanced options
bench: build
	@echo "Running benchmark tests..."
	fmi-ai-judge run go/solution --benchmark --timeout 60

# Custom validation
validate: build
	@echo "Validating solution format..."
	fmi-ai-judge validate go/solution

# Custom format
fmt:
	@echo "Formatting Go code..."
	gofmt -w go/*.go
	@which goimports >/dev/null 2>&1 && goimports -w go/*.go || echo "goimports not found, using gofmt only"

# Custom clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f go/solution

.PHONY: build run test bench validate fmt clean
