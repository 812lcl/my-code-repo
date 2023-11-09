package leetcode

type MyQueue struct {
	s1 *[]int
	s2 *[]int
}

func Constructor() MyQueue {
	return MyQueue{s1: &[]int{}, s2: &[]int{}}
}

func (q *MyQueue) Push(x int) {
	*q.s1 = append(*q.s1, x)
}

func (q *MyQueue) Pop() int {
	s2 := q.s2

	item := q.Peek()

	*s2 = (*s2)[:len(*s2)-1]

	return item
}

func (q *MyQueue) Peek() int {
	s1 := q.s1
	s2 := q.s2

	if len(*s2) == 0 {
		for len(*s1) > 0 {
			item := (*s1)[len(*s1)-1]
			*s1 = (*s1)[:len(*s1)-1]

			*s2 = append(*s2, item)
		}
	}

	return (*s2)[len(*s2)-1]
}

func (q *MyQueue) Empty() bool {
	return len(*q.s1)+len(*q.s2) == 0
}
