package leetcode

import "testing"

func Test_minDepth(t *testing.T) {
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

	want := 2
	got := minDepth(root)
	if got != want {
		t.Errorf("minDepth() = %v, want %v", got, want)
	}
	got = minDepth1(root)
	if got != want {
		t.Errorf("minDepth() = %v, want %v", got, want)
	}
}
