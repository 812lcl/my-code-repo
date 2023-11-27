package leetcode

import "math"

func findSubsequences(nums []int) [][]int {
	c := []int{}
	res := [][]int{}
	dfs(nums, 0, c, &res)
	return res
}

func dfs(nums []int, index int, path []int, res *[][]int) {
	p := len(path)

	if p > 1 {
		*res = append(*res, path)
	}

	last, visited := math.MinInt, make(map[int]bool)

	if p > 0 {
		last = path[p-1]
	}

	for i := index; i < len(nums); i++ {
		num := nums[i]
		if visited[num] {
			continue
		}

		if num >= last {
			newPath := make([]int, p+1)
			copy(newPath, path)
			newPath[p] = num
			dfs(nums, i+1, newPath, res)
			visited[num] = true
		}
	}
}
