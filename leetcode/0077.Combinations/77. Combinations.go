package leetcode

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	var res [][]int
	dfs(n, k, 1, []int{}, &res)
	return res
}

func dfs(n, k, index int, c []int, res *[][]int) {
	if len(c) == k {
		tmp := make([]int, k)
		copy(tmp, c)
		*res = append(*res, tmp)
	}
	for i := index; i <= n; i++ {
		c = append(c, i)
		dfs(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}
}
