package leetcode

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.Front().Value.(*Node)
			queue.Remove(queue.Front())
			if i != l-1 {
				node.Next = queue.Front().Value.(*Node)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}
