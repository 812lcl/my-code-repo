package leetcode

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	if m < n {
		a, b = b, a
		m, n = n, m
	}
	c := 0
	var res string
	for i := 0; i < m; i++ {
		sum := 0
		sum += int(a[m-i-1] - '0')
		if i < n {
			sum += int(b[n-i-1] - '0')
		}
		sum += c
		if sum > 1 {
			c = 1
		} else {
			c = 0
		}
		res = itoa(sum%2) + res
	}
	if c == 1 {
		res = string('1') + res
	}
	return res
}

func itoa(num int) string {
	switch num {
	case 0:
		return string('0')
	case 1:
		return string('1')
	default:
		return string('0')
	}
}
