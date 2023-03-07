package algo

import "fmt"

func test02152() {
	fmt.Println(findKthLargest2([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest2(nums []int, k int) int {
	// 快排方法

	var (
		res  int
		stop bool
		loc  = len(nums) - k
	)

	check := func(i, pivot int) {
		//fmt.Println("check", i)
		if i == loc {
			res = pivot
			stop = true
		}
	}

	var partition func(left, right int)
	partition = func(left, right int) {
		if left == right {
			check(left, nums[left])
			return
		}

		i, j, pivot := right, right, nums[left]
		for i != left && !stop {
			if nums[i] > pivot {
				nums[i], nums[j] = nums[j], nums[i]
				j--
			}
			i--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
		// i==j
		nums[i] = pivot
		check(i, pivot)
		// 快速选择，只选择一半的区间
		if i > loc && left <= i-1 {
			partition(left, i-1)
		}
		if i < loc && i+1 <= right {
			partition(i+1, right)
		}
	}

	partition(0, len(nums)-1)
	//fmt.Println(nums)

	return res
}
