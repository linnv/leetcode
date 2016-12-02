package demo

func levelOrderNoRecursive(root *TreeNode) []int {
	queue := make([]*TreeNode, 0, 3)
	result := make([]int, 0, 3)
	processing := root
	for {
		if processing.Left != nil {
			queue = append(queue, processing.Left)
		}
		if processing.Right != nil {
			queue = append(queue, processing.Right)
		}
		result = append(result, processing.Val)
		if len(queue) == 0 {
			break
		}
		processing = queue[0]
		queue = queue[1:]
	}
	return result
}

func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftCount := 1 + getMaxDepth(root.Left)
	rightCount := 1 + getMaxDepth(root.Right)

	if leftCount > rightCount {
		return leftCount
	}
	return rightCount
}

func levelScanTree(root *TreeNode, level int, vals *[]int) {
	if root == nil {
		return
	}
	if level == 1 { //key point to ger the level data
		*vals = append(*vals, root.Val)
		return
	}
	levelScanTree(root.Left, level-1, vals)
	levelScanTree(root.Right, level-1, vals)
	return
}

func levelOrder(root *TreeNode) [][]int {
	maxDepth := getMaxDepth(root)
	var result [][]int
	for i := 1; i <= maxDepth; i++ {
		var tmp []int
		levelScanTree(root, i, &tmp)
		result = append(result, tmp)
	}
	return result
}

func levelOrderIterate(root *TreeNode) []int {
	maxDepth := getMaxDepth(root)
	var result []int
	for i := 1; i <= maxDepth; i++ {
		levelScanTree(root, i, &result)
	}
	return result
}
