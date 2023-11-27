package leetcode

import (
	"testing"
)

func Test_isBalanced(t *testing.T) {
	// Test case 1: Balanced tree
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	if !isBalanced(root1) {
		t.Errorf("isBalanced() = false, want true")
	}

	// Test case 2: Unbalanced tree
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	if isBalanced(root2) {
		t.Errorf("isBalanced() = true, want false")
	}
}
