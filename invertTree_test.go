package demo

import (
	"reflect"
	"testing"
)

func Test_levelOrderScan(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	n1 := newTreeNodeChilds(1, newTreeNode(2), newTreeNode(3))
	n2 := newTreeNodeChilds(1, newTreeNode(2), newTreeNode(3))
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{root: newTreeNode(1)}, []int{1}},
		{"normal", args{root: n1}, []int{1, 2, 3}},
		{"normal", args{root: invertTree(n2)}, []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderScan(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				// if got := levelOrderScan(n1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderScan() = %v, want %v", got, tt.want)
			}
		})
	}
}
