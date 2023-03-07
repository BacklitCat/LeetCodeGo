package algo

func invertTree(root *TreeNode) *TreeNode {
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		travel(node.Left)
		travel(node.Right)
	}
	travel(root)
	return root
}
