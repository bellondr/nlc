package main

import "fmt"

//# 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//#
//# 示例:
//#
//# 输入: [1,2,3]
//# 输出:
//# [
//#   [1,2,3],
//#   [1,3,2],
//#   [2,1,3],
//#   [2,3,1],
//#   [3,1,2],
//#   [3,2,1]
//# ]

func main()  {
	fmt.Println(permute([]int{1,2,3}))
}

func permute(nums []int) [][]int {
	res := make([][]int,0)
	pDfs(nums, 0, &res)
	return res
}

func pDfs(nums []int, start int, res*[][]int) {
	if start == len(nums) {
		*res = append(*res, append([]int{}, nums...))
		return
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		pDfs(nums, start+1, res)
		nums[i], nums[start] = nums[start], nums[i]
	}
}
