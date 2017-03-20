package demo

import "testing"

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"normal", args{[]int{0, 1, 2, 3}}, newTreeNodes(2, newTreeNodes(1, newTreeNode(0), nil), newTreeNode(3))},
		{"normal", args{[]int{0, 1, 2}}, newTreeNodes(1, newTreeNode(0), newTreeNode(2))},
		{"normal", args{[]int{0, 1}}, newTreeNodes(1, newTreeNode(0), nil)},
		{"normal", args{[]int{0}}, newTreeNodes(0, nil, nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !isSameTree(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
