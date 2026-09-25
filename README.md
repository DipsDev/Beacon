# Beacon

A lightweight, robust black-box end-to-end testing framework built in Go.

Beacon allows you to embed test specifications directly inside source files using comment directives, making it effortless to test compilers, interpreters, virtual machines, and CLI utilities.

## Writing a Test (`.beacon`)

Beacon files are executable source files containing inline test directives. Here is an example testing array creation, indexing, and membership:
```python
numbers = [10, 20, 30];
print(numbers[0]);
//beacon: expect 10

print(numbers[1]);
//beacon: expect 20

let has_thirty = 30 in numbers;
print(has_thirty);
//beacon: expect true

//beacon: exit 0
```

## Getting Started
Place your .beacon test files in your project's test directory (e.g., tests/).

Run your test runner pointing to your target executable:

```bash
beacon run --executable ./bin/my_language --dir ./tests
```
