package leetcode

import (
	"testing"
)

func Test_deepestLeavesSum(t *testing.T) {
	// Test case 1: Single node tree
	root1 := &TreeNode{Val: 1}
	want1 := 1
	if got1 := deepestLeavesSum(root1); got1 != want1 {
		t.Errorf("deepestLeavesSum() = %v, want %v", got1, want1)
	}
	if got1 := deepestLeavesSum1(root1); got1 != want1 {
		t.Errorf("deepestLeavesSum() = %v, want %v", got1, want1)
	}
	if got1 := deepestLeavesSum2(root1); got1 != want1 {
		t.Errorf("deepestLeavesSum() = %v, want %v", got1, want1)
	}

	// Test case 2: Tree with two levels
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	want2 := 5
	if got2 := deepestLeavesSum(root2); got2 != want2 {
		t.Errorf("deepestLeavesSum() = %v, want %v", got2, want2)
	}
	if got2 := deepestLeavesSum1(root2); got2 != want2 {
		t.Errorf("deepestLeavesSum() = %v, want %v", got2, want2)
	}
	if got2 := deepestLeavesSum2(root2); got2 != want2 {
		t.Errorf("deepestLeavesSum() = %v, want %v", got2, want2)
	}

	// Test case 3: Tree with three levels
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 6}}}
	want3 := 15
	if got3 := deepestLeavesSum(root3); got3 != want3 {
		t.Errorf("deepestLeavesSum() = %v, want %v", got3, want3)
	}
	if got3 := deepestLeavesSum1(root3); got3 != want3 {
		t.Errorf("deepestLeavesSum() = %v, want %v", got3, want3)
	}
	if got3 := deepestLeavesSum2(root3); got3 != want3 {
		t.Errorf("deepestLeavesSum() = %v, want %v", got3, want3)
	}
}
