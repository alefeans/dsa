# Clojure Solutions

[![Clojure](https://img.shields.io/badge/clojure-1.11+-blue.svg)]()

Clojure implementations of data structures, algorithms, and solutions to various coding challenges.

## Getting Started

### Prerequisites

- Java 8 or higher
- Leiningen (Clojure build tool)

### Installation

```bash
cd clojure/
lein deps  # Download dependencies defined in project.clj
```

### Running Solutions

All solutions use Clojure core libraries:

```bash
# Run a specific namespace
lein run -m exercises.hackerrank.functional-programming.hello-word-n-times

# Start a REPL
lein repl

# Examples:
lein run -m exercises.clojure-for-the-brave-and-true.chapter-3.exercise-4
lein run -m exercises.hackerrank.functional-programming.filter-array
```

### Running Tests

```bash
# Run all tests
lein test

# Run specific test namespace
lein test common-algorithms.missing-number-test

# Run tests with output
lein test :all
```