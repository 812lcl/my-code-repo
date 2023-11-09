package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	lists := []*ListNode{list1, list2, list3}
	want := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}}
	if got := mergeKLists(lists); !reflect.DeepEqual(got, want) {
		t.Errorf("mergeKLists() = %v, want %v", got, want)
	}
}
