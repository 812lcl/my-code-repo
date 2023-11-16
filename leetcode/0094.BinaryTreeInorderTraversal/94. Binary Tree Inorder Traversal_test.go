package leetcode

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	want := []int{1, 3, 2}
	if got := inorderTraversal(root); !reflect.DeepEqual(got, want) {
		t.Errorf("inorderTraversal() = %v, want %v", got, want)
	}
}
