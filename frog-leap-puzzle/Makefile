# Custom build - uses root-level go.mod
build:
	@echo "Building frog-leap solution..."
	go build -o go/frog-leap ./go

# Custom test - uses enhanced judge with time=out
test: build
	@echo "Running fmi-ai-judge tests..."
	judge run --bench go/frog-leap

# Custom clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f go/frog-leap

.PHONY: build run test bench validate fmt clean
