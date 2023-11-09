package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	uniq := make(map[*ListNode]bool)
	for i := head; i != nil; i = i.Next {
		if _, ok := uniq[i]; ok {
			return i
		}
		uniq[i] = true
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	cycleLen := 0
	isCycle := false
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
		cycleLen++
		if slow == fast {
			isCycle = true
			break
		}
	}
	if !isCycle {
		return nil
	}
	slow := head
	fast := head
	for i := 0; i < cycleLen; i++ {
		fast = fast.Next
	}
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
