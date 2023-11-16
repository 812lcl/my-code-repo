package leetcode

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if searchWord(board, word, &visited, 0, i, j) {
				return true
			}
		}
	}
	return false
}

var dir = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func searchWord(board [][]byte, word string, visited *[][]bool, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		(*visited)[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInboard(board, nx, ny) && !(*visited)[nx][ny] && searchWord(board, word, visited, index+1, nx, ny) {
				return true
			}
		}
		(*visited)[x][y] = false
	}
	return false
}

func isInboard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
