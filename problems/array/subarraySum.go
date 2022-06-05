package main

import "fmt"

func subarraySum(nums []int, k int) int {
	sm := make(map[int]int, 0)
	sm[0] = 1

	count := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		left := sum - k
		val, ok := sm[left]
		if ok {
			count += val
		}
		val2, ok := sm[sum]
		if ok {
			sm[sum] = val2+1
		} else {
			sm[sum] = 1
		}
	}
	fmt.Println(sm)
	return count
}
