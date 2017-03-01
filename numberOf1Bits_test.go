package demo

import "testing"

func Test_hanmmingWeight(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{1}, 1},
		{"normal", args{2}, 1},
		{"normal", args{3}, 2},
		{"normal", args{4}, 1},
		{"normal", args{8}, 1},
		{"normal", args{10}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hanmmingWeight(tt.args.n); got != tt.want {
				t.Errorf("hanmmingWeight(%d) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
