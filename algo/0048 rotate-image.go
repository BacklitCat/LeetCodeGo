package algo

import "log"

func Test() {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(m)
	log.Println(m)
	m = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(m)
	log.Println(m)
}

func rotate(matrix [][]int) {
	if len(matrix) == 1 {
		return
	}
	n := len(matrix)
	var (
		diagonal   func([][]int)
		horizontal func([][]int)
	)
	// twice
	diagonal = func(matrix [][]int) {
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
	horizontal = func(matrix [][]int) {
		for i := 0; i < n; i++ {
			for j := 0; j < n/2; j++ {
				matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
			}
		}
	}
	diagonal(matrix)
	horizontal(matrix)
}
