# Algorithm Benchmarking in Go and Python

## Overview
This repository contains benchmarking tests for various algorithms implemented in **Go** and **Python**. The benchmarks are designed to evaluate the performance, memory usage, and scalability of different algorithms, including sorting algorithms (e.g., Bubble Sort, Merge Sort) as input sizes increase.

## Features
- **Benchmarking Algorithms**: Compare the execution time of various algorithms like sorting algorithms (Bubble Sort, Merge Sort, etc.) in both Go and Python.
- **Memory Profiling**: Measure memory allocation and usage during execution using **Go's memory profiling** and **Python's tracemalloc**.
- **Scalability Analysis**: Analyze how algorithms behave with increasing input sizes, and visualize the results using **Matplotlib** in Python.

## Tools Used
- **Go Testing Package**: Used for running benchmarks and collecting performance metrics.
- **Go Memory Profiling**: `b.ReportAllocs()` for tracking memory allocations during benchmark runs.
- **Python's `timeit`**: For measuring execution time.
- **Python's `tracemalloc`**: For tracking memory usage.
- **Python's `matplotlib`**: For plotting scalability graphs.

## Running Benchmarks

### Go
To run the benchmarks for all algorithms in the Go package:

```bash
go test -bench .
```
To run a specific benchmark in Go:

```bash
go test -bench BenchmarkBubbleSort
```
Python
To run the benchmarks in Python:

```bash
python main.py
```

Make sure to install the required dependencies first:

```bash
pip install matplotlib
```
## Contribution
Feel free to contribute by adding more algorithms, improving the benchmark tests, or supporting additional languages. Please ensure all contributions are accompanied by benchmarks to validate performance.

## License
MIT License 