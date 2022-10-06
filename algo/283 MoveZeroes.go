package algo

/*
283. 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]
*/

// MoveZeroesBubble 冒泡思想
/*每一次遍历移动一个零到末尾，直到没有移动元素停止
执行用时：156 ms, 在所有 Go 提交中击败了5.56%的用户
内存消耗：6.4 MB, 在所有 Go 提交中击败了98.85%的用户*/
func MoveZeroesBubble(nums []int) {
	length := len(nums)
	var flag bool
	for i1 := 0; i1 < length-1; i1++ {
		flag = false
		for i2 := i1; i2 < length-1; i2++ {
			if nums[i2] == 0 && nums[i2+1] != 0 {
				nums[i2], nums[i2+1] = nums[i2+1], nums[i2]
				flag = true
			}
		}
		// [0,0,1]：0与0交换情况，外层循环游标需回原位置
		if i1 > 0 && nums[i1-1] == 0 {
			i1 -= 2
			continue
		}
		// 如果没有交换，提前退出
		if !flag {
			return
		}
	}
	return
}

// MoveZeroesInsert 双指针插入赋值思想
/*维护两个指针，左指针维护非0元素，右指针探索
右指针遇到0元素直接右移，遇到非0则左指针赋值
等到右指针到边界时，左右指针之间赋0
执行用时：12 ms, 在所有 Go 提交中击败了99.74%的用户
内存消耗：6.5 MB, 在所有 Go 提交中击败了6.16%的用户*/
func MoveZeroesInsert(nums []int) {
	length := len(nums)
	// p1指向的位置是下一个等待插入的位置
	// p2指向的位置是等待探索的位置
	// 由于go语言不推荐移动指针，需要调用unsafe库，所以这里指针用游标替代
	p1, p2 := 0, 0
	for p2 != length {
		if nums[p2] == 0 {
			p2 += 1
		} else {
			nums[p1] = nums[p2]
			p1 += 1
			p2 += 1
		}
	}
	for p1 != length {
		nums[p1] = 0
		p1 += 1
	}
	return
}

// MoveZerosPivot pivot思想
/*参照快排思路，将0视为pivot，将非0元素放到pivot左边，将0放在pivot右边*/
func MoveZerosPivot(nums []int) {
	length := len(nums)
	// pivot探测0，p是待交换位置
	p, pivot := 0, 0
	for pivot != length {
		if nums[pivot] != 0 {
			nums[p], nums[pivot] = nums[pivot], nums[p]
			p += 1
			pivot += 1
		} else {
			pivot += 1
		}
	}
	return
}
