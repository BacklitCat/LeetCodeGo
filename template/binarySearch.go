package template

import "fmt"

func testBinarySearch() {
	fmt.Println(BinarySearchLowerBound([]int{1, 2, 3, 4, 5, 5, 5, 6, 7}, 5, 0, 9)) //4
	fmt.Println(BinarySearchUpperBound([]int{1, 2, 3, 4, 5, 5, 5, 6, 7}, 5, 0, 9)) //7
	fmt.Println(BinarySearchLowerBound([]int{1, 2, 3, 4, 5, 5, 5, 6, 7}, 7, 0, 9)) //8
	fmt.Println(BinarySearchUpperBound([]int{1, 2, 3, 4, 5, 5, 5, 6, 7}, 7, 0, 9)) //9
	fmt.Println(BinarySearchLowerBound([]int{0, 4, 12}, 4, 0, 3))                  //1
	fmt.Println(BinarySearchUpperBound([]int{0, 4, 12}, 4, 0, 3))                  //2
	fmt.Println(BinarySearchLowerBound([]int{0, 4, 12}, 2, 0, 3))                  //1
	fmt.Println(BinarySearchUpperBound([]int{0, 4, 12}, 2, 0, 3))                  //1
	fmt.Println(BinarySearchUpperBound([]int{0, 2, 4, 12}, 2, 0, 3))               //2
	fmt.Println(BinarySearchLowerBound([]int{2, 5}, 3, 0, 2))                      //1
	fmt.Println(BinarySearchUpperBound([]int{2, 5}, 3, 0, 2))                      //1
	fmt.Println(BinarySearchLowerBound([]int{2, 5}, 1, 0, 2))                      //0
	fmt.Println(BinarySearchUpperBound([]int{2, 5}, 1, 0, 2))                      //0
}

// BinarySearchLowerBound 求非降序范围[l, r)内第一个不小于value的值的位置
func BinarySearchLowerBound(nums []int, target, l, r int) (idx int) {
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l //l==r
}

// BinarySearchUpperBound 求非降序范围[l, r)内第一个大于value的值的位置
func BinarySearchUpperBound(nums []int, target, l, r int) (idx int) {
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l //l==r
}
