package demo

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{100000}, 1},
		{"normal", args{-100}, -1},
		{"normal", args{0}, 0},
		{"normal", args{-0}, 0},
		{"normal", args{-1}, -1},
		{"normal", args{12}, 21},
		{"normal", args{-12}, -21},
		{"normal", args{-123}, -321},
		{"normal", args{-1234}, -4321},
		{"normal", args{1234}, 4321},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				// if got := reverseByswapChart(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
