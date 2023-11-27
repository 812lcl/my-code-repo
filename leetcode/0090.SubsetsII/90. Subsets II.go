package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var c []int
	var res [][]int
	sort.Ints(nums)
	for i := 0; i <= len(nums); i++ {
		dfs(nums, i, 0, c, &res)
	}
	return res
}

func dfs(nums []int, k, index int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, k)
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		dfs(nums, k, i+1, append(c, nums[i]), res)
	}
}
