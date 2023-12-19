package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	return inOrder(root, &k)
}

func inOrder(root *TreeNode, k *int) int {
	if root == nil {
		return 0
	}
	res := inOrder(root.Left, k)
	if res != 0 {
		return res
	}
	*k--
	if *k == 0 {
		return root.Val
	}
	return inOrder(root.Right, k)
}
