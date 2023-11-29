package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
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

func reverseList1(head *ListNode) *ListNode {
	var reverseHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = reverseHead
		reverseHead = head
		head = next
	}
	return reverseHead
}
