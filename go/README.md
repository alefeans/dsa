# Go Solutions

[![Go](https://img.shields.io/badge/go-1.19+-blue.svg)]()

Go implementations of data structures, algorithms, and solutions to various coding challenges.

## Getting Started

### Prerequisites

- Go 1.19 or higher

### Installation

```bash
cd go/
go mod download  # No external dependencies, but ensures module is initialized
```

### Running Solutions

All solutions use only the Go standard library:

```bash
# Run a specific file
go run hackerrank/30_days_challenge/day_1/data_types.go

# Run gophercises quiz game (from project root)
go run go/gophercises/quiz_game/quiz.go

# Build and run
go build -o quiz gophercises/quiz_game/quiz.go
./quiz
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests in a specific directory
go test ./hackerrank/30_days_challenge/day_2/

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```