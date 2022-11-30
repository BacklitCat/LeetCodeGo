package algo

//func Test() {
//	fmt.Println(generateParenthesis(3))
//
//}

func generateParenthesis(n int) []string {
	var res []string
	var backtrack func(choice []byte, lCount, rCount int)
	backtrack = func(choice []byte, lCount, rCount int) {
		if len(choice) == n*2 && lCount == rCount {
			res = append(res, string(choice))
			return
		} else if lCount > n || rCount > lCount { // 剪枝
			return
		}
		choice = append(choice, '(')
		backtrack(choice, lCount+1, rCount)
		choice = choice[:len(choice)-1]
		choice = append(choice, ')')
		backtrack(choice, lCount, rCount+1)
	}
	backtrack([]byte{}, 0, 0)
	return res
}
