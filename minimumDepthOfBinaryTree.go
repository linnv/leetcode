package demo

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	}
	leftCount := 1 + minDepth(root.Left)
	rightCount := 1 + minDepth(root.Right)
	if leftCount > rightCount {
		return rightCount
	}
	return leftCount
}
