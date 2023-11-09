package leetcode

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
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
				x:    5,
			},
			want: nil,
		},
		{
			name: "single node list less than x",
			args: args{
				head: &ListNode{Val: 3},
				x:    5,
			},
			want: &ListNode{Val: 3},
		},
		{
			name: "single node list greater than x",
			args: args{
				head: &ListNode{Val: 7},
				x:    5,
			},
			want: &ListNode{Val: 7},
		},
		{
			name: "multiple node list with some nodes less than x",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}},
				x:    3,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}},
		},
		{
			name: "multiple node list with all nodes less than x",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
				x:    6,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		},
		{
			name: "multiple node list with all nodes greater than x",
			args: args{
				head: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10, Next: &ListNode{Val: 11}}}}},
				x:    6,
			},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10, Next: &ListNode{Val: 11}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %+v, want %+v", spew.Sdump(got), spew.Sdump(tt.want))
			}
		})
	}
}
