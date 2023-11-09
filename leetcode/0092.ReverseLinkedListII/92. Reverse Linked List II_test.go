package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case 1",
			args: args{
				head:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
				left:  2,
				right: 4,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}}},
		},
		{
			name: "case 2",
			args: args{
				head:  &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
				left:  2,
				right: 4,
			},
			want: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 1}}}}},
		},
		{
			name: "case 3",
			args: args{
				head:  &ListNode{Val: 1},
				left:  1,
				right: 1,
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "case 4",
			args: args{
				head:  &ListNode{Val: 3, Next: &ListNode{Val: 5}},
				left:  1,
				right: 2,
			},
			want: &ListNode{Val: 5, Next: &ListNode{Val: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBetween1(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case 1",
			args: args{
				head:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
				left:  2,
				right: 4,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}}},
		},
		{
			name: "case 2",
			args: args{
				head:  &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
				left:  2,
				right: 4,
			},
			want: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 1}}}}},
		},
		{
			name: "case 3",
			args: args{
				head:  &ListNode{Val: 1},
				left:  1,
				right: 1,
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "case 4",
			args: args{
				head:  &ListNode{Val: 3, Next: &ListNode{Val: 5}},
				left:  1,
				right: 2,
			},
			want: &ListNode{Val: 5, Next: &ListNode{Val: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween1(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
