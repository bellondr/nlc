package main

func testSplitArray() {
	nums := []int{7,2,5,10,8}
	m := 2
	print(splitArray(nums, m))
}
func splitArray(nums []int, m int) int {
	left, right := 0, 0
	for i := range nums {
		if nums[i] > left {
			left = nums[i]
		}
		right += nums[i]
	}
	for left < right {
		mid := left + (right - left)/2
		sn := splitNums(nums, mid, m)
		if sn > 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func splitNums(nums []int, l, m int) int {
	step := 1
	cap := l
	for i := range nums {
		if cap >= nums[i] {
			cap -= nums[i]
		} else {
			step++
			cap = l - nums[i]
			if cap < 0 {
				return 1
			}
		}
	}
	return step - m
}
