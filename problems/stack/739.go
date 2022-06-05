package stack

type node struct {
	val int
	index int
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]*node, 0)
	res := make([]int, len(temperatures))

	for i:=len(temperatures) - 1; i>=0; i-- {
		for len(stack) != 0 && stack[len(stack)-1].val <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = 0
		} else if stack[len(stack)- 1].val > temperatures[i] {
			res[i] = stack[len(stack)- 1].index - i
		} else {
			res[i] = 0
		}
		stack = append(stack, &node{val: temperatures[i], index: i})
	}
	return res
}
