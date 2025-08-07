# Python Solutions

[![codecov](https://codecov.io/gh/alefeans/python-exercises/branch/master/graph/badge.svg)](https://codecov.io/gh/alefeans/python-exercises) [![Python](https://img.shields.io/badge/python-3.11+-blue.svg)]()

Python implementations of data structures, algorithms, and solutions to various coding challenges.

## Getting Started

### Prerequisites

- Python 3.11 or higher
- Poetry (for dependency management) or pip

### Installation

```bash
# Using Poetry (recommended)
cd python/
poetry install

# Or using pip
cd python/
pip install .
```

### Running Solutions

All solutions use only the Python standard library (dependencies are for testing only):

```bash
python {folder}/{challenge}/{challenge}.py

# Examples:
python dsa/binary_search.py
python hacker_rank/30_days_challenge/day_1/data_types.py
python project_euler/problem_1.py
```

### Running Tests

```bash
# Run all tests
pytest

# Run with coverage
pytest --cov

# Watch mode for development
pytest-watch
```


