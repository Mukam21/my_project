.PHONY: fmt lint test

fmt:
	@echo "🔧 Formatting code..."
	gofmt -w .
	goimports -w .

lint: fmt
	@echo "🔍 Running linter..."
	golangci-lint run ./...

test:
	@echo "🧪 Running tests..."
	go test ./... -count=1 -v -cover

all: lint test
	@echo "✅ All checks passed!"