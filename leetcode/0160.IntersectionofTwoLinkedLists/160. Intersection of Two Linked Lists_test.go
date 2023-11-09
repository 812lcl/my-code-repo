package leetcode

import (
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	// Test case 1: No intersection
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	list2 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}
	if got := getIntersectionNode(list1, list2); got != nil {
		t.Errorf("getIntersectionNode() = %v, want %v", got, nil)
	}

	// Test case 2: Intersection at the beginning
	common := &ListNode{Val: 7, Next: &ListNode{Val: 8}}
	list1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: common}}
	list2 = &ListNode{Val: 3, Next: common}
	if got := getIntersectionNode(list1, list2); got != common {
		t.Errorf("getIntersectionNode() = %v, want %v", got, common)
	}

	// Test case 3: Intersection at the end
	common = &ListNode{Val: 9, Next: &ListNode{Val: 10}}
	list1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: common}}
	list2 = &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: common}}
	if got := getIntersectionNode(list1, list2); got != common {
		t.Errorf("getIntersectionNode() = %v, want %v", got, common)
	}

	// Test case 4: Intersection in the middle
	common = &ListNode{Val: 11, Next: &ListNode{Val: 12}}
	list1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: common}}
	list2 = &ListNode{Val: 3, Next: common}
	if got := getIntersectionNode(list1, list2); got != common {
		t.Errorf("getIntersectionNode() = %v, want %v", got, common)
	}

	// Test case 5: One list is empty
	list1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	list2 = nil
	if got := getIntersectionNode(list1, list2); got != nil {
		t.Errorf("getIntersectionNode() = %v, want %v", got, nil)
	}

	// Test case 6: Both lists are empty
	list1 = nil
	list2 = nil
	if got := getIntersectionNode(list1, list2); got != nil {
		t.Errorf("getIntersectionNode() = %v, want %v", got, nil)
	}
	if got := getIntersectionNode1(list1, list2); got != nil {
		t.Errorf("getIntersectionNode() = %v, want %v", got, nil)
	}
}
