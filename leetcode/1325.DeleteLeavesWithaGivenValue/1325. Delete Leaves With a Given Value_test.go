package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeLeafNodes(t *testing.T) {
	// Test case 1: Remove leaf nodes with value 3
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	target1 := 3
	want1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	if got1 := removeLeafNodes(root1, target1); !reflect.DeepEqual(got1, want1) {
		t.Errorf("removeLeafNodes() = %v, want %v", got1, want1)
	}

	// Test case 2: Remove leaf nodes with value 2
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	target2 := 2
	want2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	if got2 := removeLeafNodes(root2, target2); !reflect.DeepEqual(got2, want2) {
		t.Errorf("removeLeafNodes() = %v, want %v", got2, want2)
	}

	// Test case 3: Remove leaf nodes with value 5
	root3 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	target3 := 5
	want3 := (*TreeNode)(nil)
	if got3 := removeLeafNodes(root3, target3); !reflect.DeepEqual(got3, want3) {
		t.Errorf("removeLeafNodes() = %v, want %v", got3, want3)
	}
}
