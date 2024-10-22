# Guess IT

This Go project provides essential functions for statistical analysis, including calculating the **mean**, **variance**, **standard deviation**, and **range** of a dataset. It is designed to be simple and efficient, making it suitable for data analysis tasks in Go applications.

## Features

- Calculate statistical metrics with ease.
- Handles edge cases such as empty slices.
- Robust error checking to prevent overflow.

## Installation

To install this package, clone the repository and use `go get` to fetch it:

```bash
git clone https://github.com/kh3rld/guess-it-1.git
cd guess-it-1
```

## Usage

Here’s an example of how to use the functions in this project:

- create a folder called student
- copy the files that are needed to run your program into this folder
- write an executable shell script named script.sh containing the command(s) to run your program, from the root folder of the provided tester (see below where to find it)

## Functions

### Mean

```go
func Mean(data []float64) float64
```

Calculates the average of a slice of `float64` numbers.

**Parameters:**

- `data`: A slice of `float64` numbers.

**Returns:**

- The mean as a `float64`. Returns `0` for an empty slice.

### Variance

```go
func Variance(x []float64) float64
```

Calculates the variance of a slice of `float64` numbers.

**Parameters:**

- `x`: A slice of `float64` numbers.

**Returns:**

- The variance as a `float64`. Returns `0` for an empty slice.

### StandardD

```go
func StandardD(data []float64) float64
```

Calculates the standard deviation of a slice of `float64` numbers.

**Parameters:**

- `data`: A slice of `float64` numbers.

**Returns:**

- The standard deviation as a `float64`. Returns `0` for an empty slice.

### Range

```go
func Range(data []float64) (float64, float64)
```

Calculates the range of values based on the mean and standard deviation.

**Parameters:**

- `data`: A slice of `float64` numbers.

**Returns:**

- Two `float64` values representing the lower and upper bounds.

## Testing

To run the tests for this package, execute the following command:

```bash
go test ./...
```

This will run all the test cases defined in the package.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request. Ensure that your code adheres to the existing style and includes tests where applicable.

