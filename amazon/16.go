package main

import (
	"fmt"
	"math"
	"sort"
)

func testThreeSumClosest() {
	res :=threeSumClosest([]int{1,1,1,1}, -100)
	fmt.Println("res is: ", res)
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := sum(nums)
	res += int(math.Abs(float64(target)))
	thelper(nums, 3, 0, target, &res)
	return target + res
}
func sum(nums[]int) int {
	s := 0
	for _, v := range nums{
		s+=v
	}
	return s
}
func thelper(nums []int, n, start, target int, res *int) {
	if n < 2 || len(nums) < n || *res == 0 {
		return
	}
	if n == 2 {
		left, right := start, len(nums) - 1

		for left < right {
			tmp := (nums[left] + nums[right]) - target
			if math.Abs(float64(tmp)) <= math.Abs(float64(*res)) {
				*res = tmp
			}

			if tmp > 0 {
				right--
			} else if tmp < 0 {
				left++
			} else {
				return
			}
		}

	} else {
		for index := start; index < len(nums); index++ {
			thelper(nums, n-1, index+1, target-nums[index], res)
		}
	}
}
