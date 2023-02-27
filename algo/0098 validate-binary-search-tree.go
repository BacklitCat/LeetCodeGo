package algo

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var traversal func(node *TreeNode)
	left := math.MinInt
	res := true
	traversal = func(node *TreeNode) {
		if !res || node == nil {
			return
		}
		traversal(node.Left)
		//fmt.Printf("%d ", node.Val)
		if node.Val > left {
			left = node.Val
		} else {
			res = false
			return
		}
		traversal(node.Right)
	}
	traversal(root)
	return res
}
