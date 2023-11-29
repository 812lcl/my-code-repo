package leetcode

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	visited, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	dfs(nums, p, &res, &visited)
	return res

}

func dfs(nums []int, p []int, res *[][]int, visited *[]bool) {
	if len(p) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, p)
		(*res) = append((*res), tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*visited)[i] {
			(*visited)[i] = true
			p = append(p, nums[i])
			dfs(nums, p, res, visited)
			p = p[:len(p)-1]
			(*visited)[i] = false
		}
	}
}
