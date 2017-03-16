package demo

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{newListNodeWithNext(1, newListNodeWithNext(1, newListNodeWithNext(2, newListNodeWithNext(3, newListNodeWithNext(3, nil)))))}, []int{1, 2, 3}},
		{"normal", args{newListNodeWithNext(1, newListNodeWithNext(1, newListNode(2)))}, []int{1, 2}},
		{"normal", args{newListNodeWithNext(1, newListNode(2))}, []int{1, 2}},
		{"normal", args{newListNode(1)}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
