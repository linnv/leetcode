package demo

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{[]int{1, 2, 4, 1}}, true},
		{"normal", args{[]int{1, 1}}, true},
		{"normal", args{[]int{2, 1, 1, 3}}, true},
		{"normal", args{[]int{2, 9, 1, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
