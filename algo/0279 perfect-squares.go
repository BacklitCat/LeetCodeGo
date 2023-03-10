package algo

import (
	"fmt"
	"math"
)

func test279() {
	fmt.Println(numSquares(1))
	fmt.Println(numSquares(2))
	fmt.Println(numSquares(3))
	fmt.Println(numSquares(4))
	fmt.Println(numSquares(5))
	fmt.Println(numSquares(6))
	fmt.Println(numSquares(7))
	fmt.Println(numSquares(8))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(14))
	fmt.Println(numSquares(43))
	fmt.Println(numSquares(266)) // 超时？

	//for i := 1; i < 10000; i++ {
	//	fmt.Printf("%d,", numSquares(i))
	//}
}

// numSquares dp
func numSquares(n int) int {
	if n == 1 {
		return 1
	}
	var (
		dp     = make([]int, n+1)
		jMax   int
		minVal int
		min    = func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
	)
	dp[0], dp[1] = 0, 1
	for i := 2; i < n+1; i++ {
		jMax = int(math.Sqrt(float64(i)))
		minVal = math.MaxInt
		for j := 1; j <= jMax; j++ {
			minVal = min(minVal, dp[i-j*j])
		}
		dp[i] = minVal + 1
	}
	return dp[n]
}

// DFS+回溯 会超时
//func numSquares(n int) int {
//	nSqrt := int(math.Sqrt(float64(n)))
//	if n == nSqrt*nSqrt {
//		return 1
//	}
//	var (
//		res    = math.MaxInt
//		search func(num, target, sum int)
//	)
//	search = func(num, target, sum int) {
//		if num == 0 || target < 0 {
//			return
//		}
//		if target == 0 {
//			if sum < res {
//				res = sum
//			}
//			return
//		}
//		if num <= target {
//			search(num, target-num*num, sum+1)
//		}
//		search(num-1, target, sum)
//	}
//	search(nSqrt, n, 0)
//	return res
//}
