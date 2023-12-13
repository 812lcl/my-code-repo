package leetcode

func xorQueries(arr []int, queries [][]int) []int {
	var res []int
	for _, nums := range queries {
		result := arr[nums[0]]
		for i := nums[0] + 1; i <= nums[1]; i++ {
			result ^= arr[i]
		}
		res = append(res, result)
	}
	return res
}

func xorQueries1(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr))
	xors[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		xors[i] = arr[i] ^ xors[i-1]
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = xors[q[1]]
		if q[0] > 0 {
			res[i] ^= xors[q[0]-1]
		}
	}
	return res
}
