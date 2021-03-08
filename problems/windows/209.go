package main

import (
	"fmt"
)

func main() {
	test()
}

func minSubArrayLen(s int, nums []int) int {
	left := 0
	right := 0
	max := len(nums)
	valid := false
	moveRight := true
	sum := 0
	for right < len(nums) {
		if moveRight {
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}
		if sum >= s {
			tmp := right - left+1
			if tmp < max {
				max = tmp
			}

			valid = true
			moveRight=false
		}else {
			moveRight=true
		}

		if moveRight {
			right++
		}
	}
	if valid {
		return max
	}
	return 0
}