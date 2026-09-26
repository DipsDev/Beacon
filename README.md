# Beacon

A lightweight, robust black-box end-to-end testing framework built in Go.

Beacon allows you to embed test specifications directly inside source files using comment directives, making it effortless to test compilers, interpreters, virtual machines, and CLI utilities.

## Writing a Test (`.beacon`)

Beacon files are executable source files containing inline test directives. Here is an example testing array creation, indexing, and membership:
```python
numbers = [10, 20, 30]
print(numbers[0])
#beacon: expect 10

print(numbers[1])
#beacon: expect 20

has_thirty = 30 in numbers
print(has_thirty)
#beacon: expect True

#beacon: exit 0
```

## Getting Started
Place your .beacon test files in your project's test directory (e.g., tests/).

Run your test runner pointing to your target executable:

```bash
beacon run --executable ./bin/my_language --dir ./tests
```

Watch the output and fix your code!
```
running beacon tests...
----------------------------------------
✕ test_arrays.beacon (9)
   └─ should be able to push to arrays > expect 30
      Expected: 30
      Got: 33

   └─ should be able to pop from an array > expect [20,30,33]
      Expected: [20,30,33]
      Got: [10,20,30]


✓ test_branching.beacon (3)
✓ test_math.beacon (5)
✕ test_recursion.beacon (1)
   └─ unnamed_group > expect 10
      Expected: 10
      Got: runtime error: variable 'count' is not defined


✓ test_while.beacon (1)
==================================================
Test Files:  3 passed, 5 total
Test Suites: 4 passed, 7 total
Tests:       6 failed, 13 passed, 19 total
==================================================
```
