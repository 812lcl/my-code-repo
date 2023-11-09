package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	for i, j := head, head.Next; i != nil; {
		for j != nil && j.Val == i.Val {
			j = j.Next
		}
		i.Next = j
		i = j
		if j != nil {
			j = j.Next
		}
	}

	return head
}
