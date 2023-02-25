package algo

import (
	"log"
)

func Test53() {
	log.Println(maxSubArray([]int{1}))
	log.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	log.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	log.Println(maxSubArray([]int{-1}))
	log.Println(maxSubArray([]int{-1, -2}))
}

func maxSubArray(nums []int) int {
	// o(n) 0(1)
	//maxV := -1 << (32<<(^uint(0)>>63) - 1)
	maxV := nums[0]
	for i, temp := 1, 0; i < len(nums); i++ {
		temp += nums[i]
		if temp > maxV {
			maxV = temp
		}
		if temp <= 0 {
			temp = 0
		}
	}
	return maxV
}
