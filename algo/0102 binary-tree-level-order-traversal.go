package algo

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var (
		res   [][]int
		nums  []int
		queue []*TreeNode
		n     = make([]int, 2)
	)

	p, queue, k := root, append(queue, root), 0
	n[k] = 1

	for p != nil && n[k] > 0 {
		p = queue[0]
		queue = queue[1:]
		nums = append(nums, p.Val)
		//fmt.Printf("%d ", p.Val)
		n[k]--
		if p.Left != nil {
			queue = append(queue, p.Left)
			n[(k+1)%2]++
		}
		if p.Right != nil {
			queue = append(queue, p.Right)
			n[(k+1)%2]++
		}
		if n[k] == 0 {
			res = append(res, nums)
			nums = []int{}
			k = (k + 1) % 2
		}
	}
	return res
}
