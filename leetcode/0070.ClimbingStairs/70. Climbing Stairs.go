package leetcode

var res map[int]int

func climbStairs(n int) int {
	if res == nil {
		res = make(map[int]int)
	}
	if v, ok := res[n]; ok {
		return v
	}

	if n <= 2 {
		res[n] = n
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs1(n int) int {
	res := make(map[int]int)
	res[0] = 1
	res[1] = 1
	res[2] = 2
	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}
