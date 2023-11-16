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

func reverseWords(s string) string {
	var st stack
	var word string

	for _, v := range s {
		if v != ' ' {
			word = word + string(v)
		} else if len(word) != 0 {
			st.Push(word)
			word = ""
		}
	}
	if len(word) != 0 {
		st.Push(word)
	}

	var res string
	for !st.Empty() {
		if res == "" {
			res = st.Pop()
		} else {
			res = res + " " + st.Pop()
		}
	}
	return res
}

func reverseWords1(s string) string {
	var words []string
	var word string
	for _, v := range s {
		if v != ' ' {
			word = word + string(v)
		} else if len(word) != 0 {
			words = append(words, word)
			word = ""
		}
	}
	if len(word) != 0 {
		words = append(words, word)
	}

	for i, j := 0, len(words)-1; i <= j; {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}
	return strings.Join(words, " ")
}
