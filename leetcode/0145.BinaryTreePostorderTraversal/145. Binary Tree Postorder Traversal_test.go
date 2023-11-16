package leetcode

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	want := []int{3, 2, 1}
	if got := postorderTraversal(root); !reflect.DeepEqual(got, want) {
		t.Errorf("postorderTraversal() = %v, want %v", got, want)
	}
}
