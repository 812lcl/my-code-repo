package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	cur := head.Next
	dummy.Next.Next = nil
	for cur != nil {
		for pre, behind := dummy, dummy.Next; behind != nil; {
			if cur.Val < behind.Val {
				next := cur.Next
				cur.Next = behind
				pre.Next = cur
				cur = next
				break
			}
			pre = pre.Next
			behind = behind.Next
		}
	}
	return dummy.Next
}

func sortList1(head *ListNode) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length <= 1 {
		return head
	}

	middleNode := middleNode(head)
	cur = middleNode.Next
	middleNode.Next = nil
	middleNode = cur

	left := sortList1(head)
	right := sortList1(middleNode)
	return mergeTwoLists(left, right)
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
