package demo

func midTravelTree2(root *TreeNode, sum, target int, result *[]int, mergedResult *[][]int) {
	if root == nil {
		return
	}
	sum += root.Val
	*result = append(*result, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == target {
			*mergedResult = append(*mergedResult, *result)
			return
		}
	}
	leftPath := make([]int, len(*result))
	rightPath := make([]int, len(*result))
	copy(leftPath, *result)
	copy(rightPath, *result)
	midTravelTree2(root.Left, sum, target, &leftPath, mergedResult)
	midTravelTree2(root.Right, sum, target, &rightPath, mergedResult)
}

func hasPathSum2(root *TreeNode, sum int) [][]int {
	var result []int
	var mergedResult [][]int
	midTravelTree2(root, 0, sum, &result, &mergedResult)
	return mergedResult
}
