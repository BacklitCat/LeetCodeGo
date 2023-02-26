package algo

import "log"

func Test78() {
	log.Println(subsets([]int{1}))
	log.Println(subsets([]int{1, 2, 3}))
	log.Println(subsets([]int{9, 0, 3, 5, 7}))
	log.Println(subsets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func subsets(nums []int) [][]int {
	if len(nums) == 1 {
		return append([][]int{nil}, nums)
	}
	var (
		res       [][]int
		recursive func([]int, int)
	)
	recursive = func(state []int, i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, state...))
			return
		}
		recursive(state, i+1)
		recursive(append(state, nums[i]), i+1)
	}
	recursive([]int(nil), 0)
	return res
}
