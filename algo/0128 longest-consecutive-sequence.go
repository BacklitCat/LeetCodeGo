package algo

import "log"

func Test() {
	test128()
}

func test128() {
	log.Println(longestConsecutive([]int{}))                             //0
	log.Println(longestConsecutive([]int{1}))                            //1
	log.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))         //4
	log.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) //9
}

func longestConsecutive(nums []int) int {
	var (
		cache     = make(map[int]int, len(nums))
		maxLength = 0
	)
	for i := 0; i < len(nums); i++ {
		if _, ok := cache[nums[i]]; !ok {
			leftL := cache[nums[i]-1]
			rightL := cache[nums[i]+1]
			newLength := leftL + rightL + 1
			cache[nums[i]] = -1
			cache[nums[i]-leftL] = newLength
			cache[nums[i]+rightL] = newLength
			if newLength > maxLength {
				maxLength = newLength
			}
		}
	}
	return maxLength
}
