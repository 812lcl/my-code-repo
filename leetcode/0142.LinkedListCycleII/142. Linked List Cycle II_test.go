package leetcode

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	// Create a linked list with a cycle
	head := &ListNode{Val: 3}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: -4}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1

	// Test the function
	want := node1
	if got := detectCycle(head); got != want {
		t.Errorf("detectCycle() = %v, want %v", got, want)
	}

	// Create a linked list without a cycle
	head = &ListNode{Val: 1}
	node1 = &ListNode{Val: 2}
	head.Next = node1

	// Test the function
	want = nil
	if got := detectCycle(head); got != want {
		t.Errorf("detectCycle() = %v, want %v", got, want)
	}
	if got := detectCycle1(head); got != want {
		t.Errorf("detectCycle() = %v, want %v", got, want)
	}
}
