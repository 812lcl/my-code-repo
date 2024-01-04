package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	num, _, _ := sumOfSubtree(root)
	return num
}

func sumOfSubtree(root *TreeNode) (num, sum, count int) {
	if root == nil {
		return 0, 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return 1, root.Val, 1
	}
	numl, suml, countl := sumOfSubtree(root.Left)
	numr, sumr, countr := sumOfSubtree(root.Right)
	if root.Val == (suml+sumr+root.Val)/(countl+countr+1) {
		return numl + numr + 1, suml + sumr + root.Val, countl + countr + 1
	} else {
		return numl + numr, suml + sumr + root.Val, countl + countr + 1
	}
}
