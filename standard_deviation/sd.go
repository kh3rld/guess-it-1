package standarddeviation

import (
	"math"

	"github.com/kh3rld/guess-it-1/mean"
	"github.com/kh3rld/guess-it-1/variance"
)

func StandardD(x []float64) float64 {
	m := mean.Mean(x)
	return math.Sqrt(variance.Variance(x, m))
}
