package main

import (
	"sort"
)

type Helpstruct struct {
	Val int
	Index int
}

func advantageCount(nums1 []int, nums2 []int) []int {
	hss := make([]Helpstruct, len(nums2))
	for i := range nums2 {
		hss[i] = Helpstruct{
			Val: nums2[i],
			Index: i,
		}
	}
	sort.Slice(hss, func(i, j int) bool {
		return hss[i].Val < hss[j].Val
	})
	sort.Ints(nums1)
	left, right := 0, len(nums1) - 1
	index := right
	res := make([]int, len(nums1))
	for left <= right {
		if nums1[right] > hss[index].Val {
			res[hss[index].Index] = nums1[right]
			right--
		} else {
			res[hss[index].Index] = nums1[left]
			left++
		}
		index--
	}
	return res
}
