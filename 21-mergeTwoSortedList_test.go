package demo

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"edge", args{newListNode(1), nil}, newListNodeWithNext(1, nil)},
		{"edge", args{nil, newListNode(1)}, newListNodeWithNext(1, nil)},
		{"normal", args{newListNode(1), newListNode(2)}, newListNodeWithNext(1, newListNode(2))},
		{"normal", args{newListNodeWithNext(1, newListNode(2)), newListNodeWithNext(1, newListNode(2))}, newListNodeWithNext(1, newListNodeWithNext(1, newListNodeWithNext(2, newListNode(2))))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
