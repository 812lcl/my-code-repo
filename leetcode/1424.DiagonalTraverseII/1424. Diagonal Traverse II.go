package leetcode

import (
	"container/list"
	"math"
)

type Node struct {
	Val int
	X   int
	Y   int
}

func findDiagonalOrder(nums [][]int) []int {
	if nums == nil {
		return nil
	}
	queue := list.New()
	var res []int
	queue.PushBack(
		&Node{
			Val: nums[0][0],
			X:   0,
			Y:   0,
		})
	nums[0][0] = math.MaxInt
	for queue.Len() != 0 {
		node := queue.Front().Value.(*Node)
		res = append(res, node.Val)
		children := getChildren(nums, node.X, node.Y)
		if children != nil || children.Len() != 0 {
			queue.PushBackList(children)
		}
		queue.Remove(queue.Front())
	}
	return res
}

func getChildren(nums [][]int, x, y int) *list.List {
	if nums == nil {
		return nil
	}
	if x > len(nums)-1 {
		return nil
	}
	if y > len(nums[x])-1 {
		return nil
	}
	x1, y1, x2, y2 := x+1, y, x, y+1
	res := list.New()
	if x1 < len(nums) && y1 < len(nums[x1]) && nums[x1][y1] != math.MaxInt {
		res.PushBack(
			&Node{
				Val: nums[x1][y1],
				X:   x1,
				Y:   y1,
			})
		nums[x1][y1] = math.MaxInt
	}
	if x2 < len(nums) && y2 < len(nums[x2]) && nums[x2][y2] != math.MaxInt {
		res.PushBack(
			&Node{
				Val: nums[x2][y2],
				X:   x2,
				Y:   y2,
			})
		nums[x2][y2] = math.MaxInt
	}
	return res
}
