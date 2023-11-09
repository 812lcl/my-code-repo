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

func calPoints(operations []string) int {
	var st stack
	for _, v := range operations {
		num, err := strconv.Atoi(v)
		if err == nil {
			st.Push(num)
		} else {
			switch v {
			case "+":
				a, b := st.Pop(), st.Top()
				res := b + a
				st.Push(a)
				st.Push(res)
			case "D":
				st.Push(2 * st.Top())
			case "C":
				st.Pop()
			}
		}
	}
	var sum int
	for !st.Empty() {
		sum += st.Pop()
	}
	return sum
}
