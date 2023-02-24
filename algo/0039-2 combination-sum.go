package algo

import "log"

// 2周目
func Test003902() {
	log.Println(combinationSum003902([]int{2, 3, 6, 7}, 7))
}

func combinationSum003902(candidates []int, target int) [][]int {
	var res [][]int
	candidatesLen := len(candidates)
	var dfsBacktrace func([]int, int, int)
	dfsBacktrace = func(state []int, target, idx int) {
		if target == 0 {
			res = append(res, append([]int{}, state...))
			//log.Println(res)
			return
		}
		if idx == candidatesLen {
			return
		}
		dfsBacktrace(state, target, idx+1)
		nextNum := candidates[idx]
		if nextNum <= target {
			state = append(state, nextNum)
			dfsBacktrace(state, target-nextNum, idx)
			state = state[:len(state)-1]
		}
	}
	var state []int
	dfsBacktrace(state, target, 0)
	return res
}
