package leetcode

import (
	"testing"
)

func Test_isSameTree(t *testing.T) {
	// Test case 1: Both trees are empty
	if isSameTree(nil, nil) != true {
		t.Error("Expected true for two empty trees")
	}

	// Test case 2: One tree is empty, the other is not
	p1 := &TreeNode{Val: 1}
	if isSameTree(p1, nil) != false {
		t.Error("Expected false for one empty tree")
	}

	// Test case 3: The trees have different values
	p2 := &TreeNode{Val: 1}
	q2 := &TreeNode{Val: 2}
	if isSameTree(p2, q2) != false {
		t.Error("Expected false for trees with different values")
	}

	// Test case 4: The trees have the same values but different structures
	p3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q3 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	if isSameTree(p3, q3) != false {
		t.Error("Expected false for trees with different structures")
	}

	// Test case 5: The trees have the same values and structures
	p4 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q4 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if isSameTree(p4, q4) != true {
		t.Error("Expected true for trees with same values and structures")
	}
}
