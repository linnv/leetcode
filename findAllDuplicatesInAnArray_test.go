package demo

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{2, 3}},
		{"normal", args{[]int{2, 2, 1}}, []int{2}},
		{"normal", args{[]int{3, 3, 2, 2, 1}}, []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
