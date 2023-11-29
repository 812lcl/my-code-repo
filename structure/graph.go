package structure

import (
	"fmt"
	"strconv"
	"strings"
)

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

/* 基于邻接矩阵实现的无向图类 */
type graphAdjMat struct {
	// 顶点列表，元素代表“顶点值”，索引代表“顶点索引”
	vertices []int
	// 邻接矩阵，行列索引对应“顶点索引”
	adjMat [][]int
}

/* 构造函数 */
func newGraphAdjMat(vertices []int, edges [][]int) *graphAdjMat {
	// 添加顶点
	n := len(vertices)
	adjMat := make([][]int, n)
	for i := range adjMat {
		adjMat[i] = make([]int, n)
	}
	// 初始化图
	g := &graphAdjMat{
		vertices: vertices,
		adjMat:   adjMat,
	}
	// 添加边
	// 请注意，edges 元素代表顶点索引，即对应 vertices 元素索引
	for i := range edges {
		g.addEdge(edges[i][0], edges[i][1])
	}
	return g
}

/* 获取顶点数量 */
func (g *graphAdjMat) size() int {
	return len(g.vertices)
}

/* 添加顶点 */
func (g *graphAdjMat) addVertex(val int) {
	n := g.size()
	// 向顶点列表中添加新顶点的值
	g.vertices = append(g.vertices, val)
	// 在邻接矩阵中添加一行
	newRow := make([]int, n)
	g.adjMat = append(g.adjMat, newRow)
	// 在邻接矩阵中添加一列
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
}

/* 删除顶点 */
func (g *graphAdjMat) removeVertex(index int) {
	if index >= g.size() {
		return
	}
	// 在顶点列表中移除索引 index 的顶点
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	// 在邻接矩阵中删除索引 index 的行
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)
	// 在邻接矩阵中删除索引 index 的列
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}
}

/* 添加边 */
// 参数 i, j 对应 vertices 元素索引
func (g *graphAdjMat) addEdge(i, j int) {
	// 索引越界与相等处理
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Printf("%s", "Index Out Of Bounds Exception")
	}
	// 在无向图中，邻接矩阵沿主对角线对称，即满足 (i, j) == (j, i)
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

/* 删除边 */
// 参数 i, j 对应 vertices 元素索引
func (g *graphAdjMat) removeEdge(i, j int) {
	// 索引越界与相等处理
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Printf("%s", "Index Out Of Bounds Exception")
	}
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

/* 打印邻接矩阵 */
func (g *graphAdjMat) print() {
	fmt.Printf("\t顶点列表 = %v\n", g.vertices)
	fmt.Printf("\t邻接矩阵 = \n")
	for i := range g.adjMat {
		fmt.Printf("\t\t\t%v\n", g.adjMat[i])
	}
}

type Vertex struct {
	Val int
}

/* 基于邻接表实现的无向图类 */
type graphAdjList struct {
	// 邻接表，key: 顶点，value：该顶点的所有邻接顶点
	adjList map[Vertex][]Vertex
}

/* 构造函数 */
func newGraphAdjList(edges [][]Vertex) *graphAdjList {
	g := &graphAdjList{
		adjList: make(map[Vertex][]Vertex),
	}
	// 添加所有顶点和边
	for _, edge := range edges {
		g.addVertex(edge[0])
		g.addVertex(edge[1])
		g.addEdge(edge[0], edge[1])
	}
	return g
}

/* 获取顶点数量 */
func (g *graphAdjList) size() int {
	return len(g.adjList)
}

/* 添加边 */
func (g *graphAdjList) addEdge(vet1 Vertex, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error")
	}
	// 添加边 vet1 - vet2, 添加匿名 struct{},
	g.adjList[vet1] = append(g.adjList[vet1], vet2)
	g.adjList[vet2] = append(g.adjList[vet2], vet1)
}

/* 删除边 */
func (g *graphAdjList) removeEdge(vet1 Vertex, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error")
	}
	// 删除边 vet1 - vet2
	g.adjList[vet1] = DeleteSliceElms(g.adjList[vet1], vet2)
	g.adjList[vet2] = DeleteSliceElms(g.adjList[vet2], vet1)
}

/* 添加顶点 */
func (g *graphAdjList) addVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if ok {
		return
	}
	// 在邻接表中添加一个新链表
	g.adjList[vet] = make([]Vertex, 0)
}

/* 删除顶点 */
func (g *graphAdjList) removeVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if !ok {
		panic("error")
	}
	// 在邻接表中删除顶点 vet 对应的链表
	delete(g.adjList, vet)
	// 遍历其他顶点的链表，删除所有包含 vet 的边
	for v, list := range g.adjList {
		g.adjList[v] = DeleteSliceElms(list, vet)
	}
}

/* 打印邻接表 */
func (g *graphAdjList) print() {
	var builder strings.Builder
	fmt.Printf("邻接表 = \n")
	for k, v := range g.adjList {
		builder.WriteString("\t\t" + strconv.Itoa(k.Val) + ": ")
		for _, vet := range v {
			builder.WriteString(strconv.Itoa(vet.Val) + " ")
		}
		fmt.Println(builder.String())
		builder.Reset()
	}
}

// DeleteSliceElms 删除切片指定元素
func DeleteSliceElms[T any](a []T, elms ...T) []T {
	if len(a) == 0 || len(elms) == 0 {
		return a
	}
	// 先将元素转为 set
	m := make(map[any]struct{})
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 过滤掉指定元素
	res := make([]T, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}
