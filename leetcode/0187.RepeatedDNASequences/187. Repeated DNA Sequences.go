package leetcode

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	cache := make(map[string]int)
	res := []string{}
	for i := 0; i <= len(s)-10; i++ {
		cur := string(s[i : i+10])
		if cache[cur] == 1 {
			res = append(res, cur)
		}
		cache[cur]++
	}
	return res
}
