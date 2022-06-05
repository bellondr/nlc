package main

import "fmt"

func main() {
	//data := []int{-1,0,1,2,-1,-4}
	//data := []int{0, 0, 0}
	//data := []int{1,0,-1,0,-2,2}
	data := []int {2,2,2,2,2}
	sortNum(data, 0, len(data) - 1)
	res := nSum(data, 4, 0, 8)
	for i := range res {
		for j := range res[i] {
			fmt.Println(res[i][j])
		}
	}
}

func twoSumBySortedArr(nums []int, target int, left, right int) [][]int {
	res := [][]int{}
	if right - left < 1 {
		return res
	}
	for ; left < right; {
		if nums[left] + nums[right] > target {
			right = right - 1
		} else if nums[left] + nums[right] < target {
			left = left + 1
		} else {
			res = append(res, []int{left, right})
			for ;right > left && nums[right] == nums[right-1]; {
				right = right - 1
			}
			for ;left < right && nums[left] == nums[left+1]; {
				left = left + 1
			}
			left = left + 1
			right = right - 1

		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sortNum(nums, 0, len(nums) - 1)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		twoRes := twoSumBySortedArr(nums, target, i+1, len(nums)-1)
		for _, r := range twoRes {
			print(r[0], r[1], "\n")
			res = append(res, []int{nums[i], nums[r[0]], nums[r[1]]})
		}
		for ;i < len(nums) - 1 && nums[i] == nums[i+1]; {
			i++
		}
	}
	return res
}

func sortNum(nums[]int, left, right int) {
	if left < right {
		i, j := left, right
		key := nums[(left+right)/2]
		for i < j {
			for nums[i] < key {
				i++
			}
			for nums[j] > key {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		if left < j {
			sortNum(nums, left, j)
		}
		if right > i {
			sortNum(nums, i ,right)
		}
	}
}

func nSum(nums []int, n, start, target int) [][]int {
	for (n < 2 || len(nums) < n) {
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
			} else {
				res = append(res, []int{nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for right > left && nums[right] == nums[right - 1] {
					right--
				}
				left++
				right--
			}
		}
		return res
	} else {
		res := [][]int{}
		for index := start; index < len(nums); index++ {
			sub := nSum(nums, n - 1, index + 1, target - nums[index])
			for _, n := range sub {
				res = append(res, append([]int{nums[index]}, n...))
			}
			for index < len(nums) - 1 && nums[index] == nums[index+1] {
				index++
			}
		}
		return res
	}
}

