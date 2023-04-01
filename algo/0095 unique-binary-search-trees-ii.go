package algo

func generateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{{1, nil, nil}}
	}
	var gen func(start, end int) []*TreeNode
	gen = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		var res []*TreeNode
		for i := start; i <= end; i++ {
			leftTrees := gen(start, i-1)
			rightTrees := gen(i+1, end)

			for _, left := range leftTrees {
				for _, right := range rightTrees {
					root := &TreeNode{
						Val:   i,
						Left:  left,
						Right: right,
					}
					res = append(res, root)
				}
			}
		}
		return res
	}

	return gen(1, n)
}
