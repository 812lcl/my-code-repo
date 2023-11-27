package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MaxInt
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, depth int, res *int) {
	depth++
	if isLeave(root) {
		if depth < *res {
			*res = depth
		}
	}
	if root.Left != nil {
		dfs(root.Left, depth, res)
	}
	if root.Right != nil {
		dfs(root.Right, depth, res)
	}
}

func isLeave(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
