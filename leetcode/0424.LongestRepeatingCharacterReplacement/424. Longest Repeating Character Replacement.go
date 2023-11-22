package leetcode

func characterReplacement(s string, k int) int {
	uniq := make(map[byte]int)
	res, dup := 0, 0
	for i := 0; i < len(s); i++ {
		if uniq[s[i]] != 0 {
			continue
		}
		uniq[s[i]]++
	}
	for c := range uniq {
		dup = 0
		for left, right := 0, 0; right < len(s); {
			if s[right] == c {
				right++
			} else if dup < k {
				right++
				dup++
			} else if s[left] == c {
				left++
			} else {
				left++
				dup--
			}
			if res < right-left {
				res = right - left
			}
		}
	}
	return res
}
