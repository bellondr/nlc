package main

func widthOfBinaryTree(root *TreeNode) int {
	log := make(map[int][]int)
	helper662(root, 0, 1, log)
	max := -1
	for _, v := range log {
		if v[len(v)-1] - v[0] > max {
			max = v[len(v)-1]- v[0]
		}
	}
	return max + 1
}

func helper662(root *TreeNode, level, num int, log map[int][]int) {
	if root == nil {
		return
	}

	helper662(root.Left, level+1, num*2, log)
	helper662(root.Right, level+1, num*2+1, log)
	v, ok := log[level]
	if !ok {
		log[level] = []int{num}
	} else {
		log[level] = append(v, num)
	}
}
