package demo

import "testing"

func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"edge", args{newTreeNodewithChilds(0, newTreeNode(1), newTreeNode(2))}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
