package leetcode

type MyCircularDeque struct {
	capacity int
	length   int
	head     *Node
	tail     *Node
}

type Node struct {
	val  int
	next *Node
	prev *Node
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		capacity: k,
		length:   0,
		head:     nil,
		tail:     nil,
	}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}
	q.length++
	node := &Node{
		val:  value,
		next: q.head,
		prev: q.tail,
	}
	if q.length == 1 {
		node.next = node
		node.prev = node
		q.head = node
		q.tail = node
		return true
	}
	q.head.prev = node
	q.head = node
	q.tail.next = node
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}
	q.length++
	node := &Node{
		val:  value,
		next: q.head,
		prev: q.tail,
	}
	if q.length == 1 {
		node.next = node
		node.prev = node
		q.head = node
		q.tail = node
		return true
	}
	q.tail.next = node
	q.tail = node
	q.head.prev = node
	return true

}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	q.length--
	if q.length == 0 {
		q.head = nil
		q.tail = nil
		return true
	}
	q.tail.next = q.head.next
	q.head.next.prev = q.tail
	q.head = q.head.next
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	q.length--
	if q.length == 0 {
		q.head = nil
		q.tail = nil
		return true
	}
	q.tail.prev.next = q.head
	q.head.prev = q.tail.prev
	q.tail = q.tail.prev
	return true
}

func (q *MyCircularDeque) GetFront() int {
	if q.head == nil {
		return -1
	}
	return q.head.val
}

func (q *MyCircularDeque) GetRear() int {
	if q.tail == nil {
		return -1
	}
	return q.tail.val
}

func (q *MyCircularDeque) IsEmpty() bool {
	return q.length == 0
}

func (q *MyCircularDeque) IsFull() bool {
	return q.length == q.capacity
}
