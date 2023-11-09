package leetcode

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	if s.Empty() {
		return 0
	}
	res := (*s)[s.Length()-1]
	*s = (*s)[:s.Length()-1]
	return res
}

func (s *stack) Top() int {
	if s.Empty() {
		return 0
	}
	return (*s)[s.Length()-1]
}

func (s *stack) Empty() bool {
	return s.Length() == 0
}

func (s *stack) Length() int {
	return len(*s)
}

func longestValidParentheses(s string) int {
	var st stack
	res := 0
	st.Push(-1)
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '(' {
			st.Push(i)
		} else {
			st.Pop()
			if st.Empty() {
				st.Push(i)
			} else {
				res = max(res, i-st.Top())
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses1(s string) int {
	var l, r, m int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else {
			r++
		}

		if l == r {
			m = max(m, 2*r)
		} else if r > l {
			l, r = 0, 0
		}
	}
	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		} else {
			r++
		}

		if l == r {
			m = max(m, 2*l)
		} else if r < l {
			l, r = 0, 0
		}
	}
	return m
}
