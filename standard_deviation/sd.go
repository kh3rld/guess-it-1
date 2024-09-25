package standarddeviation

import (
	"math"

	"github.com/kh3rld/guess-it-1/variance"
)

func StandardD(x []float64, m float64) float64 {
	return math.Sqrt(variance.Variance(x, m))
}
