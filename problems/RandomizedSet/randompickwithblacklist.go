package main

import "math/rand"

type Solution struct {
	blackMap map[int]int
	n int
}


func BConstructor(n int, blacklist []int) Solution {
	blackMap := make(map[int]int)
	for i := 0; i < len(blacklist); i++ {
		blackMap[blacklist[i]] = 1

	}
	last := n - 1
	bz := n - len(blacklist)
	for i:= range blacklist {
		if blacklist[i] >= bz {
			continue
		}
		for {
			_, ok := blackMap[last]
			if ok {
				last--
			} else{
				break
			}
		}
		blackMap[blacklist[i]] = last
		last--
	}
	return Solution{
		blackMap: blackMap,
		n:  bz,
	}
}


func (this *Solution) Pick() int {
	r := rand.Intn(this.n)
	val, ok := this.blackMap[r]
	if ok {
		return val
	}
	return r
}