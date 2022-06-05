package main

func maxArea(height []int) int {
	left, right := 0, len(height) - 1
	ans := 0
	for left < right {
		if height[left] >= height[right] {
			ans = max(ans, (right-left) * height[right])
			right--
		} else {
			ans = max(ans, (right-left) * height[left])
			left++
		}
	}
	return ans
}
