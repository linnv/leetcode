package demo

import (
	"reflect"
	"testing"
)

func Test_levelOrderNoRecursive(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Right: &TreeNode{},
				Left:  &TreeNode{},
			},
			Right: &TreeNode{},
		}}, []int{1, 2, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderNoRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderNoRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrderIterate(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Right: &TreeNode{},
				Left:  &TreeNode{},
			},
			Right: &TreeNode{},
		}}, []int{1, 2, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderIterate(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderIterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_levelOrderIterate(b *testing.B) {

	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Right: &TreeNode{},
			Left:  &TreeNode{},
		},
		Right: &TreeNode{},
	}
	for i := 0; i < b.N; i++ {
		levelOrderIterate(root)
	}
}

func Benchmark_levelOrderNoRecursive(b *testing.B) {

	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Right: &TreeNode{},
			Left:  &TreeNode{},
		},
		Right: &TreeNode{},
	}
	for i := 0; i < b.N; i++ {
		levelOrderNoRecursive(root)
	}
}

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {"edge", args{nil}, [][]int{}}, DeepEqual just is bad to tackle this
		{"edge", args{&TreeNode{}}, [][]int{[]int{0}}},
		{"edge", args{&TreeNode{Left: &TreeNode{}}}, [][]int{[]int{0}, []int{0}}},
		{"normal", args{&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Right: &TreeNode{},
				Left:  &TreeNode{},
			},
			Right: &TreeNode{},
		}}, [][]int{[]int{1}, []int{2, 0}, []int{0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
