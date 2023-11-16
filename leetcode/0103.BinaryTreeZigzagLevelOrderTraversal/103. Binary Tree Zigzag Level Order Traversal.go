package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := levelOrder(root)
	for i := 0; i < len(res); i++ {
		if i%2 == 0 {
			continue
		}
		for j, k := 0, len(res[i])-1; j <= k; {
			res[i][j], res[i][k] = res[i][k], res[i][j]
			j++
			k--
		}
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
