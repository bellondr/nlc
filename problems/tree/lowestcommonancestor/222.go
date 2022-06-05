package main

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := root, root
	lh, rh := 0, 0
	for l != nil {
		l = l.Left
		lh++
	}
	for r != nil {
		r = r.Right
		rh++
	}
	if lh == rh {
		return int(math.Pow(float64(2), float64(lh))) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}