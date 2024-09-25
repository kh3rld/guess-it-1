package diff

import (
	"github.com/kh3rld/guess-it-1/mean"
	sd "github.com/kh3rld/guess-it-1/standard_deviation"
)

func Range(x []float64) (float64, float64) {
	mean := mean.Mean(x)
	sd := sd.StandardD(x)
	// upper limit
	upper := mean + sd
	// lower limit
	lower := mean - sd

	return lower, upper
}
