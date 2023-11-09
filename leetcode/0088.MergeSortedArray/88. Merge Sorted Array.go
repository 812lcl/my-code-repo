package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := 0, 0, 0
	tmp := make([]int, m+n)
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			tmp[k] = nums1[i]
			i++
		} else {
			tmp[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		tmp[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		tmp[k] = nums2[j]
		j++
		k++
	}
	for k := 0; k < m+n; k++ {
		nums1[k] = tmp[k]
	}
}
