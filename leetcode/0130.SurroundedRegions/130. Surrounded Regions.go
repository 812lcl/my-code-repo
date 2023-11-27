package leetcode

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 {
				if board[i][j] == 'O' {
					dfs(board, i, j)
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

var dir = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[x])-1 {
		return
	}
	if board[x][y] == 'O' {
		board[x][y] = '*'
		for i := 0; i < 4; i++ {
			dfs(board, x+dir[i][0], y+dir[i][1])
		}
	}
}
