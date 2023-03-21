package algo

func buildTree(preorder []int, inorder []int) *TreeNode {

	var (
		m     = make(map[int]int)
		n     = len(preorder)
		build func(l1, r1, l2, r2 int) *TreeNode
	)

	for i, num := range inorder {
		m[num] = i
	}

	build = func(l1, r1, l2, r2 int) *TreeNode {
		if l1 > r1 {
			return nil
		}

		rootIdx := l1
		rootIdxIn := m[preorder[rootIdx]]
		leftNum := rootIdxIn - l2

		node := &TreeNode{
			Val:   preorder[rootIdx],
			Left:  build(l1+1, l1+leftNum, l2, rootIdxIn-1),
			Right: build(l1+1+leftNum, r1, rootIdxIn+1, r2),
		}

		return node
	}
	return build(0, n-1, 0, n-1)
}
