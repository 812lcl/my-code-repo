package leetcode

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	tests := []struct {
		name      string
		inorder   []int
		postorder []int
		want      *TreeNode
	}{
		{
			name:      "example 1",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want: &TreeNode{
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
			},
		},
		{
			name:      "example 2",
			inorder:   []int{2, 1},
			postorder: []int{2, 1},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			name:      "empty tree",
			inorder:   []int{},
			postorder: []int{},
			want:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.inorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
