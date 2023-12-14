package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{
			name: "case 1",
			nums: []int{-10, -3, 0, 5, 9},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -10,
					Right: &TreeNode{
						Val: -3,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
		},
		{
			name: "case 2",
			nums: []int{1, 3, 5, 7, 9},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
		},
		{
			name: "case 3",
			nums: []int{1},
			want: &TreeNode{
				Val: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
