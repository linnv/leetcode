package demo

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{[]int{0, 2}}, 1},
		{"edge", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
			if got := missingNumberByBitsOperation(tt.args.nums); got != tt.want {
				t.Errorf("missingNumberByBitsOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
