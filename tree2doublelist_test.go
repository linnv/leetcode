package demo

import (
	"reflect"
	"testing"
)

func TestTree2DoubleLinkedList(t *testing.T) {
	leftOnly, leftTmp := newTreeNodes(1, newTreeNodes(2, newTreeNodes(3, nil, nil), nil), nil), &TreeNode{}
	rightOnly, rightTmp := newTreeNodes(1, nil, newTreeNodes(2, nil, newTreeNodes(3, nil, nil))), &TreeNode{}
	leftRight, leftRightTmp := newTreeNodes(1, newTreeNode(2), newTreeNode(3)), &TreeNode{}
	leftRight1, leftRightTmp1 := newTreeNodes(1, newTreeNode(2), newTreeNodes(4, newTreeNode(5), newTreeNode(6))), &TreeNode{}
	type args struct {
		root *TreeNode
		next **TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"left only", args{leftOnly, &leftTmp}},
		{"right only ", args{rightOnly, &rightTmp}},
		{"left right", args{leftRight, &leftRightTmp}},
		{"left right 1", args{leftRight1, &leftRightTmp1}},
	}

	intSlice := make([]int, 0, 2)
	retIntSlice := make([]int, 0, 2)
	for _, tt := range tests {
		intSlice = intSlice[:]
		retIntSlice = retIntSlice[:]

		tt.args.root.MidOrderData(&intSlice)

		Tree2DoubleLinkedList(tt.args.root, tt.args.next)

		// ToLeft(&tt.args.root) or
		tt.args.root = ToLeftYah(tt.args.root)
		for tt.args.root != nil {
			retIntSlice = append(retIntSlice, tt.args.root.Val)
			tt.args.root = tt.args.root.Right
		}

		if !reflect.DeepEqual(intSlice, retIntSlice) {
			t.Errorf("Tree2DoubleLinkedList[%s] Get retIntSlice: %+v,cap(retIntSlice):%d,len(retIntSlice):%d arrd:%v  want \n", tt.name, retIntSlice, cap(retIntSlice), len(retIntSlice), &retIntSlice[0])
			t.Errorf("intSlice: %+v,cap(intSlice):%d,len(intSlice):%d arrd:%v \n", intSlice, cap(intSlice), len(intSlice), &intSlice[0])
		}
	}
}
