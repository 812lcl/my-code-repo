package sort

import "fmt"

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivot := arr[left]
	i, j := left+1, right
	for i < j {
		for i <= right && arr[i] <= pivot {
			i++
		}
		for j >= left+1 && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left], arr[j] = arr[j], arr[left]

	quickSort(arr, left, j-1)
	quickSort(arr, j+1, right)
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
