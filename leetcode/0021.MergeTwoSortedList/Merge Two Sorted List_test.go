package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}
	result := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mergeTwoLists(%v, %v) = %v; want %v", list1, list2, result, expected)
	}
}
