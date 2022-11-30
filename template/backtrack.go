package template

/*
例子1
从1到4，数字两两组合的所有可能性。
*/

func BackTrackEg() [][]int {
	numList := []int{1, 2, 3, 4}
	var res [][]int
	var backtrack func(state []int, idx int)
	backtrack = func(state []int, idx int) {
		if len(state) == 2 {
			res = append(res, append([]int(nil), state...))
			return
		}
		for i := idx; i < len(numList); i++ {
			state = append(state, numList[i])
			backtrack(state, i+1)
			state = state[:len(state)-1]
		}
		return
	}
	backtrack([]int{}, 0)
	return res
}
