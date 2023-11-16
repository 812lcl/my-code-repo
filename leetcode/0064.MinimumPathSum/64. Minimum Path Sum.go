package leetcode

func minPathSum(grid [][]int) int {
	if len(grid) == 1 {
		var sum int
		for _, num := range grid[0] {
			sum += num
		}
		return sum
	}
	if len(grid[0]) == 1 {
		var sum int
		for i := 0; i < len(grid); i++ {
			sum += grid[i][0]
		}
		return sum
	}
	var grid1 [][]int
	grid2 := make([][]int, len(grid))
	for i, v := range grid {
		if i == 0 {
			continue
		}
		grid1 = append(grid1, v)
	}
	res := grid[0][0]
	for i, v := range grid {
		grid2[i] = append(grid2[i], v[1:]...)
	}
	res1 := minPathSum(grid1)
	res2 := minPathSum(grid2)
	if res1 < res2 {
		return res + res1
	}
	return res + res2
}

func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]

}
