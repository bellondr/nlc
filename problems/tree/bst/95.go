package main

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return build(1, n)
}

func build(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	res := []*TreeNode{}
	for start := l; start <= r; start++ {
		left := build(l, start-1)
		right := build(start+1, r)
		for _, l := range left {
			for _, r := range right {
				res = append(res, &TreeNode{
					Val: start,
					Left: l,
					Right: r,
				})
			}
		}
	}
	return res
}
