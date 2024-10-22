package variance

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Decimals", args: args{x: []float64{1, 2, 3, 4}}, want: 1.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.x); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}
