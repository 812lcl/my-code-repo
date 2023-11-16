package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		var tmp []int
		len := l.Len()
		for i := 0; i < len; i++ {
			node := l.Front().Value.(*TreeNode)
			tmp = append(tmp, node.Val)
			l.Remove(l.Front())
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
