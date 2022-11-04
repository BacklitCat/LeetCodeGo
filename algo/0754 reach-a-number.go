package algo

import "math"

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	n := int(math.Ceil((-1 + math.Sqrt(float64(8*target+1))) / 2))
	if (n*(n+1)/2-target)%2 == 0 {
		return n
	}
	return n + 1 + n%2
}
