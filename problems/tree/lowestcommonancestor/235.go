package main

func bstLowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	max, min := 0, 0
	if p.Val > q.Val {
		max = p.Val
		min = q.Val
	}else {
		max = q.Val
		min = p.Val
	}
	return helper(root, max, min)
}
func helper(root *TreeNode, max, min int )*TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > max {
		return helper(root.Left, max, min)
	}
	if root.Val < min {
		return helper(root.Right, max, min)
	}
	return root
}