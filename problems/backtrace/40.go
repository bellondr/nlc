package main

import "fmt"

func main()  {
	aar := []int{10,1,2,7,6,1,5}
	fmt.Println(combinationSum4(aar, 8))
}


func combinationSum4(candidates []int, target int) [][]int {
	sort(candidates)
	res := make([][]int, 0)
	comBfs2(target, candidates, []int{}, &res)
	return res
}

func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j ++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
func comBfs2(target int, candidates, nums []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, nums)
		return
	}
	if target < 0 {
		return
	}
	for i := range candidates {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		sub := make([]int, 0)
		sub = append(sub, nums...)
		comBfs2(target-candidates[i], candidates[i+1:], append(sub, candidates[i]), res)
	}
}
