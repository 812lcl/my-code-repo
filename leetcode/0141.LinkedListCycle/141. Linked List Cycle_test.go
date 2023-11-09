package leetcode

import (
	"testing"
)

func Test_hasCycle(t *testing.T) {
	// Test case 1: cycle exists
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = head1.Next
	if !hasCycle(head1) {
		t.Errorf("Test case 1 failed: expected true, but got false")
	}

	// Test case 2: cycle does not exist
	head2 := &ListNode{Val: 1}
	head2.Next = &ListNode{Val: 2}
	head2.Next.Next = &ListNode{Val: 3}
	if hasCycle(head2) {
		t.Errorf("Test case 2 failed: expected false, but got true")
	}

	// Test case 3: empty list
	head3 := (*ListNode)(nil)
	if hasCycle(head3) {
		t.Errorf("Test case 3 failed: expected false, but got true")
	}

	// Test case 4: single node list
	head4 := &ListNode{Val: 1}
	if hasCycle(head4) {
		t.Errorf("Test case 4 failed: expected false, but got true")
	}
}
