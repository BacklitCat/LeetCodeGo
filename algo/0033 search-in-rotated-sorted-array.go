package algo

import "fmt"

//func Test() {
//	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
//	fmt.Println(search([]int{1, 3}, 1))    // debug
//	fmt.Println(search([]int{1, 3}, 3))    // debug
//	fmt.Println(search([]int{1, 3, 5}, 5)) // debug
//	fmt.Println(search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8)) // debug
//}

func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left, right, mid := 0, numsLen-1, -1
	for left != right-1 {
		mid = (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid
		} else {
			right = mid
		}
	}
	mid = left

	if nums[mid] < nums[numsLen-1] {
		mid = numsLen - 1
	}
	fmt.Println("mid:", mid)

	if target < nums[0] {
		left = mid + 1
		if left == numsLen { // 处理越界
			return -1
		}
		right = numsLen - 1
	} else {
		left = 0
		right = mid
	}
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
