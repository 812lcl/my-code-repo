package leetcode

import (
	"container/list"
	"math"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	count := make(map[int][]int)
	min := math.MaxInt
	m := make([][]int, n)
	for _, v := range edges {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}
	for i := 0; i < n; i++ {
		h := bfs(i, m)
		if h <= min {
			min = h
			count[min] = append(count[min], i)
		}
	}
	return count[min]
}

func bfs(n int, m [][]int) int {
	visited := make(map[int]bool)
	q := list.New()
	q.PushBack(n)
	var h int
	for q.Len() != 0 {
		l := q.Len()
		h++
		for i := 0; i < l; i++ {
			num := q.Front().Value.(int)
			for _, v := range m[num] {
				if visited[v] {
					continue
				}
				q.PushBack(v)
			}
			visited[num] = true
			q.Remove(q.Front())
		}
	}
	return h
}
