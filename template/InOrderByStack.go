package template

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrderByStack(root *TreeNode) {
	var (
		stack []*TreeNode
	)
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%d ", p.Val)
			p = p.Right
		}
	}
}
