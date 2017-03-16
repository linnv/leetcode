package demo

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{newTreeNodewithChilds(1, newTreeNode(3), newTreeNode(3)), newTreeNodewithChilds(1, newTreeNode(3), newTreeNode(3))}, true},
		{"normal", args{newTreeNodewithChilds(1, newTreeNode(4), newTreeNode(3)), newTreeNodewithChilds(1, newTreeNode(3), newTreeNode(3))}, false},
		{"normal", args{newTreeNode(3), newTreeNode(4)}, false},
		{"normal", args{newTreeNode(3), newTreeNode(3)}, true},
		{"edge", args{nil, newTreeNode(3)}, false},
		{"edge", args{nil, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
