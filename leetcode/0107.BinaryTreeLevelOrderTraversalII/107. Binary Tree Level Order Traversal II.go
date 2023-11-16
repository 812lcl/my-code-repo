package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := levelOrder(root)
	for i, j := 0, len(res)-1; i <= j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
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
