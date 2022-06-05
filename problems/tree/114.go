package main

var arr = []int{}
func preorderTraversal(root *TreeNode) []int {
	arr = make([]int, 0)
	preorder(root)
	return arr
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	preorder(root.Left)
	preorder(root.Right)
}
