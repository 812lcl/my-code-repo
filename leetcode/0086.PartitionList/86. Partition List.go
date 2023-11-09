package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	j, k := dummy1, dummy2
	for i := head; i != nil; {
		if i.Val < x {
			j.Next = i
			j = j.Next
		} else {
			k.Next = i
			k = k.Next
		}
		i = i.Next
	}
	j.Next = dummy2.Next
	k.Next = nil
	return dummy1.Next
}
