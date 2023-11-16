package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
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
	want := [][]int{{3}, {9, 20}, {15, 7}}
	got := levelOrder(root)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrder() = %v, want %v", got, want)
	}
}
