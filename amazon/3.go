package main



func lengthOfLongestSubstring(s string) int {
	index := make(map[int]int)
	left, right := 0, 0
	ans := 0
	for right < len(s) {
		num := int(s[right] - 'a')
		val, ok := index[num]
		if ok {
			if left >= val {
				delete(index, num)
			}
			left = maxInt(left, val+1)
			ans = maxInt(ans, right-left+1)
		} else {
			ans = maxInt(ans, right-left+1)
		}
		index[num] = right
		right++
	}
	return ans

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
