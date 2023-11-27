package leetcode

// DFS
func subsets(nums []int) [][]int {
	var c []int
	var res [][]int
	for k := 0; k <= len(nums); k++ {
		dfs(nums, k, 0, c, &res)
	}
	return res
}

func dfs(nums []int, k, index int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(nums); i++ {
		dfs(nums, k, i+1, append(c, nums[i]), res)
	}
}

// BFS
func subsets1(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for i := 0; i < len(nums); i++ {
		length := len(res)

		for j := 0; j < length; j++ {
			item := res[j]

			tmp := make([]int, len(item))
			copy(tmp, item)
			tmp = append(tmp, nums[i])

			res = append(res, tmp)
		}
	}

	return res
}

// BIT
func subsets2(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 1<<n)
	for x := 0; x < 1<<n; x++ {
		res[x] = make([]int, 0, n)
		for b := 0; b < n; b++ {
			if x&(1<<b) > 0 {
				res[x] = append(res[x], nums[b])
			}
		}
	}
	return res
}
