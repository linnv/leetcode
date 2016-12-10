package demo

import "testing"

func Test_sumOfLeftLeaves(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	n1 := newTreeNodeChilds(1, newTreeNode(2), newTreeNode(3))
	n2 := newTreeNodeChilds(1, newTreeNodeChilds(2, newTreeNode(3), newTreeNode(3)), newTreeNode(3))
	n3 := newTreeNodeChilds(1, newTreeNodeChilds(2, newTreeNode(3), newTreeNode(3)), newTreeNodeChilds(2, newTreeNode(3), newTreeNode(3)))
	n4 := newTreeNodeChilds(1, nil, newTreeNodeChilds(2, newTreeNode(3), newTreeNode(3)))
	n5 := newTreeNodeChilds(1, nil, newTreeNodeChilds(2, nil, newTreeNode(3)))
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{root: n1}, 2},
		{"normal", args{root: n2}, 3},
		{"normal", args{root: n3}, 6},
		{"normal", args{root: n4}, 3},
		{"edge", args{root: n5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
