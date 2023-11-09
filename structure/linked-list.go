package structure

import "fmt"

type ILinkedList interface {
	New()
	Get(val int) *Node
	Append(n *Node)
	Remove(n *Node)
}

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type DoubleLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *DoubleLinkedList) New() {
	l.Head = nil
	l.Tail = nil
	l.Length = 0
}

func (l *DoubleLinkedList) Get(val int) *Node {
	for currentNode := l.Head; currentNode != nil; currentNode = currentNode.Next {
		if currentNode.Data == val {
			return currentNode
		}
	}
	return nil
}

func (l *DoubleLinkedList) Append(n *Node) {
	l.Length++
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	l.Tail.Next = n
	n.Prev = l.Tail
	l.Tail = n
}

func (l *DoubleLinkedList) Remove(n *Node) {
	if l.Head == nil {
		return
	}

	l.Length--
	if l.Head == n {
		l.Head = l.Head.Next
		l.Head.Prev = nil
		return
	}

	if l.Tail == n {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
		return
	}

	// 如果 n 不存在与链表中呢？
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

type CircularLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *CircularLinkedList) New() {
	l.Head = nil
	l.Tail = nil
	l.Length = 0
}

func (l *CircularLinkedList) Get(val int) *Node {
	// 引入哨兵节点是否有利于简化逻辑？
	var currentNode *Node
	for currentNode = l.Head; currentNode != nil && currentNode.Next != l.Head; currentNode = currentNode.Next {
		if currentNode.Data == val {
			return currentNode
		}
	}
	if currentNode == nil {
		return nil
	} else if currentNode.Next == l.Head {
		if currentNode.Data == val {
			return currentNode
		} else {
			return nil
		}
	}
	return nil
}

func (l *CircularLinkedList) Append(n *Node) {
	l.Length++
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		n.Next = n
		return
	}

	l.Tail.Next = n
	n.Next = l.Head
	l.Tail = n
}

func (l *CircularLinkedList) Remove(n *Node) {
	if l.Head == nil {
		return
	}
	currentNode := l.Head
	preNode := l.Tail
	for currentNode.Next != l.Head {
		if currentNode == n {
			break
		}
		preNode = currentNode
		currentNode = currentNode.Next
	}
	if currentNode != n {
		return // n is not in the list
	}
	l.Length--
	if l.Length == 0 {
		l.Head = nil
		l.Tail = nil
	} else if n == l.Head {
		l.Head = n.Next
		l.Tail.Next = l.Head
	} else if n == l.Tail {
		preNode.Next = l.Head
		l.Tail = preNode
	} else {
		preNode.Next = n.Next
	}
	n.Next = nil
}

type CircularDoubleLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (cdl *CircularDoubleLinkedList) New() {
	cdl.Head = nil
	cdl.Tail = nil
	cdl.Length = 0
}

func (cdl *CircularDoubleLinkedList) Get(val int) *Node {
	// 引入哨兵节点是否有利于简化逻辑？
	var currentNode *Node
	for currentNode = cdl.Head; currentNode != nil && currentNode.Next != cdl.Head; currentNode = currentNode.Next {
		if currentNode.Data == val {
			return currentNode
		}
	}
	if currentNode == nil {
		return nil
	} else if currentNode.Next == cdl.Head {
		if currentNode.Data == val {
			return currentNode
		} else {
			return nil
		}
	}
	return nil
}

func (cdl *CircularDoubleLinkedList) Append(n *Node) {
	cdl.Length++
	if cdl.Head == nil {
		cdl.Head = n
		cdl.Tail = n
		n.Next = n
		n.Prev = n
		return
	}

	cdl.Tail.Next = n
	n.Next = cdl.Head
	cdl.Head.Prev = n
	n.Prev = cdl.Tail
	cdl.Tail = n

}

func (cdl *CircularDoubleLinkedList) Remove(n *Node) {
	if cdl.Head == nil {
		return
	}
	currentNode := cdl.Head
	for currentNode.Next != cdl.Head {
		if currentNode == n {
			break
		}
		currentNode = currentNode.Next
	}
	if currentNode != n {
		return // n is not in the list
	}
	cdl.Length--
	if cdl.Length == 0 {
		cdl.Head = nil
		cdl.Tail = nil
	} else if n == cdl.Head {
		cdl.Head = n.Next
		cdl.Head.Prev = cdl.Tail
		cdl.Tail.Next = cdl.Head
	} else if n == cdl.Tail {
		n.Prev.Next = cdl.Head
		cdl.Tail = n.Prev
		cdl.Head.Prev = n.Prev
	} else {
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	}
	n.Next = nil

}

func LinkedListFunc() {
	node1 := &Node{Data: 2, Next: nil}
	node2 := &Node{Data: 3, Next: nil}
	headNode := &Node{Data: 1, Next: node1}
	fmt.Printf("address of headNode is %p\n", headNode)
	fmt.Printf("address of node1 is %p\n", node1)
	fmt.Printf("address of node2 is %p\n", node2)

	list := &DoubleLinkedList{}
	list.New()
	list.Append(headNode)
	list.Append(node1)
	list.Append(node2)
	fmt.Printf("Value of list is %+v\n", list)
	fmt.Printf("Address of head is %p, Value of head is %+v\n", list.Head, list.Head)
	fmt.Printf("Address of node1 is %p, Value of node1 is %+v\n", node1, node1)
	fmt.Printf("Address of tail is %p, Value of tail is %+v\n", list.Tail, list.Tail)
	fmt.Printf("Get Value =2 in list is %+v\n", list.Get(2))
	list.Remove(headNode)
	fmt.Printf("After remove head, Value of list is %+v\n", list)
	list.Remove(node2)
	fmt.Printf("After remove tail, Value of list is %+v\n", list)
	node1.Next = nil
	node1.Prev = nil
	node2.Next = nil
	node2.Prev = nil
	headNode.Next = nil
	headNode.Prev = nil

	fmt.Println()
	clist := &CircularLinkedList{}
	clist.New()
	clist.Append(headNode)
	clist.Append(node1)
	clist.Append(node2)
	fmt.Printf("Value of list is %+v\n", clist)
	fmt.Printf("Address of head is %p, Value of head is %+v\n", clist.Head, clist.Head)
	fmt.Printf("Address of node1 is %p, Value of node1 is %+v\n", node1, node1)
	fmt.Printf("Address of tail is %p, Value of tail is %+v\n", clist.Tail, clist.Tail)
	fmt.Printf("Get Value =2 in clist is %+v\n", clist.Get(2))
	clist.Remove(headNode)
	fmt.Printf("After remove head, Value of clist is %+v\n", clist)
	clist.Remove(node2)
	fmt.Printf("After remove tail, Value of clist is %+v\n", clist)
	node1.Next = nil
	node1.Prev = nil
	node2.Next = nil
	node2.Prev = nil
	headNode.Next = nil
	headNode.Prev = nil

	fmt.Println()
	cdlist := &CircularDoubleLinkedList{}
	cdlist.New()
	cdlist.Append(headNode)
	cdlist.Append(node1)
	cdlist.Append(node2)
	fmt.Printf("Value of list is %+v\n", cdlist)
	fmt.Printf("Address of head is %p, Value of head is %+v\n", cdlist.Head, cdlist.Head)
	fmt.Printf("Address of node1 is %p, Value of node1 is %+v\n", node1, node1)
	fmt.Printf("Address of tail is %p, Value of tail is %+v\n", cdlist.Tail, cdlist.Tail)
	fmt.Printf("Get Value =2 in cdlist is %+v\n", cdlist.Get(2))
	cdlist.Remove(headNode)
	fmt.Printf("After remove head, Value of cdlist is %+v\n", cdlist)
	cdlist.Remove(node2)
	fmt.Printf("After remove tail, Value of cdlist is %+v\n", cdlist)
}
