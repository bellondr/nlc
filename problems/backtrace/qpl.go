package main

import "fmt"

var res = [][]string{}

func main()  {
	AllCa(3)
}
func AllCa(a int) {
	dfs(a, a, []string{})
	fmt.Println(res)
}

func dfs(left, right int, tmp []string) {
	if right == 0 {
		res = append(res, tmp[:])
		return
	}
	if left == 0 {
		dfs(left, right-1, append(tmp, "}"))
	} else if left == right {
		dfs(left-1, right, append(tmp, "{"))
	} else {
		dfs(left-1, right, append(tmp, "{"))
		dfs(left, right-1, append(tmp, "}"))
	}
}
