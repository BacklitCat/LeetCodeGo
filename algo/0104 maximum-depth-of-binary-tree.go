package algo

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		maxD      int
		traversal func(node *TreeNode, depth int)
	)
	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		traversal(node.Left, depth+1)
		if depth > maxD {
			maxD = depth
		}
		traversal(node.Right, depth+1)
	}
	traversal(root, 1)
	return maxD
}
