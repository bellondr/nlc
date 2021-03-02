package main

import "fmt"

func main() {
	aar := []int{2,3,6,7}
	fmt.Println(combinationSum(aar, 7))
}
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	comBfs(target, candidates, []int{}, &res)
	return res
}

func comBfs(target int, candidates, nums []int, res*[][]int) {
	if target == 0 {
		*res = append(*res, nums)
	}
	if target < 0 {
		return
	}
	for i := range candidates {
		subSetCopy := make([]int, 0)
		subSetCopy = append(subSetCopy, nums...)

		comBfs(target-candidates[i], candidates[i:], append(subSetCopy, candidates[i]), res)
	}
}
