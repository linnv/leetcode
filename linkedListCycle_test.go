package demo

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		ListNode *ListNode
	}
	one := newListNodeWithNext(21, newListNode(22))
	two := newListNode(23)
	root := newListNode(1)
	two.Next = root
	root.Next = two
	one.Next = root

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{one}, true},
		{"normal", args{newListNodeWithNext(1, newListNode(2))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.ListNode); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
