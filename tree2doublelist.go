package demo

func Tree2DoubleLinkedList(root *TreeNode, leftNode **TreeNode) {
	if root == nil {
		return
	}
	var convertNode *TreeNode
	convertNode = root
	if convertNode.Left != nil {
		Tree2DoubleLinkedList(convertNode.Left, leftNode)
	}
	convertNode.Left = *leftNode
	if convertNode != nil {
		convertNode.Left.Right = convertNode
	}
	(*leftNode) = convertNode
	if convertNode.Right != nil {
		Tree2DoubleLinkedList(convertNode.Right, leftNode)
	}
}

func ToLeftYah(root *TreeNode) *TreeNode {
	if (root) == nil {
		return root
	}
	for root.Left != nil {
		root = root.Left
	}
	//why?
	//var left **TreeNode
	//because *left doesn't belong to tree, but to use *left, we must make sure left can't be nil and *left==&TreeNode{}, which doesn't belong to the Tree
	if root.Right != nil {
		root = root.Right
	}
	return root
}

func ToLeft(root **TreeNode) {
	if root == nil || (*root) == nil {
		return
	}
	for (*root).Left != nil {
		*root = (*root).Left
	}
	if (*root).Right != nil {
		*root = (*root).Right
	}
}
