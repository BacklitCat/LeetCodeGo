package algo

import (
	"fmt"
)

func test322() {
	//fmt.Println(coinChange([]int{1}, 1))                    // 1
	//fmt.Println(coinChange([]int{2}, 4))                    // 2
	//fmt.Println(coinChange([]int{1, 2, 3}, 6))              // 2
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249)) // 20
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	var (
		dp  = make([]int, amount+1)
		min = func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
	)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1 // dp[0]=0
	}

	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin])
			}
		}
		dp[i] = dp[i] + 1
	}

	if dp[amount] >= amount+1 {
		return -1
	}

	return dp[amount]
}
