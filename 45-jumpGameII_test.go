package demo

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"edge", args{[]int{3}}, 0},
		{"edge", args{[]int{3, 1}}, 1},
		{"edge", args{[]int{3, 1, 1, 1}}, 1},
		{"normal", args{[]int{2, 3, 2, 3, 2, 3, 2, 3, 2, 3}}, 5},
		{"normal", args{[]int{2, 3, 4, 2, 5, 2, 3, 5, 1, 2, 3, 4}}, 4},
		{"normal", args{[]int{1, 2, 3, 4, 5, 6, 7}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.num); got != tt.want {
				t.Errorf("jump(%v) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}
