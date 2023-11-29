package leetcode

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
