package main

import (
	"math"
	"strconv"
)
func reverse(x int) int {
	a := strconv.Itoa(x)
	left, right := 0, len(a) - 1
	if a[0] == '-' {
		left = 1
	}
	runes := []rune(a)
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	ans, _ := strconv.Atoi(string(runes))
	if ans >= int(math.Pow(2, 31)) || ans < -int(math.Pow(2, 31)) {
		return 0
	}
	return ans
}