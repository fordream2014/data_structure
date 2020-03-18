package algrithm

//var max int

//二叉树的树高
func getHeight(node *BinaryTreeNode, height int) int{
	if node == nil {
		return height
	}

	height ++
	leftH, rightH := height, height
	if node.Left != nil {
		leftH = getHeight(node.Left, height)
	}
	if node.Right != nil {
		rightH = getHeight(node.Right, height)
	}

	if leftH > rightH {
		return leftH
	} else {
		return rightH
	}
}

//getHeight(root, 0)
