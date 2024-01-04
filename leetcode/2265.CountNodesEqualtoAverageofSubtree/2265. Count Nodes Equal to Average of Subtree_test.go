package leetcode

import "testing"

func Test_averageOfSubtree(t *testing.T) {
	// Create a sample tree
	root := &TreeNode{
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

	want := 3

	got := averageOfSubtree(root)

	if got != want {
		t.Errorf("averageOfSubtree() = %v, want %v", got, want)
	}
}
