package demo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftCount := 1 + maxDepth(root.Left)
	rightCount := 1 + maxDepth(root.Right)

	if leftCount > rightCount {
		return leftCount
	}
	return rightCount
}
