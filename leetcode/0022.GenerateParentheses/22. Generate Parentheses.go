package leetcode

import "sort"

func generateParenthesis(n int) []string {
	var p []string
	for i := 0; i < n; i++ {
		p = append(p, "(")
		p = append(p, ")")
	}
	return permuteUnique(p)
}

func permuteUnique(parenthesis []string) []string {
	if len(parenthesis) == 0 {
		return []string{}
	}
	sort.Strings(parenthesis)
	visited, p, res := make([]bool, len(parenthesis)), "", []string{}
	dfs(parenthesis, p, &res, &visited)
	return res

}

func dfs(parenthesis []string, p string, res *[]string, visited *[]bool) {
	if len(p) == len(parenthesis) && isValid(p) {
		tmp := p
		(*res) = append((*res), tmp)
		return
	}

	for i := 0; i < len(parenthesis); i++ {
		if !(*visited)[i] {
			if i > 0 && parenthesis[i] == parenthesis[i-1] && !(*visited)[i-1] {
				continue
			}
			(*visited)[i] = true
			p += parenthesis[i]
			dfs(parenthesis, p, res, visited)
			p = p[:len(p)-1]
			(*visited)[i] = false
		}
	}
}

type stack []rune

func (s *stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *stack) Pop() rune {
	if s.Empty() {
		return 0
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return res
}

func (s *stack) Top() rune {
	if s.Empty() {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *stack) Empty() bool {
	return len(*s) == 0
}

func (s *stack) length() int {
	return len(*s)
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	var st stack
	opposite := make(map[rune]rune)
	opposite[')'] = '('
	opposite[']'] = '['
	opposite['}'] = '{'
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			st.Push(v)
		} else {
			res := st.Pop()
			if res != opposite[v] {
				return false
			}
		}
	}
	return st.length() == 0
}

func generateParenthesis1(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	findGenerateParenthesis(n, n, "", &res)
	return res
}

func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	if lindex == 0 && rindex == 0 {
		*res = append(*res, str)
		return
	}
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}
	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}
