package main

import "sort"
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return helper(nums, 3, 0, 0)
}

func helper(nums[]int, n, start, target int)[][]int {
	if n < 2 || len(nums) < n {
		return [][]int{}
	}
	if n == 2 {
		res := [][]int{}
		left, right := start, len(nums) - 1
		for left < right {
			if nums[left] + nums[right] > target {
				right--
			} else if nums[left] + nums[right] < target {
				left++
			} else if nums[left] + nums[right] == target {
				res = append(res, []int{nums[left], nums[right]})
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				right--
				left++
			}
		}
		return res
	} else {
		res := [][]int{}
		for index := start; index < len(nums); index++ {
			tmp := helper(nums, n - 1, index + 1, target-nums[index])
			for _, v := range tmp {
				res = append(res, append([]int{nums[index]}, v...))
			}
			for index < len(nums) - 1&& nums[index] == nums[index+1] {
				index++
			}
		}
		return res
	}
}
