package leetcode

func removeOuterParentheses(s string) string {
	res := ""
	count := 0
	for _, v := range s {
		if v == '(' {
			if count > 0 {
				res += string(v)
			}
			count++
		} else if v == ')' {
			count--
			if count > 0 {
				res += string(v)
			}
		} else {
			res += string(v)
		}
	}
	return res
}
