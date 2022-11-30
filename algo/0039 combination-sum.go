package algo

//func Test() {
//	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
//
//}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backtrack func(choice []int, currSum, idx int)
	backtrack = func(choice []int, currSum, idx int) {
		if currSum == target {
			res = append(res, append([]int(nil), choice...))
			return
		} else if currSum > target || idx == len(candidates) {
			return
		}
		choice = append(choice, candidates[idx])
		currSum += candidates[idx]
		backtrack(choice, currSum, idx)
		choice = choice[:len(choice)-1]
		currSum -= candidates[idx]
		backtrack(choice, currSum, idx+1)
	}
	backtrack([]int{}, 0, 0)
	return res
}
