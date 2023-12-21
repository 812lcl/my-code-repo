package leetcode

func findLUSlength(a string, b string) int {
	if len(a) > len(b) {
		return len(a)
	} else if len(a) < len(b) {
		return len(b)
	}
	freq := make(map[byte]int)
	for i := 0; i < len(a); i++ {
		freq[a[i]]++
	}
	for i := 0; i < len(b); i++ {
		if freq[b[i]] == 0 {
			return len(b)
		}
		freq[b[i]]--
	}
	if a != b {
		return len(a)
	}
	return -1
}

func findLUSlength1(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}
