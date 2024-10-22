package diff

import "testing"

func TestRange(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{name: "Decimals", args: args{data: []float64{1, 2, 3, 4}}, want: 0.2639320225002102, want1: 4.73606797749979},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Range(tt.args.data)
			if got != tt.want {
				t.Errorf("Range() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Range() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
