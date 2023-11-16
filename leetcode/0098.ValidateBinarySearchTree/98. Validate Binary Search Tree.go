package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil {
		if root.Val >= leftVal(root.Right) {
			return false
		}
		return isValidBST(root.Right)
	} else if root.Right == nil {
		if root.Val <= rightVal(root.Left) {
			return false
		}
		return isValidBST(root.Left)
	} else if root.Val <= root.Left.Val || root.Val >= root.Right.Val || root.Val <= rightVal(root.Left) || root.Val >= leftVal(root.Right) {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}

func leftVal(root *TreeNode) int {
	if root.Left != nil {
		return leftVal(root.Left)
	} else {
		return root.Val
	}
}

func rightVal(root *TreeNode) int {
	if root.Right != nil {
		return rightVal(root.Right)
	} else {
		return root.Val
	}
}
