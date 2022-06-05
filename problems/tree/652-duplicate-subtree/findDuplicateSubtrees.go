package main

import "fmt"
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	dtn := make([]*TreeNode, 0)
	dfs(root, make(map[string]int), &dtn)
	return dtn
}

func dfs(root *TreeNode, hasArr map[string]int, dtn *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	l := dfs(root.Left, hasArr, dtn)
	r := dfs(root.Right, hasArr, dtn)
	val := fmt.Sprintf("(%s)(%d)(%s)", l, root.Val, r)
	hasArr[val]++
	if hasArr[val] == 2 {
		*dtn = append(*dtn, root)
	}
	return val
}