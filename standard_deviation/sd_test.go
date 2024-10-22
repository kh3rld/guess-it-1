package standarddeviation

import "testing"

func TestStandardD(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Decimals", args: args{data: []float64{1, 2, 3, 4}}, want: 1.118033988749895},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardD(tt.args.data); got != tt.want {
				t.Errorf("StandardD() = %v, want %v", got, tt.want)
			}
		})
	}
}
