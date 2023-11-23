package leetcode

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	freq := make(map[int]int)
	reverseFreq := make(map[int][]int)
	for _, num := range barcodes {
		freq[num]++
	}
	for k, v := range freq {
		reverseFreq[v] = append(reverseFreq[v], k)
	}
	var keys []int
	for k := range reverseFreq {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	codes := []int{}
	for _, k := range keys {
		for i := 0; i < len(reverseFreq[k]); i++ {
			for j := 0; j < k; j++ {
				codes = append(codes, reverseFreq[k][i])
			}
		}
	}
	res := []int{}
	j := (len(codes)-1)/2 + 1
	for i := 0; i <= (len(codes)-1)/2; i++ {
		res = append(res, codes[i])
		if j < len(codes) {
			res = append(res, codes[j])
		}
		j++
	}

	return res
}
