package leetcode

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
