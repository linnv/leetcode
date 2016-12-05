package demo

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{1}, 1},
		{"normal", args{2}, 2},
		{"normal", args{6}, 6},
		{"normal", args{7}, 8},
		{"normal", args{8}, 9},
		{"normal", args{9}, 10},
		{"normal", args{10}, 12},
		{"normal", args{11}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumber(%d) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
