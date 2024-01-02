package leetcode

func minOperations(n int) int {
	i := 1
	c := 0
	for i <= n {
		if n&i != 0 && n&(i<<1) != 0 {
			n += i
			c++
		} else if n&i != 0 {
			c++
		}
		i = i << 1
	}
	return c
}
