package demo

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{[]int{1, 2, 3, 3}}, 3},
		{"normal", args{[]int{1, 2, 1}}, 1},
		{"normal", args{[]int{2, 1, 1}}, 1},
		{"normal", args{[]int{1, 1, 2}}, 1},
		{"normal", args{[]int{1, 2, 3, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
