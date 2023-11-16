package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	length := 0
	for i, num := range inorder {
		if num == root.Val {
			length = i + 1
			break
		}
	}
	root.Left = buildTree(preorder[1:length], inorder[0:length-1])
	root.Right = buildTree(preorder[length:], inorder[length:])
	return root
}
