package demo

import (
	"reflect"
	"testing"
)

func Test_removeLinkedListElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	var emptySlice []int
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{newListNodeWithNext(1, newListNode(2)), 1}, []int{2}},
		{"normal", args{newListNodeWithNext(1, newListNode(1)), 1}, emptySlice},
		{"normal", args{newListNodeWithNext(1, newListNodeWithNext(1, newListNode(2))), 1}, []int{2}},
		{"normal", args{newListNodeWithNext(1, newListNodeWithNext(1, newListNodeWithNext(2, newListNode(1)))), 1}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLinkedListElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("removeLinkedListElements() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
