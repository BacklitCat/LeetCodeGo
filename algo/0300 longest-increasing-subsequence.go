package algo

import "fmt"

func test0300() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))           // 4
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))        // 1
	fmt.Println(lengthOfLIS([]int{5, 4, 3, 2, 1}))              // 1
	fmt.Println(lengthOfLIS([]int{1}))                          // 1
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))          // 3
}

func lengthOfLIS(nums []int) int {
	var (
		d            = []int{nums[0]}
		binarySearch func(arr []int, l, r, target int) int // lower bound
	)

	binarySearch = func(arr []int, l, r, target int) int {
		var mid int
		for l < r {
			mid = l + (r-l)/2
			if arr[mid] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l // l==r
	}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > d[len(d)-1] {
			d = append(d, num)
		} else {
			loc := binarySearch(d, 0, len(d), num)
			// [4,10] find 4, loc=0, do nothing
			// [3,10] find 4, loc=1, let d[1]=4
			d[loc] = num
		}
	}
	return len(d)
}
