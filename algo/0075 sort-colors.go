package algo

import "log"

func Test75() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	log.Println(nums)
	nums = []int{2, 0, 1}
	sortColors(nums)
	log.Println(nums)
	nums = []int{0, 0}
	sortColors(nums)
	log.Println(nums)
	nums = []int{0, 0, 0}
	sortColors(nums)
	log.Println(nums)
	nums = []int{0, 0, 1}
	sortColors(nums)
	log.Println(nums)
}

func sortColors(nums []int) {
	if len(nums) == 1 {
		return
	}
	for i, j, k, p := -1, -1, -1, 0; p < len(nums); p++ {
		if nums[p] == 0 {
			i++
			j++
			k++
			nums[i] = 0
			if i != j {
				nums[j] = 1
			}
			if j != k {
				nums[k] = 2
			}
		} else if nums[p] == 1 {
			j++
			k++
			nums[j] = 1
			if j != k {
				nums[k] = 2
			}
		} else {
			k++
		}
	}
}
