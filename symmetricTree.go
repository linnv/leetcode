package demo

func isNodeSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val == right.Val {
		return isNodeSymmetric(left.Left, right.Right) && isNodeSymmetric(left.Right, right.Left)
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isNodeSymmetric(root.Left, root.Right)
}

//it works, but code is too much long
// func getMaxDepthSymmetricTree(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
//
// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}
// 	leftCount := 1 + getMaxDepthSymmetricTree(root.Left)
// 	rightCount := 1 + getMaxDepthSymmetricTree(root.Right)
//
// 	if leftCount > rightCount {
// 		return leftCount
// 	}
// 	return rightCount
// }
//
// func levelScanTreeSymmetricTree(root *TreeNode, level int, vals *[]int) {
// 	if level == 1 { //key point to get the level data
// 		if root == nil {
// 			*vals = append(*vals, -1)
// 			return
// 		}
// 		*vals = append(*vals, root.Val)
// 		return
// 	}
// 	if root == nil {
// 		return
// 	}
// 	levelScanTreeSymmetricTree(root.Left, level-1, vals)
// 	levelScanTreeSymmetricTree(root.Right, level-1, vals)
// 	return
// }
//
// func isSymmetric(root *TreeNode) bool {
// 	isSymmetricSlice := func(num []int) bool {
// 		numLen := len(num)
// 		if numLen%2 == 1 {
// 			return false
// 		}
// 		half := numLen / 2
// 		for i := 0; i < half; i++ {
// 			if num[i] != num[numLen-i-1] {
// 				return false
// 			}
// 		}
// 		return true
// 	}
//
// 	maxDepth := getMaxDepthSymmetricTree(root)
// 	var result []int
// 	for i := 2; i <= maxDepth; i++ {
// 		levelScanTreeSymmetricTree(root, i, &result)
// 		if !isSymmetricSlice(result) {
// 			return false
// 		}
// 		result = result[:0]
// 	}
// 	return true
// }
