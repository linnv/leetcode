package demo

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"edge", args{[]int{2, 4}}, 1},
		{"edge", args{[]int{0}}, 1},
		{"edge", args{[]int{2}}, 1},
		{"edge", args{[]int{2, -2}}, 1},
		{"edge", args{[]int{-1}}, 1},
		{"edge", args{[]int{-1, -1}}, 1},
		{"edge", args{[]int{0, 2, 2, 4}}, 1},
		{"edge", args{[]int{2, 2, 2, 4}}, 1},
		{"edge", args{[]int{1, 2, 2, 2, 4}}, 3},
		{"edge", args{[]int{1, 2, 2, 4}}, 3},
		{"normal", args{[]int{0, 2, 4}}, 1},
		{"normal", args{[]int{1, 2, 0}}, 3},
		{"normal", args{[]int{1, 2, 4}}, 3},
		{"edge", args{[]int{-2, -5, 2, 3}}, 1},
		{"normal", args{[]int{3, 4, -1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
