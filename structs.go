package demo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func newTreeNodewithChilds(val int, r, l *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: l, Right: r}
}
