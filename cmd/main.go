package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kh3rld/guess-it-1/diff"
)

func main() {
	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
			continue
		}

		numbers = append(numbers, num)
		if len(numbers) > 1 {
			lower, upper := diff.Range(numbers[:len(numbers)-1])
			fmt.Printf("%.0f %.0f\n", lower, upper)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from input: %v\n", err)
	}
}
