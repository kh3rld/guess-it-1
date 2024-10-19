package variance

import "github.com/kh3rld/guess-it-1/mean"

// Variance is the expected value of the squared deviation from the mean of a random variable.
func Variance(x []float64) float64 {
	m := mean.Mean(x)
	r := []float64{}
	for _, x1 := range x {
		y := x1 - m
		r = append(r, y*y)
	}
	return mean.Mean(r)
}
