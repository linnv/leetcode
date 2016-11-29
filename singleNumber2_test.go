package demo

import "testing"

func Test_singleNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"edge", args{[]int{-2147483648}}, -2147483648},
		{"normal", args{[]int{1, 2, 2, 3, 3, 3, 2}}, 1},
		{"edge", args{[]int{1}}, 1},
		{"edge", args{[]int{-1}}, -1},
		{"edge", args{[]int{-2, -2, 1, 1, -3, 1, -3, -3, 4, -2}}, 4},
		{"edge", args{[]int{-401451, -177656, -2147483646, -473874, -814645, -2147483646, -852036, -457533, -401451, -473874, -401451, -216555, -917279, -457533, -852036, -457533, -177656, -2147483646, -177656, -917279, -473874, -852036, -917279, -216555, -814645, 2147483645, -2147483648, 2147483645, -814645, 2147483645, -216555}}, -2147483648},
		{"edge", args{[]int{-19, -46, -19, -46, -9, -9, -19, 17, 17, 17, -13, -13, -9, -13, -46, -28}}, -28},
		{"edge", args{[]int{-1, 2, 2, 3, 3, 3, 2}}, -1},
		{"edge", args{[]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}}, -4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber2(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}
