package algo

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	traversal94(root, &res)
	return res
}

func traversal94(node *TreeNode, res *[]int) {
	if node != nil {
		traversal94(node.Left, res)
		*res = append(*res, node.Val)
		traversal94(node.Right, res)
	}
}
