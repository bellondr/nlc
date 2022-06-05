package main

func largestVariance(s string) (ans int) {
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if b == a {
				continue
			}
			diff, diffWithB := 0, -len(s)
			for _, ch := range s {
				if ch == a {
					diff++
					diffWithB++
				} else if ch == b {
					diff--
					diffWithB = diff // 记录包含 b 时的 diff
					diff = max(diff, 0)
				}
				ans = max(ans, diffWithB)
			}
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }

