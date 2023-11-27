package leetcode

import (
	"reflect"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		name  string
		input *TreeNode
		want  bool
	}{
		{
			name:  "Case 1",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}},
			want:  true,
		},
		{
			name:  "Case 2",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want:  false,
		},
		{
			name:  "Case 3",
			input: nil,
			want:  true,
		},
		{
			name:  "Case 4",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
			want:  true,
		},
		{
			name:  "Case 5",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSymmetric1(t *testing.T) {
	tests := []struct {
		name  string
		input *TreeNode
		want  bool
	}{
		{
			name:  "Case 1",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}},
			want:  true,
		},
		{
			name:  "Case 2",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want:  false,
		},
		{
			name:  "Case 3",
			input: nil,
			want:  true,
		},
		{
			name:  "Case 4",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
			want:  true,
		},
		{
			name:  "Case 5",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
