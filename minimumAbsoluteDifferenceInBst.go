package demo

func middleTravelTree(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	middleTravelTree(root.Left, data)
	*data = append(*data, root.Val)
	middleTravelTree(root.Right, data)
}

func getMinimumDifference(root *TreeNode) int {
	var data []int
	middleTravelTree(root, &data)
	dataLen := len(data)
	min := data[1] - data[0]
	for i := 2; i < dataLen; i++ {
		if data[i]-data[i-1] < min {
			min = data[i] - data[i-1]
		}
	}
	return min
}
