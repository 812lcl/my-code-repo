package leetcode

func findWords(board [][]byte, words []string) []string {
	var res []string
	for _, word := range words {
		if exist(board, word) {
			res = append(res, word)
		}
	}
	return res
}

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

func findWords1(board [][]byte, words []string) []string {
	trie := new(TrieNode)
	for _, word := range words {
		trie.insert(word)
	}

	result := make([]string, 0, len(words))
	for i, row := range board {
		for j, ch := range row {
			idx := ch - 'a'
			if tn := trie.Children[idx]; tn != nil {
				dfs(board, i, j, trie, &result)
			}
		}
	}

	return result
}

func dfs(board [][]byte, i, j int, node *TrieNode, result *[]string) {
	letter := board[i][j]
	node = node.Children[letter-'a']
	if node == nil {
		return
	} else if node.Word != "" {
		*result = append(*result, node.Word)
		node.Word = ""
	}

	board[i][j] = '.'
	for _, d := range dir {
		x, y := i+d[0], j+d[1]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] != '.' {
			dfs(board, x, y, node, result)
		}
	}
	board[i][j] = letter
}

type TrieNode struct {
	Children [26]*TrieNode
	Word     string
}

func (tn *TrieNode) insert(word string) {
	for _, ch := range word {
		idx := ch - 'a'
		if tn.Children[idx] == nil {
			tn.Children[idx] = new(TrieNode)
		}
		tn = tn.Children[idx]
	}
	tn.Word = word
}
