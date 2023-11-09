package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}

	r, rnext := head, head.Next
	for i := 0; i < right-1; i++ {
		r = r.Next
		rnext = rnext.Next
	}
	r.Next = nil
	if left == 1 {
		subList := reverseList(head)
		head.Next = rnext
		return subList
	}

	lpre, l := head, head.Next
	for i := 0; i < left-2; i++ {
		l = l.Next
		lpre = lpre.Next
	}

	subList := reverseList(l)
	if lpre != nil {
		lpre.Next = subList
	}
	l.Next = rnext
	if left != 1 {
		return head
	}
	return subList
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	i, j := head, head.Next
	for j != nil {
		t := j.Next
		j.Next = i
		i = j
		j = t
	}
	head.Next = nil
	return i
}

func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	// If the list is empty, has only one node, or m >= n, return the head
	if head == nil || head.Next == nil || m >= n {
		return head
	}

	// Create a new node before the head to handle the case where m = 1
	newHead := &ListNode{Val: 0, Next: head}

	// Find the node before the mth node
	preM := newHead
	for i := 0; i < m-1; i++ {
		preM = preM.Next
	}

	// Reverse the nodes between m and n
	cur := preM.Next
	for i := 0; i < n-m; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = preM.Next
		preM.Next = next
	}

	// Return the head of the reversed list
	return newHead.Next
}
