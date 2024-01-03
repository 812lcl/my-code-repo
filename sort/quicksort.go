package sort

import "fmt"

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	for left < right {
		pivot := partition(arr, left, right)
		if pivot-left < right-pivot {
			quickSort(arr, left, pivot-1)
			left = pivot + 1
		} else {
			quickSort(arr, pivot+1, right)
			right = pivot - 1
		}
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[left]
	i, j := left, right
	for i < j {
		for j > i && arr[j] >= pivot {
			j--
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[left], arr[j] = arr[j], arr[left]
	return j
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
