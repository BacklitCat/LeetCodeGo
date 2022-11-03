package algo

import "math"

func reverse(x int) int {
	res := 0
	d := x % 10
	for x != 0 {
		res = res*10 + d
		x = x / 10
		d = x % 10
	}
	if res < int(math.Pow(2, 31)) && res > int(-math.Pow(2, 31)+1) {
		return res
	}
	return 0
}
