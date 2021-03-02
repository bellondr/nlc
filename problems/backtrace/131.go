package main

import "fmt"

func main()  {
	fmt.Println(partition("aab"))
}

func partition(s string) [][]string {
	res := make([][]string, 0)
	partHelper(s, []string{}, &res)
	return res
}

func partHelper(s string, tmp []string, res *[][]string) {
	if len(s) == 0 {
		*res = append(*res, append([]string{}, tmp...))
		return
	}
	for i := 1 ; i <= len(s); i++ {
		if isPar(s[:i]) {
			tp := append([]string{}, tmp...)
			tp = append(tp, s[:i])
			partHelper(s[i:], tp, res)
		}
	}
}

func isPar(s string) bool {
	if len(s) == 0 {
		return false
	}
	begin, end := 0, len(s) - 1
	for ; begin < end ; {
		if s[begin] != s[end] {
			return false
		}
		begin++
		end--
	}
	return true
}
