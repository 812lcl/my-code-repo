package leetcode

import "812lcl/my-code-repo/structure"

func maxDepth(root *structure.TreeNode) int {
	if root == nil {
		return 0
	}
	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)
	if ld > rd {
		return ld + 1
	} else {
		return rd + 1
	}
}
