package mean

import (
	"testing"
)

// TestMean tests the Mean function.
func TestMean(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name     string
		args     args
		expected float64
	}{
		{name: "Decimals", args: args{x: []float64{1, 2, 3, 4}}, expected: 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.x); got != tt.expected {
				t.Errorf("Mean() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
