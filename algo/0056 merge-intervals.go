package algo

import (
	"log"
	"sort"
)

func Test56() {
	log.Println(merge([][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}))
	log.Println(merge([][]int{{1, 4}, {4, 5}}))
	log.Println(merge([][]int{{0, 0}}))
	log.Println(merge([][]int{{0, 0}, {0, 0}}))
	log.Println(merge([][]int{{0, 1}, {0, 1}}))
	log.Println(merge([][]int{{1, 4}, {2, 3}}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	var res [][]int
	for i, l, r := 1, intervals[0][0], intervals[0][1]; i < len(intervals); i++ {
		if r < intervals[i][0] {
			res = append(res, []int{l, r})
			l, r = intervals[i][0], intervals[i][1]
		} else {
			r = Max(intervals[i][1], r)
		}
		if i == len(intervals)-1 {
			res = append(res, []int{l, r})
		}
	}

	return res
}
