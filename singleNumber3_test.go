package demo

import (
	"reflect"
	"testing"
)

func Test_singleNumber3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{[]int{1, 2, 3, 3, 2, 4}}, []int{1, 4}},
		{"normal", args{[]int{1, 2, 2, 3, 3, 4}}, []int{1, 4}},
		{"normal", args{[]int{1, 2, 2, 4}}, []int{1, 4}},
		{"edge", args{[]int{-1, 2, 2, 4}}, []int{-1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber3(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber3() = %v, want %v", got, tt.want)
			}
		})
	}
}
