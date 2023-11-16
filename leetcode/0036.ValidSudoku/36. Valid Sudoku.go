package leetcode

func isValidSudoku(board [][]byte) bool {
	visitedCol := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		visitedRow := make([]bool, 9)
		visitedCol[i] = make([]bool, 9)
		for j := 0; j < 9; j++ {
			if string(board[i][j]) != "." {
				if visitedRow[board[i][j]-49] {
					return false
				} else {
					visitedRow[board[i][j]-49] = true
				}
			}
			if string(board[j][i]) != "." {
				if visitedCol[i][board[j][i]-49] {
					return false
				} else {
					visitedCol[i][board[j][i]-49] = true
				}
			}
		}
	}
	for i := 0; i < 9; i++ {
		visitedRow := make([]bool, 9)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if string(board[j+i%3*3][k+i/3*3]) == "." {
					continue
				}
				if visitedRow[board[j+i%3*3][k+i/3*3]-49] {
					return false
				} else {
					visitedRow[board[j+i%3*3][k+i/3*3]-49] = true
				}
			}
		}
	}
	return true
}
