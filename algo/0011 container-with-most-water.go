package algo

func maxArea(height []int) int {
	l, r, maxS := 0, len(height)-1, 0

	for l < r {
		s := Min(height[l], height[r]) * (r - l)
		maxS = Max(s, maxS)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxS
}
