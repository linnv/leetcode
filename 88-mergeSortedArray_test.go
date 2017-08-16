package demo

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{[]int{1}, []int{}}, []int{1}},
		{"normal", args{[]int{1, 1}, []int{}}, []int{1, 1}},
		{"normal", args{[]int{}, []int{1}}, []int{1}},
		{"normal", args{[]int{2}, []int{1}}, []int{1, 2}},
		{"normal", args{[]int{1}, []int{2}}, []int{1, 2}},
		{"normal", args{[]int{1, 2}, []int{3}}, []int{1, 2, 3}},
		{"normal", args{[]int{1, 3}, []int{2, 4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"normal", args{[]int{1, 3}, []int{2, 3, 4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedArray(tt.args.num1, tt.args.num2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergedSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
