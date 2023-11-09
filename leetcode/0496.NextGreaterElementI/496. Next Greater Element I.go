package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	index := -1
	for j, v := range nums1 {
		res[j] = -1
		for i := 0; i < len(nums2); i++ {
			if nums2[i] == v {
				index = i
			}
		}
		for i := index; i < len(nums2); i++ {
			if nums2[i] > v {
				res[j] = nums2[i]
				index = -1
				break
			}
		}
	}
	return res
}

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	if s.Empty() {
		return 0
	}
	res := (*s)[s.Length()-1]
	*s = (*s)[:s.Length()-1]
	return res
}

func (s *stack) Top() int {
	if s.Empty() {
		return 0
	}
	return (*s)[s.Length()-1]
}

func (s *stack) Empty() bool {
	return s.Length() == 0
}

func (s *stack) Length() int {
	return len(*s)
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	var st stack
	greater := make(map[int]int)
	for _, v := range nums2 {
		for !st.Empty() && st.Top() < v {
			greater[st.Pop()] = v
		}
		st.Push(v)
	}

	for i, v := range nums1 {
		if num, ok := greater[v]; ok {
			res[i] = num
		} else {
			res[i] = -1
		}
	}
	return res
}
