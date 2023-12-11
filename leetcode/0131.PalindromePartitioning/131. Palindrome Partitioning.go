package leetcode

func partition(s string) [][]string {
	var res [][]string
	n := len(s)
	if n == 0 {
		return res
	}
	current := make([]string, 0, n)
	dfs(s, 0, current, &res)
	return res
}

func dfs(s string, index int, c []string, res *[][]string) {
	start, end := index, len(s)
	if start == end {
		b := make([]string, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := start; i < len(s); i++ {
		if isPal(s, start, i) {
			dfs(s, i+1, append(c, s[start:i+1]), res)
		}
	}
}

func isPal(str string, s, e int) bool {
	for s < e {
		if str[s] != str[e] {
			return false
		}
		s++
		e--
	}
	return true
}
