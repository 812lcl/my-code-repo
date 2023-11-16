package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	arr := inOrder(root)
	var node1, node2 *TreeNode
	var second bool
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].Val > arr[i+1].Val && !second {
			node1 = arr[i]
			node2 = arr[i+1]
			second = true
			continue
		}
		if arr[i].Val > arr[i+1].Val && second {
			node2 = arr[i+1]
			break
		}
	}
	node1.Val, node2.Val = node2.Val, node1.Val

}

func inOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	var res []*TreeNode
	res = append(res, inOrder(root.Left)...)
	res = append(res, root)
	res = append(res, inOrder(root.Right)...)
	return res
}

func recoverTree1(root *TreeNode) {
	var prev, first, second *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil {
			if first == nil && prev.Val >= node.Val {
				first = prev
			}
			if first != nil && prev.Val >= node.Val {
				second = node
			}
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)

	first.Val, second.Val = second.Val, first.Val
}
