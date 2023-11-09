package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
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
			name: "single node",
			args: args{
				head: &ListNode{Val: 1},
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "two nodes",
			args: args{
				head: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name: "three nodes",
			args: args{
				head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name: "four nodes",
			args: args{
				head: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortList1(t *testing.T) {
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
			name: "single node",
			args: args{
				head: &ListNode{Val: 1},
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "two nodes",
			args: args{
				head: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name: "three nodes",
			args: args{
				head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name: "four nodes",
			args: args{
				head: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
