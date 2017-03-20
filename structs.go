package demo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) ToArray() []int {
	var data []int
	head := ln
	for head != nil {
		data = append(data, head.Val)
		head = head.Next

	}
	return data
}

func (tn *TreeNode) MidOrderData(treeData *[]int) []int {
	if tn == nil {
		return []int{}
	}
	tn.Left.MidOrderData(treeData)
	*treeData = append(*treeData, tn.Val)
	tn.Right.MidOrderData(treeData)
	return *treeData
}

func (tn *TreeNode) PreOrderData(treeData *[]int) []int {
	if tn == nil {
		return []int{}
	}
	*treeData = append(*treeData, tn.Val)
	tn.Left.PreOrderData(treeData)
	tn.Right.PreOrderData(treeData)
	return *treeData
}

func (tn *TreeNode) PosOrderData(treeData *[]int) []int {
	if tn == nil {
		return []int{}
	}
	tn.Left.PosOrderData(treeData)
	tn.Right.PosOrderData(treeData)
	*treeData = append(*treeData, tn.Val)
	return *treeData
}

func (tn *TreeNode) LevelData() []int {
	var data []int
	treeQueue := []*TreeNode{tn}
	var working *TreeNode
	for len(treeQueue) > 0 {
		working = treeQueue[0]
		data = append(data, working.Val)
		if working.Left != nil {
			treeQueue = append(treeQueue, working.Left)
		}
		if working.Right != nil {
			treeQueue = append(treeQueue, working.Right)
		}
		treeQueue = treeQueue[1:]
	}
	return data
}

func newListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func newListNodeWithNext(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func newTreeNodes(val int, left, right *TreeNode) *TreeNode {
	return newTreeNodewithChilds(val, right, left)
}

func newTreeNodewithChilds(val int, r, l *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: l, Right: r}
}
