package main

func isValidBST(root *TreeNode) bool {
	return helperV(root, nil, nil)
}

func helperV(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil {
		if root.Val <= min.Val {
			return false
		}
	}
	if max != nil {
		if root.Val >= max.Val {
			return false
		}
	}
	return helperV(root.Left, min, root) && helperV(root.Right, root, max)
}
