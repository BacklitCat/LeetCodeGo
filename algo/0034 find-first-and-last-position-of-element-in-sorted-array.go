package algo

//func Test() {
//	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
//	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
//	fmt.Println(searchRange([]int{}, 0))
//	fmt.Println(searchRange([]int{2, 2}, 1))
//}

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{-1, -1}
	} else if numsLen == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	l, r, m := 0, numsLen-1, -1

	for l < r {
		m = (l + r) / 2
		if nums[m] == target {
			break
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if l == r {
		if nums[l] == target {
			return []int{l, l}
		} else {
			return []int{-1, -1}
		}
	}

	for i := 0; ; i++ {
		if m-i >= 0 && nums[m-i] == target {
			l = m - i
		}
		if m+i < numsLen && nums[m+i] == target {
			r = m + i
		}
		if l != m-i && r != m+i {
			break
		}
	}
	if l <= r {
		return []int{l, r}
	}
	return []int{-1, -1}
}
