package leetcode

import "container/heap"

func (h *KthLargest) Len() int {
	return len(h.data)
}

func (h *KthLargest) Less(i int, j int) bool {
	return h.data[i] < h.data[j]
}

// Swap swaps the elements with indexes i and j.
func (h *KthLargest) Swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *KthLargest) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *KthLargest) Pop() any {
	item := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return item
}

type KthLargest struct {
	data []int
	top  int
}

func Constructor(k int, nums []int) KthLargest {
	h := KthLargest{
		top: k,
	}
	for _, num := range nums {
		h.Add(num)
	}
	return h
}

func (k *KthLargest) Add(val int) int {
	heap.Push(k, val)
	if k.Len() > k.top {
		heap.Pop(k)
	}
	return k.data[0]
}

func TopK(k int, nums []int) int {
	// return quickSortTopK(nums, 0, len(nums)-1, k)
	h := Constructor(k, nums)
	return h.Add(0)
}

func QuickTopK(k int, nums []int) int {
	return quickSortTopK(nums, 0, len(nums)-1, k)
}

func quickSortTopK(arr []int, left, right, k int) int {
	if k < 0 {
		return 0
	}
	if left > right {
		return 0
	}
	pivot := arr[left]
	i, j := left+1, right
	for i < j {
		for i <= right && arr[i] > pivot {
			i++
		}
		for j >= left+1 && arr[j] < pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left], arr[j] = arr[j], arr[left]
	if j-left+1 == k {
		return arr[j]
	} else if j-left+1 > k {
		return quickSortTopK(arr, left, j-1, k)
	} else {
		return quickSortTopK(arr, j+1, right, k-j-1)
	}
}
