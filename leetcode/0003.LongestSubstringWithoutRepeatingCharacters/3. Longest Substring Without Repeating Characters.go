package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	uniq := make(map[byte]bool)
	res := 0
	for i, j := 0, 0; i < len(s); {
		if uniq[s[j]] {
			uniq[s[i]] = false
			i++
		} else {
			uniq[s[j]] = true
			j++
		}
		if res < j-i {
			res = j - i
		}
		if i+res >= len(s) || j >= len(s) {
			break
		}
	}
	return res
}
