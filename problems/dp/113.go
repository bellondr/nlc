package dp

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)

	pdfs(root, []int{}, targetSum, &res)
	return res
}

func pdfs(node *TreeNode, nums[]int, target int, res *[][]int) {
	if node == nil {
		return
	}
	nums = append(nums, node.Val)
	target = target - node.Val
	if node.Left == nil && node.Right == nil {
		if target == 0 {
			c := make([]int, len(nums))
			copy(c, nums)
			*res = append(*res, c)
		}
		return
	}
	pdfs(node.Left, nums, target, res)
	pdfs(node.Right, nums, target, res)
}

