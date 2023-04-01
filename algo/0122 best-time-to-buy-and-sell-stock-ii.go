package algo

import "math"

/*
dp[i][0] 第i天的最大利润(不持有股票)
dp[i][1] 第i天的最大利润(持有股票)

dp[i][0] = dp[i-1][0] 继续不买
         = dp[i-1][1] + prices[i-1] 昨天还持有，今天卖出 (prices下标从0)
dp[i][1] = dp[i-1][1] 继续持有
         = dp[i-1][0] - prices[i-1] 今天购买
         = dp[i][0] - prices[i] 今卖今又买

长度 n+1
初始化 dp[0][0],dp[0][1] = - price[0]
答案 max{dp[n][0],dp[n][1]} = dp[n][0]
*/

func maxProfit(prices []int) int {
	n := len(prices)
	var dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i <= n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = Maxs(dp[i-1][1], dp[i-1][0]-prices[i-1], dp[i][0]-prices[i-1])
	}
	return dp[n][0]
}

func Maxs(args ...int) int {
	maxV := math.MinInt
	for _, num := range args {
		maxV = Max(maxV, num)
	}
	return maxV
}
