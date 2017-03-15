package demo

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{[]int{1, 1, 0, 1, 1, 1}}, 3},
		{"edge", args{[]int{1, 1, 1, 1, 1}}, 5},
		{"edge", args{[]int{1, 1, 1, 1, 1, 0}}, 5},
		{"edge", args{[]int{1, 1, 1, 1, 1, 0, 1}}, 5},
		{"edge", args{[]int{1, 0, 1, 1, 1, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
