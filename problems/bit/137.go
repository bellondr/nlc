package main

import (
	"fmt"
)

// Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.
// exp1:
//Input: nums = [2,2,3,2]
// Output: 3

// exp2:
// Input: nums = [0,1,0,1,0,1,99]
// Output: 99

func main() {
	fmt.Println(singleNumber137([]int{2, 2, 3, 2}))
}

func singleNumber137(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		cnt := 0
		bit := 1 << i
		for _, n := range nums {
			cnt += bit & n
		}
		if cnt % 3 != 0 {
			res += bit
		}
	}
	if res > 2147483647 {
		return res - 4294967296
	}
	return res
}
