package algo

import (
	"fmt"
	"math"
)

func test221() {
	m := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	fmt.Println(maximalSquare(m))
	m = [][]byte{{'1', '0'}, {'0', '1'}}
	fmt.Println(maximalSquare(m))
	m = [][]byte{{'1', '1'}, {'1', '1'}}
	fmt.Println(maximalSquare(m))
	m = [][]byte{{'0'}}
	fmt.Println(maximalSquare(m))
	m = [][]byte{{'0', '0'}, {'0', '0'}}
	fmt.Println(maximalSquare(m))
}

func maximalSquare(matrix [][]byte) int {
	var (
		m   = len(matrix)
		n   = len(matrix[0])
		dp  = make([][]int, m)
		min func(args ...int) int
	)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	min = func(args ...int) (res int) {
		res = math.MaxInt
		for _, arg := range args {
			if arg < res {
				res = arg
			}
		}
		return
	}

	maxLen := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}

			if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			} else {
				dp[i][j] = 1
			}

			if maxLen < dp[i][j] {
				maxLen = dp[i][j]
			}
		}

	}
	//fmt.Println(dp)
	return maxLen * maxLen
}

//
//func maximalSquare(matrix [][]byte) int {
//	var (
//		m, n = len(matrix), len(matrix[0])
//		minL = func() int {
//			if m < n {
//				return m
//			} else {
//				return n
//			}
//		}()
//		accumulate func(i, j, k int) bool
//	)
//	var sumVal = 0
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			matrix[i][j] = matrix[i][j] - 48
//			sumVal += int(matrix[i][j])
//		}
//	}
//	if sumVal == 0 {
//		return 0
//	}
//
//	accumulate = func(i, j, k int) bool {
//		if matrix[i][j]+matrix[i+1][j]+matrix[i][j+1]+matrix[i+1][j+1] == 4 {
//			matrix[i][j] = 1
//			return true
//		} else {
//			matrix[i][j] = 0
//			return false
//		}
//	}
//
//	k := 1
//	for k < minL {
//		status := false
//		//fmt.Println(dp)
//		for i := 0; i < m-k; i++ {
//			for j := 0; j < n-k; j++ {
//				if accumulate(i, j, k) {
//					status = true
//				}
//			}
//		}
//		if !status {
//			break
//		}
//		k++
//	}
//
//	return k * k
//}
