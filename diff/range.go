package diff

import (
	"github.com/kh3rld/guess-it-1/mean"
	sd "github.com/kh3rld/guess-it-1/standard_deviation"
)

// Calculate the range based on mean and standard deviation
func Range(data []float64) (float64, float64) {
	if len(data) == 0 {
		return 0, 0
	}
	m := mean.Mean(data)
	sd := sd.StandardD(data)
	return m - 3*sd, m + 3*sd
}
