package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
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
	want := [][]int{{15, 7}, {9, 20}, {3}}
	got := levelOrderBottom(root)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrderBottom() = %v, want %v", got, want)
	}
}
