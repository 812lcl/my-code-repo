package leetcode

import "math"

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

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	num3, st := math.MinInt64, make(stack, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < num3 {
			return true
		}
		for !st.Empty() && nums[i] > st.Top() {
			// st.Top() is num3, nums[i] is num2, so just need to find if there is a num1 < num3
			num3 = st.Pop()
		}
		st.Push(nums[i])
	}
	return false
}
