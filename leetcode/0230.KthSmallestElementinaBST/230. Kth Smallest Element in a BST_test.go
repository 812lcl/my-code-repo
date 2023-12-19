package leetcode

import (
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	// Create a binary search tree for testing
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	tests := []struct {
		name string
		root *TreeNode
		k    int
		want int
	}{
		{
			name: "case 1",
			root: root,
			k:    1,
			want: 1,
		},
		{
			name: "case 2",
			root: root,
			k:    2,
			want: 2,
		},
		{
			name: "case 3",
			root: root,
			k:    3,
			want: 3,
		},
		{
			name: "case 4",
			root: root,
			k:    4,
			want: 4,
		},
		{
			name: "case 5",
			root: root,
			k:    5,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthSmallest(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
