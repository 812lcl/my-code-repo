package leetcode

import "strconv"

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	if isValid(int(s[0])-'0', 1) {
		dp[1] = 1
	}
	for i := 2; i <= n; i++ {
		n, _ := strconv.Atoi(string(s[i-1]))
		if isValid(n, 1) {
			dp[i] += dp[i-1]
		}
		n, _ = strconv.Atoi(string(s[i-2 : i]))
		if isValid(n, 2) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func isValid(num, len int) bool {
	if len == 1 {
		return num >= 1 && num <= 9
	} else {
		return num >= 10 && num <= 26
	}
}
