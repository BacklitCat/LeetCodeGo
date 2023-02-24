package algo

import "log"

func Test46() {
	log.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return append([][]int{nums})
	}
	var res [][]int
	numsLen := len(nums)
	swap := func(arr []int, i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}
	var backtrack func([]int, int)
	backtrack = func(nums []int, chosen int) {
		if chosen == numsLen {
			//log.Println(nums)
			res = append(res, append([]int{}, nums...))
			return
		}
		for i := chosen; i < numsLen; i++ {
			swap(nums, chosen, i)
			backtrack(nums, chosen+1)
			swap(nums, chosen, i)
		}
	}

	backtrack(nums, 0)
	return res
}
