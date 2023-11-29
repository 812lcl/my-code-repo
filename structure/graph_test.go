package structure

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGraph_AddNode(t *testing.T) {
	g := &Graph{}
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	node := g.getNode(3)

	// Verify the nodes are added correctly
	wantNodes := []*Node{
		{Data: 1},
		{Data: 2},
		{Data: 3},
	}
	if !reflect.DeepEqual(g.nodes, wantNodes) {
		t.Errorf("AddNode() failed to add nodes correctly. Got: %v, Want: %v", g.nodes, wantNodes)
	}
	if !reflect.DeepEqual(g.nodes[2], node) {
		t.Errorf("AddNode() failed to add nodes correctly. Got: %v, Want: %v", g.nodes, wantNodes)
	}

	// Verify the matrix is updated correctly
	wantMatrix := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	if !reflect.DeepEqual(g.matrix, wantMatrix) {
		t.Errorf("AddNode() failed to update matrix correctly. Got: %v, Want: %v", g.matrix, wantMatrix)
	}
}
func TestGraphAdjList(t *testing.T) {
	/* 初始化无向图 */
	v := ValsToVets([]int{1, 3, 2, 5, 4})
	edges := [][]Vertex{{v[0], v[1]}, {v[0], v[3]}, {v[1], v[2]}, {v[2], v[3]}, {v[2], v[4]}, {v[3], v[4]}}
	graph := newGraphAdjList(edges)
	fmt.Println("初始化后，图为:")
	graph.print()

	/* 添加边 */
	// 顶点 1, 2 即 v[0], v[2]
	graph.addEdge(v[0], v[2])
	fmt.Println("\n添加边 1-2 后，图为")
	graph.print()
	fmt.Println(graph.size())

	/* 删除边 */
	// 顶点 1, 3 即 v[0], v[1]
	graph.removeEdge(v[0], v[1])
	fmt.Println("\n删除边 1-3 后，图为")
	graph.print()

	/* 添加顶点 */
	v5 := NewVertex(6)
	graph.addVertex(v5)
	fmt.Println("\n添加顶点 6 后，图为")
	graph.print()

	/* 删除顶点 */
	// 顶点 3 即 v[1]
	graph.removeVertex(v[1])
	fmt.Println("\n删除顶点 3 后，图为")
	graph.print()
}

func TestGraphAdjMat(t *testing.T) {
	/* 初始化无向图 */
	// 请注意，edges 元素代表顶点索引，即对应 vertices 元素索引
	vertices := []int{1, 3, 2, 5, 4}
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}, {2, 4}, {3, 4}}
	graph := newGraphAdjMat(vertices, edges)
	fmt.Println("初始化后，图为:")
	graph.print()

	/* 添加边 */
	// 顶点 1, 2 的索引分别为 0, 2
	graph.addEdge(0, 2)
	fmt.Println("添加边 1-2 后，图为")
	graph.print()

	/* 删除边 */
	// 顶点 1, 3 的索引分别为 0, 1
	graph.removeEdge(0, 1)
	fmt.Println("删除边 1-3 后，图为")
	graph.print()

	/* 添加顶点 */
	graph.addVertex(6)
	fmt.Println("添加顶点 6 后，图为")
	graph.print()

	/* 删除顶点 */
	// 顶点 3 的索引为 1
	graph.removeVertex(1)
	fmt.Println("删除顶点 3 后，图为")
	graph.print()
}

// NewVertex 构造函数
func NewVertex(val int) Vertex {
	return Vertex{
		Val: val,
	}
}

// ValsToVets Generate a vertex list tree given an array
func ValsToVets(vals []int) []Vertex {
	vets := make([]Vertex, len(vals))
	for i := 0; i < len(vals); i++ {
		vets[i] = NewVertex(vals[i])
	}
	return vets
}
