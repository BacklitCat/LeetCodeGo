package algo

import "math"

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, len(triangle[i]))
		if i == 0 {
			dp[0][0] = triangle[0][0]
		} else if i > 0 {
			dp[i][0] = dp[i-1][0] + triangle[i][0]
			dp[i][i] = dp[i-1][i-1] + triangle[i][i]
		}

	}
	for i := 2; i < m; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = Min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
	}
	minV := math.MaxInt
	for j := 0; j < n; j++ {
		minV = Min(dp[m-1][j], minV)
	}
	return minV
}
