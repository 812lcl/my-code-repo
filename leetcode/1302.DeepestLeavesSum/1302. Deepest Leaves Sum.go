package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	var level int
	res := make(map[int][]int)
	dfs(root, level, res)
	max := 0
	var leaves []int
	for k, v := range res {
		if k > max {
			max = k
			leaves = v
		}
	}
	var sum int
	for _, n := range leaves {
		sum += n
	}
	return sum
}

func dfs(root *TreeNode, level int, res map[int][]int) {
	if root == nil {
		return
	}

	level++
	res[level] = append(res[level], root.Val)
	dfs(root.Left, level, res)
	dfs(root.Right, level, res)
}

func deepestLeavesSum1(root *TreeNode) int {
	maxLevel, sum := 0, 0
	dfsDeepestLeavesSum(root, 0, &maxLevel, &sum)
	return sum
}

func dfsDeepestLeavesSum(root *TreeNode, level int, maxLevel, sum *int) {
	if root == nil {
		return
	}
	if level > *maxLevel {
		*maxLevel, *sum = level, root.Val
	} else if level == *maxLevel {
		*sum += root.Val
	}
	dfsDeepestLeavesSum(root.Left, level+1, maxLevel, sum)
	dfsDeepestLeavesSum(root.Right, level+1, maxLevel, sum)
}

func deepestLeavesSum2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		var sum int
		l := queue.Len()
		for i := 0; i < l; i++ {
			p := queue.Front().Value.(*TreeNode)
			sum += p.Val
			queue.Remove(queue.Front())
			if p.Left != nil {
				queue.PushBack(p.Left)
			}
			if p.Right != nil {
				queue.PushBack(p.Right)
			}
		}
		res = sum
	}
	return res
}
