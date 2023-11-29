package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	lenghth := 0
	for n := head; n != nil; {
		lenghth++
		n = n.Next
	}
	k = k % lenghth
	for i := 0; i < k; i++ {
		if b == nil {
			return nil
		}
		b = b.Next
	}
	for b.Next != nil {
		b = b.Next
		a = a.Next
	}
	b.Next = head
	head = a.Next
	a.Next = nil
	return head
}
