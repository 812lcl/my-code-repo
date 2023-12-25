package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	var (
		i    = 1
		temp = head
		prev *ListNode
	)

	for temp != nil {
		decr := i
		var (
			tempOrig = temp
			prevOrig = prev
			tempNext *ListNode
			length   = 0
		)

		// Check length
		for decr != 0 && temp != nil {
			prev = temp
			temp = temp.Next
			tempNext = temp
			decr--
			length++
		}

		// If in even group then reverse
		if length%2 == 0 {
			temp, prev = tempOrig, prevOrig
			prev.Next = reverse(temp, length)
			tempOrig.Next = tempNext
			prev, temp = tempOrig, tempNext
		}
		i++
	}

	return head
}

func reverse(head *ListNode, k int) *ListNode {
	var (
		prev *ListNode
		temp = head
	)
	for temp != nil && k > 0 {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
		k--
	}
	return prev
}
