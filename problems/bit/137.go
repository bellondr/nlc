package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber137([]int{2, 2, 3, 2}))
}

func singleNumber137(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		cnt := 0
		bit := 1 << i
		for _, n := range nums {
			cnt += bit & n
		}
		if cnt % 3 != 0 {
			res += bit
		}
	}
	if res > 2147483647 {
		return res - 4294967296
	}
	return res
}
