package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		min := getMin(root.Right)
		root.Right = deleteNode(root.Right, min.Val)
		min.Left = root.Left
		min.Right = root.Right
		root = min
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMin(tmp *TreeNode) *TreeNode{
	root := tmp
	if root == nil {
		return root
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}