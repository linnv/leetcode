package demo

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"edge", args{&TreeNode{}}, [][]int{[]int{0}}},
		{"edge", args{&TreeNode{Val: 1, Left: &TreeNode{}}}, [][]int{[]int{0}, []int{1}}},
		{"normal", args{&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Right: &TreeNode{},
				Left:  &TreeNode{},
			},
			Right: &TreeNode{},
		}}, [][]int{[]int{0, 0}, []int{2, 0}, []int{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
