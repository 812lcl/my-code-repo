package leetcode

func uniqUser(users [][]string) int {
	emails, names := make(map[string][]string), make(map[string][]string)
	for _, user := range users {
		emails[user[1]] = append(emails[user[1]], user[0])
		names[user[0]] = append(emails[user[0]], user[1])
	}
	visited := make(map[string]bool)
	res := 0
	for _, user := range users {
		if visited[user[0]] || visited[user[1]] {
			continue
		}
		res++
		visited[user[0]] = true
		visited[user[1]] = true
		dfs(emails, names, user[1], user[0], visited)
	}
	return res
}

func dfs(emails, names map[string][]string, email, name string, visited map[string]bool) {
	if visited[email] && visited[name] {
		return
	}
	for _, v := range emails[email] {
		visited[v] = true
		dfs(emails, names, email, v, visited)
	}
	for _, v := range names[name] {
		visited[v] = true
		dfs(emails, names, v, name, visited)
	}
}
