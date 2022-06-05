package main

import "fmt"

func minWindow(s string, t string) string {
	left, right := 0, 0
	tm := make(map[string]int, 0)
	for i := range t {
		val, ok := tm[string(t[i])]
		if ok {
			tm[string(t[i])] = val + 1
		} else {
			tm[string(t[i])] = 1
		}
	}

	contain, minLeft, min := false, 0, len(s)
	for right <= len(s) {
		for containsSubString(tm) && left < right {
			if right == len(s) - 1 {
				fmt.Printf("left %d \n", left)
			}
			contain = true
			if right - left < min {
				minLeft = left
				min = right - left
			}
			val, ok := tm[string(s[left])]
			if ok {
				tm[string(s[left])] = val + 1
			}
			//fmt.Printf("%v \n", tm)
			left++
		}
		if right == len(s) {
			break
		}
		val, ok := tm[string(s[right])]
		if ok {
			tm[string(s[right])] = val - 1
		}
		if right == len(s) - 1 {
			fmt.Printf("%v \n", tm)
		}
		right++
	}
	if contain {
		return s[minLeft: minLeft + min]
	}
	return ""

}

func containsSubString(m map[string]int) bool {
	for k := range m {
		if m[k] > 0 {
			return false
		}
	}
	return true
}
