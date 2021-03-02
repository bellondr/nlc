package main

import "fmt"

//# 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//#
//# 说明：解集不能包含重复的子集。
//#
//# 示例:
//#
//# 输入: nums = [1,2,3]
//# 输出:
//# [
//#   [3],
//#   [1],
//#   [2],
//#   [1,2,3],
//#   [1,3],
//#   [2,3],
//#   [1,2],
//#   []
//# ]

func main() {
	fmt.Println(subsets([]int{1,2,3}))
}
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	subDfs(nums, []int{}, &res)
	return res
}

func subDfs(nums, tmp []int, res *[][]int) {
	*res = append(*res, tmp)
	for i := range nums {
		subDfs(append([]int{}, nums[i+1:]...), append(tmp, nums[i]), res)
	}
}
