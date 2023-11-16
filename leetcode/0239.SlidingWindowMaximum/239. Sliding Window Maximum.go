package leetcode

import (
	"container/list"
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	for i := 0; i < len(nums)-k+1; i++ {
		max := math.MinInt
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res
}

func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int, 0, k) // store the index of nums
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums { // if the left-most index is out of window, remove it
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < v { // maintain window
			window = window[0 : len(window)-1]
		}
		window = append(window, i) // store the index of nums
		if i >= k-1 {
			result = append(result, nums[window[0]]) // the left-most is the index of max value in nums
		}
	}
	return result
}

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var result []int
	queue := list.New()

	for i, num := range nums {
		// 如果队列头部的索引不在当前窗口范围内，移除该索引
		if queue.Len() > 0 && queue.Front().Value.(int) <= i-k {
			queue.Remove(queue.Front())
		}

		// 移除所有小于当前元素的值，确保队列中的元素递减
		for queue.Len() > 0 && nums[queue.Back().Value.(int)] < num {
			queue.Remove(queue.Back())
		}

		// 将当前元素索引添加到队列中
		queue.PushBack(i)

		// 将当前窗口最大值加入结果集
		if i >= k-1 {
			result = append(result, nums[queue.Front().Value.(int)])
		}
	}

	return result
}
