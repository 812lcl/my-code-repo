package leetcode

import (
	"testing"
)

func Test_isValidBST(t *testing.T) {
	tree1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	tree2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	tree3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}
	tree4 := &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}
	tree5 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	tests := []struct {
		name string
		tree *TreeNode
		want bool
	}{
		{
			name: "valid BST",
			tree: tree1,
			want: true,
		},
		{
			name: "invalid BST",
			tree: tree2,
			want: false,
		},
		{
			name: "invalid BST with duplicate values",
			tree: tree3,
			want: false,
		},
		{
			name: "invalid BST with duplicate values",
			tree: tree4,
			want: false,
		},
		{
			name: "invalid BST",
			tree: tree5,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.tree); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
