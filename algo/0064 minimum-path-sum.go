package algo

import "log"

func Test64() {
	log.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i+j == 0 {
				continue
			} else if i == 0 && j != 0 {
				grid[0][j] = grid[0][j] + grid[0][j-1]
			} else if i != 0 && j == 0 {
				grid[i][0] = grid[i][0] + grid[i-1][0]
			} else {
				minV := 0
				if grid[i][j-1] < grid[i-1][j] {
					minV = grid[i][j-1]
				} else {
					minV = grid[i-1][j]
				}
				grid[i][j] = grid[i][j] + minV
			}
		}
	}
	//log.Println(grid)
	return grid[len(grid)-1][len(grid[0])-1]
}
