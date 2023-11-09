package structure

import "fmt"

type IQueue interface {
	Enqueue(item interface{})
	Dequeue() interface{}
	IsEmpty() bool
}

type Queue []interface{}

func (q *Queue) Enqueue(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

type LinkedQueue DoubleLinkedList

func (lq *LinkedQueue) Enqueue(item interface{}) {
	newNode := &Node{Data: item}
	if lq.Head == nil {
		lq.Head = newNode
		lq.Tail = newNode
	} else {
		lq.Tail.Next = newNode
		lq.Tail = newNode
	}
	lq.Length++
}

func (lq *LinkedQueue) Dequeue() interface{} {
	if lq.IsEmpty() {
		return nil
	}
	item := lq.Head.Data
	lq.Head = lq.Head.Next
	lq.Length--
	return item
}

func (lq *LinkedQueue) IsEmpty() bool {
	return lq.Length == 0
}

type CycleQueue struct {
	data [10]interface{}
	head int
	tail int
	size int
}

func (q *CycleQueue) Enqueue(item interface{}) {
	if q.size == len(q.data) {
		return
	}
	q.data[q.tail] = item
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

func (q *CycleQueue) Dequeue() interface{} {
	if q.size == 0 {
		return nil
	}
	item := q.data[q.head]
	q.head = (q.head + 1) % len(q.data)
	q.size--
	return item
}

func (q *CycleQueue) IsEmpty() bool {
	return q.size == 0
}

func QueueFunc() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Printf("Value of q is %+v\n", q)
	fmt.Println(q.Dequeue()) // 1 true
	fmt.Println(q.Dequeue()) // 2 true
	fmt.Println(q.Dequeue()) // 3 true
	fmt.Println(q.Dequeue()) // nil false

	lq := &LinkedQueue{}
	lq.Enqueue(1)
	lq.Enqueue(2)
	lq.Enqueue(3)
	fmt.Printf("Value of lq is %+v\n", lq)
	fmt.Println(lq.Dequeue()) // 1 true
	fmt.Println(lq.Dequeue()) // 2 true
	fmt.Println(lq.Dequeue()) // 3 true
	fmt.Println(lq.Dequeue()) // nil false

	cq := &CycleQueue{}
	cq.Enqueue(1)
	cq.Enqueue(2)
	cq.Enqueue(3)
	fmt.Printf("Value of cq is %+v\n", cq)
	fmt.Println(cq.Dequeue()) // 1 true
	fmt.Println(cq.Dequeue()) // 2 true
	for i := 1; i < 12; i++ {
		cq.Enqueue(i)
	}
	fmt.Printf("Value of cq is %+v\n", cq)
	fmt.Println(cq.Dequeue()) // 3 true
	fmt.Println(cq.Dequeue()) // 1 true
	cq.Enqueue(10)
	fmt.Printf("Value of cq is %+v\n", cq)
	fmt.Println(cq.Dequeue()) // 2 true
}
