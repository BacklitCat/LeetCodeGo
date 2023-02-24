package algo

import (
	"fmt"
	"log"
	"sort"
)

func Test40() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{1, 2, 2, 2, 5}, 5))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	fmt.Println(combinationSum2([]int{14, 6, 25, 9, 30, 20, 33, 34, 28, 30, 16, 12, 31, 9, 9, 12, 34,
		16, 25, 32, 8, 7, 30, 12, 33, 20, 21, 29, 24, 17, 27, 34, 11, 17, 30, 6, 32, 21, 27, 17, 16,
		8, 24, 12, 12, 28, 11, 33, 10, 32, 22, 13, 34, 18, 12}, 27))

}

func combinationSum2(candidates []int, target int) [][]int {
	var (
		res          [][]int
		state        []int
		dfsBacktrace func(int, int)
	)
	dfsBacktrace = func(target int, idx int) {
		if target == 0 {
			res = append(res, append([]int{}, state...))
			return
		}
		if idx == len(candidates) {
			return
		}
		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			num := candidates[i]
			dst := target - num
			if dst >= 0 {
				log.Println(state)
				state = append(state, num)
				dfsBacktrace(dst, i+1)
				state = state[:len(state)-1]
			}
		}
	}
	sort.Ints(candidates)
	dfsBacktrace(target, 0)
	return res
}
