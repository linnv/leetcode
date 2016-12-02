package demo

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"edge", args{&TreeNode{Val: 1}, int(1)}, true},
		{"edge", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, int(3)}, true},
		{"edge", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, int(2)}, false},
		{"normal", args{&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Right: &TreeNode{},
				Left:  &TreeNode{},
			},
			Right: &TreeNode{},
		}, int(1)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
