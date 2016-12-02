package demo

import (
	"reflect"
	"testing"
)

func Test_hasPathSum2(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"edge", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, int(3)}, [][]int{[]int{1, 2}}},
		{"edge", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}}, int(3)}, [][]int{[]int{1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum2(tt.args.root, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hasPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
