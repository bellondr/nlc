package dp

import "strconv"

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dp[1] = 1
	for i := 1; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i-1: i+1])
		if num >= 10 && num < 26 {
			dp[i+1] += dp[i-1]
		}
		if num == 0 {
			return 0
		}
		if num % 10 != 0 {
			dp[i+1] += dp[i]
		}
	}
	return dp[len(s)]
}

