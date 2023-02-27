package algo

import (
	"fmt"
)

func Test96() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, numTrees(i))
	}
}

func numTrees(n int) int {
	g := make([]int, n+1)
	g[0], g[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}

/*
// 记忆递归 会超时 要优化为for循环
func numTrees(n int) int {
	var (
		G     func(int) int
		cache = make([]int, n+1)
	)
	cache[0], cache[1] = 1, 1
	G = func(n int) int {
		if cache[n] != 0 {
			return cache[n]
		}
		sum := 0
		for i := 1; i <= n; i++ {
			sum += G(i-1) * G(n-i)
		}
		return sum
	}
	return G(n)
}
*/
