package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	testLowestCommonAncestor()
}
func testLowestCommonAncestor() {
	n1 := &TreeNode{
		Val:   3,
	}
	n2 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 1}
	n1.Left = n2
	n1.Right = n3
	n := lowestCommonAncestor(n1, n2, n3)
	fmt.Println(n.Val)
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}
