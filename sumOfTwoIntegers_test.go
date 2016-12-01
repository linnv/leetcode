package demo

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"edge", args{4, -1}, 3},
		{"edge", args{1, 0}, 1},
		{"edge", args{-1, 0}, -1},
		{"edge", args{1, -1}, 0},
		{"normal", args{1, 2}, 3},
		{"normal", args{33333, 2}, 33335},
		{"normal", args{33333, -2}, 33331},
		{"normal", args{-1, -1}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
