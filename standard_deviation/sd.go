package standarddeviation

import (
	"math"

	"github.com/kh3rld/guess-it-1/variance"
)

// Calculate the standard deviation
func StandardD(data []float64) float64 {
	v := variance.Variance(data)
	return math.Sqrt(v)
}
