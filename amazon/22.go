package main

import "fmt"

func testGenerateParenthesis() {
	generateParenthesis(3)
}
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	genHelper("", &res, 0, 0, n)
	fmt.Println(res)
	return res
}

func genHelper(tmp string, res*[]string, left, right, n int) {
	if left > n {
		return
	}
	if right > n {
		return
	}
	if left == n && right == n {
		*res = append(*res, tmp)
		return
	}

	if left >= right {
		genHelper(tmp + "(", res, left+1, right, n)
		genHelper(tmp + ")", res, left, right+1, n)
	}
}
