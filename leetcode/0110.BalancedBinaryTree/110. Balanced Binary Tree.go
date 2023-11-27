package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := heightOfTree(root.Left) - heightOfTree(root.Right)
	return diff >= -1 && diff <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func heightOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return max(heightOfTree(root.Left), heightOfTree(root.Right)) + 1
}
