package main

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	helper543(root, &max)
	return max
}

func helper543(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		if *max < 1 {
			*max = 1
		}
		return 1
	}
	l := helper543(node.Left, max)
	r := helper543(node.Right, max)
	if l + r + 1 > *max {
		*max = l + r + 1
	}
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
