package demo

func levelOrderBottom(root *TreeNode) [][]int {
	maxDepth := getMaxDepth(root)
	var result [][]int
	for i := maxDepth; i >= 1; i-- {
		var tmp []int
		levelScanTree(root, i, &tmp)
		result = append(result, tmp)
	}
	return result
}
