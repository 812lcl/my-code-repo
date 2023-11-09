package structure

type Graph struct {
	nodes  []*Node
	matrix [][]int
}

func (g *Graph) AddNode(value int) {
	node := &Node{Data: value}
	g.nodes = append(g.nodes, node)

	// Add a new column and row to the matrix
	for i := 0; i < len(g.matrix); i++ {
		g.matrix[i] = append(g.matrix[i], 0)
	}
	newRow := make([]int, len(g.nodes))
	g.matrix = append(g.matrix, newRow)
}

func (g *Graph) AddEdge(from, to int) {
	g.matrix[from][to] = 1
	g.matrix[to][from] = 1
}

func (g *Graph) getNode(value int) *Node {
	for _, node := range g.nodes {
		if node.Data == value {
			return node
		}
	}
	return nil
}
