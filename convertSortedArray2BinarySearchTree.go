package demo

func build(indexLeft, indexRight int, nums []int) *TreeNode {
	count := indexRight - indexLeft
	if count < 0 {
		return nil
	}
	indexParent := (indexLeft + indexRight + 1) / 2
	parent := &TreeNode{Val: nums[indexParent]}
	parent.Left = build(indexLeft, indexParent-1, nums)
	parent.Right = build(indexParent+1, indexRight, nums)
	return parent
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	return build(0, len(nums)-1, nums)
}
