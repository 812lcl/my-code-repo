package leetcode

func findThePrefixCommonArray(A []int, B []int) []int {
	mapA, mapB := make(map[int]bool), make(map[int]bool)
	res := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if i == 0 {
			if A[i] == B[i] {
				res[i] = 1
			} else {
				res[i] = 0
			}
		} else {
			res[i] = res[i-1]
			if A[i] == B[i] && !mapA[B[i]] && !mapB[B[i]] {
				res[i]++
			}
			if mapA[B[i]] && !mapB[B[i]] {
				res[i]++
			}
			if mapB[A[i]] && !mapA[A[i]] {
				res[i]++
			}
		}
		mapA[A[i]] = true
		mapB[B[i]] = true
	}
	return res
}
