package demo

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"normal", args{newTreeNodewithChilds(1, newTreeNodewithChilds(2, nil, newTreeNode(5)), newTreeNode(3))}, []string{"1->3", "1->2->5"}},
		{"normal", args{newTreeNodewithChilds(1, newTreeNodewithChilds(2, nil, newTreeNode(5)), newTreeNodewithChilds(2, nil, newTreeNode(5)))}, []string{"1->2->5", "1->2->5"}},
		{"edge", args{newTreeNode(1)}, []string{"1"}},
		{"edge", args{newTreeNodewithChilds(1, nil, newTreeNode(2))}, []string{"1->2"}},
		{"edge", args{newTreeNodewithChilds(1, newTreeNode(2), nil)}, []string{"1->2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
