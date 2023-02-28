package algo

func flatten(root *TreeNode) {
	//var reversal func(node)
	p := root
	for p != nil {
		if p.Left != nil && p.Right == nil {
			p = p.Left
			continue
		} else if p.Left == nil && p.Right != nil {
			p.Left = p.Right
			p.Right = nil
		} else if p.Left != nil && p.Right != nil {
			// search the rightmost node of the left subtree
			rightmost := p.Left
			for {
				if rightmost.Right == nil {
					break
				}
				rightmost = rightmost.Right
			}
			rightmost.Right = p.Right
			p.Right = nil
		} else { // nil,nil
			break
		}
	}
	for p = root; p != nil && p.Left != nil; {
		p.Right = p.Left
		p.Left = nil
		p = p.Right
	}
	return
}
