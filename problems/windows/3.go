func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    maxString := ""
    max, right := 1, 0
    for right=0;right < len(s); right++ {
        for i, c := range maxString {
            if string(c) == string(s[right]) {
                maxString=maxString[i+1:]
            }
        }
        maxString += string(s[right])
        if len(maxString) > max {
            max = len(maxString)
        }
    }
    return max
}

func lengthOfLongestSubstringFast(s string) int {

	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right - left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}