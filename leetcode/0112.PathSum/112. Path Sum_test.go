package leetcode

import "testing"

func Test_hasPathSum(t *testing.T) {
	tree1 := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{Val: 4,
				Right: &TreeNode{Val: 1},
			},
		},
	}
	tree2 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	tree3 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
	}
	tests := []struct {
		name       string
		root       *TreeNode
		targetSum  int
		wantResult bool
	}{
		{
			name:       "test 1",
			root:       tree1,
			targetSum:  22,
			wantResult: true,
		},
		{
			name:       "test 2",
			root:       tree2,
			targetSum:  5,
			wantResult: false,
		},
		{
			name:       "test 3",
			root:       tree3,
			targetSum:  1,
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.root, tt.targetSum); got != tt.wantResult {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}
