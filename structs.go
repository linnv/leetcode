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

func newListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func newListNodeWithNext(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func newTreeNodewithChilds(val int, r, l *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: l, Right: r}
}
