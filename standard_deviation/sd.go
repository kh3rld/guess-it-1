package standarddeviation

import (
	"math"

	"github.com/kh3rld/guess-it-1/mean"
	"github.com/kh3rld/guess-it-1/variance"
)

// Calculate the standard deviation
func StandardD(data []float64) float64 {
	m := mean.Mean(data)
	return math.Sqrt(variance.Variance(data, m))
}
