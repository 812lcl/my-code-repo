package leetcode

import "strconv"

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

func evalRPN(tokens []string) int {
	var st stack
	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {
			st.Push(num)
		} else {
			a, b := st.Pop(), st.Pop()
			res := 0
			switch v {
			case "+":
				res = b + a
			case "-":
				res = b - a
			case "*":
				res = b * a
			case "/":
				if a == 0 {
					res = 0
				} else {
					res = b / a
				}
			}
			st.Push(res)
		}
	}
	return st.Pop()
}
