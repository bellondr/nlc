package main

func findKthNumber(n int, k int) int {
	p := 1
	prefix := 1
	for p < k {
		cnt := getCount(prefix, n)
		if p + cnt > k {
			prefix *= 10
			p++
		} else {
			prefix++
			p += cnt
		}
	}
	return prefix
}

func getCount(prefix, n int) int {
	cnt := 0
	a := prefix
	b := prefix+1
	for ; a <= n ; {
		cnt += min(n+1, b) - a
		a *= 10
		b *= 10
	}
	return cnt
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}