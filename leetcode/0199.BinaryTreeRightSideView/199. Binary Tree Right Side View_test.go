package leetcode

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	// Test case 1: Single node tree
	root1 := &TreeNode{Val: 1}
	want1 := []int{1}
	if got1 := rightSideView(root1); !reflect.DeepEqual(got1, want1) {
		t.Errorf("rightSideView() = %v, want %v", got1, want1)
	}

	// Test case 2: Complete binary tree
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Left = &TreeNode{Val: 4}
	root2.Left.Right = &TreeNode{Val: 5}
	root2.Right.Left = &TreeNode{Val: 6}
	root2.Right.Right = &TreeNode{Val: 7}
	want2 := []int{1, 3, 7}
	if got2 := rightSideView(root2); !reflect.DeepEqual(got2, want2) {
		t.Errorf("rightSideView() = %v, want %v", got2, want2)
	}

	// Test case 3: Skewed binary tree
	root3 := &TreeNode{Val: 1}
	root3.Right = &TreeNode{Val: 2}
	root3.Right.Right = &TreeNode{Val: 3}
	root3.Right.Right.Right = &TreeNode{Val: 4}
	want3 := []int{1, 2, 3, 4}
	if got3 := rightSideView(root3); !reflect.DeepEqual(got3, want3) {
		t.Errorf("rightSideView() = %v, want %v", got3, want3)
	}
}
