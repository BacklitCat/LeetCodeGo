package algo

import (
	"math"
)

/*
bad case
输入[0] 超时，开始时需要特判nums长度
[1,2,1,1,1] 预期3 输出2、4
对左区间计算错误;因为保证能达到最右，所以最右边界肯定更新，应该让左边界等于旧的右边界+1，避免重复计算
对右区间计算错误;忘记添加i位置，不能仅寻找最大step
rightLoc与len错误，应该>= len(nums)-1就可以，忘记-1

执行用时：8 ms, 在所有 Go 提交中击败了96.18%的用户

比官答好一些，官答每次左边界仅+1
*/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	leftLoc := 0
	rightLoc := 0
	step := 0
	maxLoc := math.MinInt
	for leftLoc <= rightLoc {
		for i := leftLoc; i <= rightLoc; i++ {
			maxLoc = Max(maxLoc, i+nums[i])
		}
		step++
		leftLoc = rightLoc + 1
		rightLoc = maxLoc
		if rightLoc >= len(nums)-1 {
			return step
		}
	}
	return -1 // error
}
