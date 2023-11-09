package structure

import "container/heap"

type Item struct {
	Value    int
	Priority int
	Index    int
}

type PriorityQueue []*Item

// Len is the number of elements in the collection.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps the elements with indexes i and j.
func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)

}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func PriorityQueueFunc() {
	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &Item{Value: 1, Priority: 1})
	heap.Push(&pq, &Item{Value: 2, Priority: 2})
	heap.Pop(&pq)
}
