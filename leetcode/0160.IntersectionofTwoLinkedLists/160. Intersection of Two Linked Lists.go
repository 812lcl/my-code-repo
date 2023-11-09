package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for i := headA; i != nil; i = i.Next {
		for j := headB; j != nil; j = j.Next {
			if i == j {
				return i
			}
		}
	}
	return nil
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	i, j := headA, headB

	for i != j {
		if i == nil {
			i = headB
		} else {
			i = i.Next
		}

		if j == nil {
			j = headA
		} else {
			j = j.Next
		}
	}
	return i
}
