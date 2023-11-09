package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	len := 0
	for p := head; p != nil; p = p.Next {
		len++
	}

	if len == 0 || len < n || (len == 1 && n == 1) {
		return nil
	}

	if len == n {
		deleteNode := head
		head = head.Next
		deleteNode.Next = nil
		return head
	}

	preNode := head
	for i := 0; i < (len - n - 1); i++ {
		preNode = preNode.Next
	}

	deleteNode := preNode.Next
	if deleteNode.Next == nil {
		preNode.Next = nil
		deleteNode = nil
		return head
	}

	preNode.Next = deleteNode.Next
	return head

}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy

	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}
