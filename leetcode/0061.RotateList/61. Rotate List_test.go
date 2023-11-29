package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	// Test case 1: Rotate by 2
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	k1 := 2
	want1 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}}
	if got1 := rotateRight(head1, k1); !reflect.DeepEqual(got1, want1) {
		t.Errorf("rotateRight() = %v, want %v", got1, want1)
	}

	// Test case 2: Rotate by 4
	head2 := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}
	k2 := 4
	want2 := &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}
	if got2 := rotateRight(head2, k2); !reflect.DeepEqual(got2, want2) {
		t.Errorf("rotateRight() = %v, want %v", got2, want2)
	}

	// Test case 3: Rotate by 0
	head3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	k3 := 0
	want3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	if got3 := rotateRight(head3, k3); !reflect.DeepEqual(got3, want3) {
		t.Errorf("rotateRight() = %v, want %v", got3, want3)
	}
}
