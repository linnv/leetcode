package demo

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func newTreeNodeChilds(val int, l, r *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: l, Right: r}
}

func levelOrderScan(root *TreeNode) []int {
	ns := make([]*TreeNode, 0, 3)
	var vals []int
	ns = append(ns, root)
	var process *TreeNode
	for len(ns) > 0 {
		process = ns[0]
		if process != nil {
			vals = append(vals, process.Val)
			ns = append(ns, process.Left)
			ns = append(ns, process.Right)
		}
		ns = ns[1:]
	}
	return vals
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftTmp := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = leftTmp
	return root
}
