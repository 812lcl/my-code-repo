package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	q := list.New()
	q.PushBack(root)
	for q.Len() != 0 {
		l := q.Len()
		for i := 0; i < l; i++ {
			node := q.Front().Value.(*TreeNode)
			if i == l-1 {
				res = append(res, node.Val)
			}
			q.Remove(q.Front())
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
	}
	return res
}
