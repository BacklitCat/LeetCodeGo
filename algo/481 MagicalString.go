package algo

import (
	"strconv"
	"strings"
)

/*
481. 神奇字符串
神奇字符串 s 仅由 '1' 和 '2' 组成，并需要遵守下面的规则：

神奇字符串 s 的神奇之处在于，串联字符串中 '1' 和 '2' 的连续出现次数可以生成该字符串。
s 的前几个元素是 s = "1221121221221121122……" 。如果将 s 中连续的若干 1 和 2 进行分组，可以得到 "1 22 11 2 1 22 1 22 11 2 11 22 ......"。
每组中 1 或者 2 的出现次数分别是 "1 2 2 1 1 2 1 2 2 1 2 2 ......" 。上面的出现次数正是 s 自身。

给你一个整数 n ，返回在神奇字符串 s 的前 n 个数字中 1 的数目。

示例 1：
输入：n = 6
输出：3
解释：神奇字符串 s 的前 6 个元素是 “122112”，它包含三个 1，因此返回 3 。

示例 2：
输入：n = 1
输出：1

提示：
1 <= n <= 105
*/

/*
GetMagicStrOneTimes
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了46.15%的用户
通过测试用例：64 / 64
*/
func GetMagicStrOneTimes(n int) int {
	if n <= 3 {
		return 1
	}

	var nums []int
	nums = make([]int, n)
	nums[0], nums[1], nums[2] = 1, 2, 2

	for i, j := 3, 2; i < n; j++ {
		insertValue := nums[i-1] ^ 3
		insertTimes := nums[j]

		if insertTimes == 1 {
			nums[i] = insertValue
			i = i + 1
		} else {
			nums[i] = insertValue
			if i+1 < n {
				nums[i+1] = insertValue
			}
			i = i + 2
		} //insert

	} //for

	SumOneTimes := 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			SumOneTimes++
		}
	}
	return SumOneTimes

}

/*
GenMagicStr
生成字符串
*/
func GenMagicStr(n int) string {
	if n <= 3 {
		return "122"[0 : n-1]
	}

	var builder strings.Builder
	builder.Grow(n)

	var nums []int
	nums = make([]int, n)
	//nums = append(nums, 1, 2, 2)
	nums[0], nums[1], nums[2] = 1, 2, 2

	for i, j := 3, 2; i < n; j++ {
		insertValue := nums[i-1] ^ 3
		insertTimes := nums[j]

		if insertTimes == 1 {
			nums[i] = insertValue
			i = i + 1
		} else {
			nums[i] = insertValue
			if i+1 < n {
				nums[i+1] = insertValue
			}
			i = i + 2
		} //insert

	} //for

	for i := 0; i < n; i++ {
		builder.WriteString(strconv.Itoa(nums[i]))
	}
	return builder.String()

}
