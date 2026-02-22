# CodingKatas-Golang

A Go implementation of common coding katas and algorithmic problems. This repository is designed for practicing and studying fundamental data structures and algorithms.

## Prerequisites

- **Go 1.25.1** or higher
- **Make** (optional, for using Makefile commands)
- **golangci-lint** (optional, for linting)

## Project Structure

```
.
├── arrays/              # Array manipulation and operations
├── backtracking/        # Backtracking algorithms (permutations, N-Queens, Sudoku)
├── binary_search/       # Binary search techniques and variations
├── bit_manipulation/    # Bit manipulation and bitwise operations
├── dynamic_programming/ # Dynamic programming problems (Fibonacci, coin change, etc.)
├── graphs/              # Graph algorithms and traversal
├── hash/                # Hash-based data structures (e.g., hash ring)
├── heap/                # Heap implementations (e.g., fixed-size heap)
├── intervals/           # Interval/range problems (merge, insert, scheduling)
├── lists/               # Linked list operations
├── matrix/              # Matrix operations (rotate, spiral, search)
├── roman/               # Roman numeral conversions
├── sorts/               # Sorting algorithms
├── stack_queue/         # Stack and Queue data structures
├── strings/             # String manipulation and algorithms
├── trees/               # Tree data structures (e.g., binary trees)
├── trie/                # Trie data structure (prefix tree)
├── two_pointers/        # Two pointer and sliding window techniques
├── union_find/          # Union-Find (Disjoint Set Union) data structure
├── cmd/                 # Command-line entry point
├── go.mod              # Go module definition
├── Makefile            # Build automation
└── README.md           # This file
```

## Getting Started

### Building the Project

```bash
make build
```

### Running the Application

```bash
make run
```

### Running Tests

Execute all tests with verbose output:

```bash
make test
```

To run tests for a specific module:

```bash
go test -v ./arrays
go test -v ./graphs
go test -v ./trees
# ... etc
```

### Linting

Check code quality with golangci-lint:

```bash
make lint
```

### Cleaning Up

Remove built binaries and clean the build cache:

```bash
make clean
```

## Testing

This project uses [testify](https://github.com/stretchr/testify) for assertions and testing utilities.

## License

See [LICENSE](LICENSE) file for details.
