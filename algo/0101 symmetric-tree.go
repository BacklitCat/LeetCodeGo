package algo

func isSymmetric(root *TreeNode) bool {
	var (
		stackL []*TreeNode
		stackR []*TreeNode
		n      int
	)
	p, q := root, root
	for p != nil || q != nil || n > 0 {
		if p != nil && q != nil {
			stackL = append(stackL, p)
			stackR = append(stackR, q)
			p = p.Left
			q = q.Right
			n++
		} else if p == nil && q == nil && n > 0 {
			p = stackL[n-1]
			stackL = stackL[:n-1]
			q = stackR[n-1]
			stackR = stackR[:n-1]
			//fmt.Println(p.Val, q.Val)
			if p.Val != q.Val {
				return false
			}
			p = p.Right
			q = q.Left
			n--
		} else {
			return false
		}
	}
	return true
}
