package template

import "fmt"

func testMaxHeap() {
	nums := []int{2, 5, 3, 7, 9, 0, 1}
	//buildMaxHeap(nums, len(nums))
	maxHeapSort(nums)
	fmt.Println(nums)
}

func maxHeapSort(arr []int) {
	heapSize := len(arr)
	buildMaxHeap(arr, heapSize)
	for i := 0; i < heapSize; i++ {
		arr[0], arr[heapSize-i-1] = arr[heapSize-i-1], arr[0]
		buildMaxHeap(arr, heapSize-i-1)
	}
}

func buildMaxHeap(arr []int, heapSize int) {
	for i := len(arr) / 2; i >= 0; i-- {
		maxHeapify(arr, i, heapSize)
	}
}

func maxHeapify(arr []int, root, heapSize int) {
	l, r, loc := root*2+1, root*2+2, root
	if l < heapSize && arr[l] > arr[loc] {
		loc = l
	}
	if r < heapSize && arr[r] > arr[loc] {
		loc = r
	}
	if root != loc {
		arr[loc], arr[root] = arr[root], arr[loc]
		maxHeapify(arr, root, heapSize)
	}
}
