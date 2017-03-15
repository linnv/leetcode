package demo

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{newTreeNodewithChilds(1, newTreeNode(3), newTreeNode(3))}, true},
		{"normal", args{newTreeNodewithChilds(1, newTreeNode(3), nil)}, false},
		{"normal", args{newTreeNode(3)}, true},
		{"normal", args{newTreeNodewithChilds(1,
			newTreeNodewithChilds(2, newTreeNode(3), nil),
			newTreeNodewithChilds(2, newTreeNode(3), nil),
		)}, false},
		{"normal", args{newTreeNodewithChilds(1,
			newTreeNodewithChilds(2, newTreeNode(3), nil),
			newTreeNodewithChilds(2, nil, newTreeNode(3)),
		)}, true},

		{"normal-edge", args{newTreeNodewithChilds(1,
			newTreeNodewithChilds(2,
				newTreeNodewithChilds(3, newTreeNode(4), nil),
				nil),
			newTreeNodewithChilds(2,
				nil,
				newTreeNodewithChilds(3, nil, newTreeNode(4))),
		)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
