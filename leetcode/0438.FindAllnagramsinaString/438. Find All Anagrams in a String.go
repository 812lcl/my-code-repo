package leetcode

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	freq := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		freq[p[i]]++
	}
	res := []int{}
	left, right, count := 0, 0, 0
	for left < len(s)-len(p)+1 && right < len(s) {
		if freq[s[right]] > 0 {
			freq[s[right]]--
			count++
			right++
			if count == len(p) {
				res = append(res, left)
			}
		} else {
			freq[s[left]]++
			count--
			left++
		}
	}
	return res
}
