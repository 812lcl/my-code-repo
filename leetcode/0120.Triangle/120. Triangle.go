package leetcode

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, len(triangle))
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(triangle[i]))
		if i == 0 {
			dp[0][0] = triangle[0][0]
		}
		if i != 0 {
			dp[i][0] = dp[i-1][0] + triangle[i][0]
			dp[i][i] = dp[i-1][i-1] + triangle[i][i]
		}
	}
	for i := 2; i < n; i++ {
		for j := 1; j < len(dp[i])-1; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
	}
	res := math.MaxInt
	for i := 0; i < n; i++ {
		if dp[n-1][i] < res {
			res = dp[n-1][i]
		}
	}
	return res
}

func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	n := len(triangle)
	dp := make([]int, n)
	for index := 0; index < len(triangle[0]); index++ {
		dp[index] = triangle[0][index]
	}
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 {
				// 最左边
				dp[j] += triangle[i][0]
			} else if j == len(triangle[i])-1 {
				// 最右边
				dp[j] += dp[j-1] + triangle[i][j]
			} else {
				// 中间
				dp[j] = min(dp[j-1]+triangle[i][j], dp[j]+triangle[i][j])
			}
		}
	}
	res := math.MaxInt
	for i := 0; i < n; i++ {
		if dp[i] < res {
			res = dp[i]
		}
	}
	return res
}

func minimumTotal2(triangle [][]int) int {
	if triangle == nil {
		return 0
	}
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}
