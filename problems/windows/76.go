package main

import "fmt"

func minWindow(s string, t string) string {
	tag := false
	left, right := 0, 0
	bitMap := make(map[string]int)
	logMap := make(map[string]int)

	for _, tc := range t {
		key := string(tc)
		v, ok := bitMap[key]
		if ok {
			bitMap[key] = v + 1
			//logMap[key] = v + 1
		} else {
			bitMap[key] = 1
		}
		logMap[key] = 0
	}
	res := ""

	for left < len(s) {

		if tag {
			v, ok := logMap[string(s[left])]
			if !ok {
				left++
				continue
			}
			if res == "" {
				res = s[left:right]
			} else {
				if right-left < len(res) {
					res = s[left:right]
				}
			}
			logMap[string(s[left])] = v - 1
			if logMap[string(s[left])] < bitMap[string(s[left])] {
				tag = false
			}
			left++
		} else {
			if right >= len(s) {
				break
			}
			v, ok := logMap[string(s[right])]
			if !ok {
				right++
				continue
			}
			logMap[string(s[right])] = v + 1
			tag = true
			for k, v := range logMap {
				if bitMap[k] > v {
					tag = false
					break
				}
			}
			right++
		}
	}
	return res
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
