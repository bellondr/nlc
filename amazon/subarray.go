package main

import "fmt"

/*

Give a binary array nums, i.e. only contains 0 or 1, and an integer goal.
Return the count of non-empty subarrays whose sum is goal.
Note: subarray means a continuous subset of array.
Sample 1:
Input: nums = [1,0,1,0,1], goal = 2
Output: 4
Reason: subarrays are [1,0,1], [1,0,1,0], [0,1,0,1], [1,0,1]
Sample 2:
Input: nums = [0,0,0,0,0], goal = 0
Output: 15

 */

func testSo()  {
	res := solution2([]int{1,0,1,0,1}, 2)
	fmt.Println(res)

	res = solution2([]int{0,0,0,0,0}, 0)
	fmt.Println(res)
}
func solution(array[]int, goal int)int {
	left, right := 0, 0
	count, sum := 0, 0
	for right < len(array) {
		sum += array[right]
		if sum == goal {
			count++
			right++
		} else if sum > goal {
			for left <= right && sum > goal {
				sum -= array[left]
				left++
			}
			if sum == goal {
				count++
			}
			right++
		} else if sum < goal {
			right++
		}
	}
	return count
}

func solution2(array[]int, goal int)int{
	preSum := make([]int, len(array)+1)
	for i := range array {
		preSum[i+1] = array[i] + preSum[i]
	}
	cnt := 0
	for i := 0; i <= len(array); i++ {
		for j := i+1; j <= len(array); j++ {
			if preSum[j] - preSum[i] == goal {
				cnt++
			}
			if preSum[j] - preSum[i] > goal {
				break
			}
		}
	}
	return cnt
}