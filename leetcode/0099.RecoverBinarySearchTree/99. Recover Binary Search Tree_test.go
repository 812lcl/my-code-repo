package leetcode

import (
	"reflect"
	"testing"
)

func Test_recoverTree(t *testing.T) {
	// Test cases
	tests := []struct {
		name  string
		input *TreeNode
		want  *TreeNode
	}{
		{
			name: "Example 1",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		},
		{
			name: "Example 2",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("recoverTree() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func Test_recoverTree1(t *testing.T) {
	// Test cases
	tests := []struct {
		name  string
		input *TreeNode
		want  *TreeNode
	}{
		{
			name: "Example 1",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		},
		{
			name: "Example 2",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree1(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("recoverTree() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}
