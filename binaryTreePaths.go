package demo

import "strconv"

func binaryTreePathsIterator(item *string, result *[]string, root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*result = append(*result, *item+"->"+strconv.Itoa(root.Val))
	}
	itemNew := *item + "->" + strconv.Itoa(root.Val)
	if root.Left != nil {
		binaryTreePathsIterator(&itemNew, result, root.Left)
	}
	if root.Right != nil {
		binaryTreePathsIterator(&itemNew, result, root.Right)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	var ret []string
	if root == nil {
		return ret
	}
	if root.Left == nil && root.Right == nil {
		ret = append(ret, strconv.Itoa(root.Val))
		return ret
	}
	item := strconv.Itoa(root.Val)
	binaryTreePathsIterator(&item, &ret, root.Left)
	binaryTreePathsIterator(&item, &ret, root.Right)
	return ret
}
