package leetcode

type stack []rune

func (s *stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *stack) Pop() rune {
	if s.Empty() {
		return 0
	}
	res := (*s)[s.Length()-1]
	*s = (*s)[:s.Length()-1]
	return res
}

func (s *stack) Top() rune {
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

func minAddToMakeValid(s string) int {
	var st stack
	for _, v := range s {
		if v == '(' {
			st.Push(v)
		} else if v == ')' && !st.Empty() && st.Top() == '(' {
			st.Pop()
		} else {
			st.Push(v)
		}
	}
	return st.Length()
}
