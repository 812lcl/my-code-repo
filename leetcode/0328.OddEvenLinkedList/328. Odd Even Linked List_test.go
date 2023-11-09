package leetcode

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "empty list",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "one element",
			args: args{
				head: &ListNode{Val: 1},
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "two elements",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name: "three elements",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}},
		},
		{
			name: "four elements",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}},
		},
		{
			name: "five elements",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
