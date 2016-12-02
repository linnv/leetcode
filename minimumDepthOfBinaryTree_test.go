package demo

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"edge", args{nil}, 0},
		{"edge", args{&TreeNode{}}, 1},
		{"edge", args{&TreeNode{Val: 9, Left: &TreeNode{}, Right: nil}}, 2},
		{"edge", args{&TreeNode{Val: 9, Right: &TreeNode{}, Left: nil}}, 2},
		{"normal", args{&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Right: &TreeNode{},
				Left:  &TreeNode{},
			},
			Right: &TreeNode{},
		}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
