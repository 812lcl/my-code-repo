package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	r, res := []int{}, [][]int{}
	dfs(candidates, target, 0, r, &res)
	return res
}

func dfs(candidates []int, target, index int, r []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(r))
		copy(tmp, r)
		(*res) = append((*res), tmp)
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] > target {
			break
		}
		r = append(r, candidates[i])
		dfs(candidates, target-candidates[i], i+1, r, res)
		r = r[:len(r)-1]
	}
}
