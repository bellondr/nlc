package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := -1
	index := 0
	for i, v := range nums {
		if v > max {
			max = v
			index = i
		}
	}
	return &TreeNode{
		Val: max,
		Left: constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}
