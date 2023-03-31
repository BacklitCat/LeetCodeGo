package algo

/*
思路：
二维dp，公式dp[i][j] = dp[i-1][j]+dp[i][j-1] if 非边界+无障碍物 else 0 是障碍物
如果有障碍物，相应部分不加和；障碍物部分dp值是0
初始化 边界1 如果有障碍物 则后续是0

bad case
[[1,0]] 预期0 输出1 初始化j起始打错了，有障碍情况下不能从1位置开始
*/
func uniquePathsWithObstacles2D(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	var dp = make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

/*
一维优化

思路：只和当前列与上一列有关，所以可以只用1维dp

bad case
[[0],[1]] 预期 0 输出1
[[1],[0]] 预期 0 输出1

应该正确初始化dp[0]

	if !(dp[0] == 1 && obstacleGrid[i][0] == 0) {
				dp[0] = 0
	}

执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了84.15%的用户
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	var dp = make([]int, n)

	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[j] = 1
	}

	for i := 1; i < m; i++ {
		if !(dp[0] == 1 && obstacleGrid[i][0] == 0) {
			dp[0] = 0
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}
