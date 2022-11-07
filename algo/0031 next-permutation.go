package algo

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	var l, r int
	for l = len(nums) - 1; nums[l-1] >= nums[l]; l-- {
		if l-1 == 0 {
			revese31(nums, 0, len(nums)-1)
			return
		}
	}
	for r = len(nums) - 1; nums[r] <= nums[l-1]; r-- {
	}
	nums[l-1], nums[r] = nums[r], nums[l-1]
	revese31(nums, l, len(nums)-1)
}

func revese31(nums []int, l, r int) {
	for i, j := l, r; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
