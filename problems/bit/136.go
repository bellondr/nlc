package main

import (
	"os"
	"fmt"
)

type singleNumberCase struct {
	Input []int
	Output int
}

func main() {
	cases := []singleNumberCase{{Input: []int{2,2,1}, Output: 1}}
	for _, ca := range cases {
		out := singleNumber(ca.Input)
		if out != ca.Output {
			fmt.Println("err: case intput：%v, output: %d, expect: %d", ca.Input, out, ca.Output)
			os.Exit(1)
		}
	}
	fmt.Println("all case pass")
}

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}

