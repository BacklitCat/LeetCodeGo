package template

import "fmt"

func testQuickSort() {
	nums := []int{1, 4, 7, 2, 3, 3, 9, 0}
	QuickSort(nums)
	fmt.Println(nums)
}

func partition(nums []int, l, r int) int {
	var pivot, i, j = l, l, r
	for i < j {
		for i < j && nums[j] >= nums[pivot] {
			j--
		}
		for i < j && nums[i] <= nums[pivot] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[pivot] = nums[pivot], nums[i] // i=j
	return i
}

func quickSort(nums []int, l, r int) {
	if l > r {
		return
	}
	i := partition(nums, l, r)
	quickSort(nums, l, i-1)
	quickSort(nums, i+1, r)
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}
