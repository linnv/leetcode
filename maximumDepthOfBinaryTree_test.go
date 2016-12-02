package demo

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"edge", args{&TreeNode{}}, 1},
		{"edge", args{&TreeNode{Val: 9, Left: &TreeNode{}, Right: nil}}, 2},
		{"edge", args{&TreeNode{Val: 9, Right: &TreeNode{}, Left: nil}}, 2},
		{"normal", args{&TreeNode{Val: 9, Right: &TreeNode{Val: 9, Right: &TreeNode{}, Left: nil}, Left: nil}}, 3},
		{"normal", args{&TreeNode{Val: 9, Right: &TreeNode{Val: 9, Right: &TreeNode{Left: &TreeNode{}}, Left: nil}, Left: &TreeNode{Val: 9, Right: &TreeNode{}, Left: nil}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
