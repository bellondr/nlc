package main

func findBottomLeftValue(root *TreeNode) int {
	res := make([][]int, 0)
	levelAppend(root, &res, 0)
	return res[len(res)-1][0]
}

func levelAppend(root *TreeNode, res *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{root.Val})
	} else {
		(*res)[level] = append((*res)[level], root.Val)
	}
	levelAppend(root.Left, res, level+1)
	levelAppend(root.Right, res, level+1)
}
