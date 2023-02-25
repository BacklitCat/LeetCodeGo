package algo

import (
	"log"
)

func Test62() {
	log.Println(uniquePaths(1, 1))   // 1
	log.Println(uniquePaths(1, 2))   // 1
	log.Println(uniquePaths(2, 2))   // 2
	log.Println(uniquePaths(3, 2))   // 3
	log.Println(uniquePaths(3, 3))   // 6
	log.Println(uniquePaths(3, 4))   // 10
	log.Println(uniquePaths(3, 7))   // 28
	log.Println(uniquePaths(7, 3))   // 28
	log.Println(uniquePaths(4, 6))   // 56
	log.Println(uniquePaths(16, 16)) // 155117520
}

func uniquePaths(m int, n int) int {
	if m > n { // make sure m <= n
		m, n = n, m
	}
	if m == 1 { // quick return
		return 1
	} else if m == 2 {
		return n
	}
	// calculate C_{p}^{m}
	m, n = m-1, n-1
	p := m + n
	res := 1
	for i, j := p-m+1, 2; i <= p; i++ {
		res *= i
		if r := res % j; j <= m && r == 0 { // divide in advance
			res /= j
			j++
		}
	}
	return res
}
