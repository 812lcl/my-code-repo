package leetcode

import "strings"

type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() string {
	if s.Empty() {
		return ""
	}
	res := (*s)[s.Length()-1]
	*s = (*s)[:s.Length()-1]
	return res
}

func (s *stack) Top() string {
	if s.Empty() {
		return ""
	}
	return (*s)[s.Length()-1]
}

func (s *stack) Empty() bool {
	return s.Length() == 0
}

func (s *stack) Length() int {
	return len(*s)
}

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	var st stack
	for _, v := range arr {
		cur := strings.TrimSpace(v)
		if cur == ".." {
			st.Pop()
		} else if cur != "." && len(cur) > 0 {
			st.Push(cur)
		}
	}
	if st.Empty() {
		return "/"
	}
	res := strings.Join(st, "/")
	return "/" + res
}
