package leetcode

import "sort"

func reorganizeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	freq := make(map[byte]int)
	reverseFreq := make(map[int][]byte)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	for k, v := range freq {
		reverseFreq[v] = append(reverseFreq[v], k)
	}

	var keys []int
	for k := range reverseFreq {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	if keys[0] > (len(s)+1)/2 {
		return ""
	}

	var tmp string
	for _, key := range keys {
		for i := 0; i < len(reverseFreq[key]); i++ {
			for j := 0; j < key; j++ {
				tmp += string(reverseFreq[key][i])
			}
		}
	}
	var res string
	for i, j := 0, (len(tmp)-1)/2+1; i <= (len(tmp)-1)/2; i++ {
		res += string(tmp[i])
		if j < len(tmp) {
			res += string(tmp[j])
		}
		j++
	}
	return res
}
