package algo

import (
	"fmt"
	"sort"
)

/*
test :[-1,0,1,2,-1,-4] [3,0,-2,-1,1,2]
*/

func threeSum(nums []int) [][]int { // bug [-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]
	var ans [][]int
	sort.Ints(nums)
	fmt.Println(nums)
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		target, p, q := -nums[i], i+1, numsLen-1
		for p < q {
			twoNumSum := nums[p] + nums[q]
			if twoNumSum < target {
				p++
			} else if twoNumSum > target {
				q--
			} else {
				if len(ans) == 0 || ans[len(ans)-1][0] != nums[i] || ans[len(ans)-1][1] != nums[p] || ans[len(ans)-1][2] != nums[q] {
					ans = append(ans, []int{nums[i], nums[p], nums[q]})
					fmt.Println("append", ans, i, p, q)
				}
				if nums[p] == nums[p-1] || nums[p+1]+nums[q] < target {
					p++
				} else if (q < numsLen-1 && nums[q] == nums[q-1]) || nums[p]+nums[q-1] > target {
					q--
				} else {
					fmt.Println("else", i, p, q)
					q--
				}
			}
		}
	}
	return ans
}

// answer

func threeSumAnswer(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
