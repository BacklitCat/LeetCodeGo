package algo

func test448() {
	findDisappearedNumbers([]int{2, 2})

}

func findDisappearedNumbers(nums []int) []int {
	var res []int
	numsLen := len(nums)
	for _, v := range nums {
		if v > numsLen {
			v = v % numsLen
		}
		if v == 0 {
			v = numsLen
		}
		nums[v-1] += numsLen
	}
	for i, v := range nums {
		if nums[i] <= numsLen {
			res = append(res, i+1)
			continue
		}
		nums[i] = v % numsLen
	}
	return res
}
