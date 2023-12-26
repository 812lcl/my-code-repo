package leetcode

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if nums1[r] == nums2[c] {
				dp[r+1][c+1] = dp[r][c] + 1
			} else {
				dp[r+1][c+1] = max(dp[r][c+1], dp[r+1][c])
			}
		}
	}

	return dp[m][n]
}
