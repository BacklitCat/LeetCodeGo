package algo

func searchMatrix(matrix [][]int, target int) bool {
	var (
		m, n = len(matrix), len(matrix[0])
		i, j = 0, n - 1
	)
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
