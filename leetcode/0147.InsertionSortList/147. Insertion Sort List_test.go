package leetcode

import (
	"reflect"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "empty list",
			head: nil,
			want: nil,
		},
		{
			name: "single node",
			head: &ListNode{Val: 1, Next: nil},
			want: &ListNode{Val: 1, Next: nil},
		},
		{
			name: "sorted list",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}},
		},
		{
			name: "reverse sorted list",
			head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}},
		},
		{
			name: "unsorted list",
			head: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertionSortList(tt.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
