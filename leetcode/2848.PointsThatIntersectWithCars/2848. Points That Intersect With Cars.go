package leetcode

func numberOfPoints(nums [][]int) int {
	uniq := make(map[int]bool)
	for _, num := range nums {
		for i := num[0]; i <= num[1]; i++ {
			uniq[i] = true
		}
	}
	return len(uniq)
}

func numberOfPoints1(nums [][]int) int {
	var h [102]int
	for _, num := range nums {
		h[num[0]]++
		h[num[1]+1]--
	}
	res, cars := 0, 0
	for _, k := range h {
		cars += k
		if cars > 0 {
			res++
		}
	}
	return res
}
