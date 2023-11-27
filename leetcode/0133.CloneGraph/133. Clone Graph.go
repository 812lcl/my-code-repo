package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	dfs(node, visited)

	return visited[node]
}

func dfs(node *Node, visited map[*Node]*Node) {
	if node == nil {
		return
	}

	newNode := new(Node)
	newNode.Val = node.Val

	visited[node] = newNode

	for _, child := range node.Neighbors {
		if _, exist := visited[child]; !exist {
			dfs(child, visited)
		}

		newNode.Neighbors = append(newNode.Neighbors, visited[child])
	}
}
