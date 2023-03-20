package algo

func pathSum(root *TreeNode, targetSum int) int {

	var (
		res    int
		preSum = map[int]int{0: 1} // curr[count]
		dfs    func(node *TreeNode, curr int)
	)

	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}
		curr += node.Val // 累计前缀和
		// ∵ 当前结点累计前缀和-起始结点累计前缀和=目标 ∴当前结点累计前缀和-目标=起始结点累计前缀和
		count := preSum[curr-targetSum] //前缀 curr-targetSum 代表起始结点的前缀和
		res += count                    // 累加方案数量
		preSum[curr]++                  // 当前结点累计前缀和为 curr 的情况+1
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]-- // 回溯
	}
	dfs(root, 0)
	return res
}

/*
func pathSum(root *TreeNode, targetSum int) int {
	var (
		res  int
		dfsA func(node *TreeNode, target int)
		dfsB func(node *TreeNode)
	)
	dfsA = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		// if target == 0 { // 不能这么写，因为到nil才为0会丢失解
		// 	res++
		// }
		if node.Val == target {
			res++
		}
		//if target >= node.Val { // 不能限制，因为可能为负值
		dfsA(node.Left, target-node.Val)
		dfsA(node.Right, target-node.Val)
		//}
	}
	dfsB = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfsA(node, targetSum)
		dfsB(node.Left)
		dfsB(node.Right)

	}
	dfsB(root)
	return res
}
*/
/*
func pathSum(root *TreeNode, targetSum int) int {
	var (
		res int
		dfs func(node *TreeNode, target int)
	)
	dfs = func(node *TreeNode, target int) {
		if node == nil {
			return
		}

		if node.Val == target {
			res++
		}

		dfs(node.Left, targetSum)
		dfs(node.Right, targetSum)
		dfs(node.Left, target-node.Val)
		dfs(node.Right, target-node.Val)
	}
	dfs(root, targetSum)
	return res
}

*/
