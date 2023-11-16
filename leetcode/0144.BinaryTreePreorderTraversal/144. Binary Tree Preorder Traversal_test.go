package leetcode

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	want := []int{1, 2, 3}
	if got := preorderTraversal(root); !reflect.DeepEqual(got, want) {
		t.Errorf("preorderTraversal() = %v, want %v", got, want)
	}
}
