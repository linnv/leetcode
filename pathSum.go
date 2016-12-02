package demo

func midTravelTree(root *TreeNode, sum, target int, found *bool) {
	if root == nil {
		return
	}
	sum += root.Val
	if sum == target && root.Left == nil && root.Right == nil {
		*found = true
		return
	}
	midTravelTree(root.Left, sum, target, found)
	midTravelTree(root.Right, sum, target, found)
}

func hasPathSum(root *TreeNode, sum int) bool {
	var findIt bool
	midTravelTree(root, 0, sum, &findIt)
	return findIt
}
