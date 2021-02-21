package main

func singleNumber260(nums []int) []int {
	res := 0

	for _, n := range nums {
		res = res ^n
	}
	h := 1
	for {
		if res & h == 0 {
			h <<= 1
		} else {
			break
		}
	}
	a, b := 0, 0
	for _, n := range nums {
		if n & h == 0 {
			a ^= n
		} else {
			b ^= n
		}
	}
	return []int{a,b}
}
