package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	odd, even := dummy1, dummy2
	for i, j := head, 1; i != nil; j++ {
		if j%2 == 1 {
			odd.Next = i
			odd = odd.Next
		} else {
			even.Next = i
			even = even.Next
		}
		i = i.Next
	}
	odd.Next = dummy2.Next
	even.Next = nil
	return dummy1.Next
}
