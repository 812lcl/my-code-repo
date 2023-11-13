package leetcode

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
		for j := (i + 1) % len(nums); j != i; {
			if nums[j] > nums[i] {
				res[i] = nums[j]
				break
			}
			j = (j + 1) % len(nums)
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

func nextGreaterElements1(nums []int) []int {
	res := make([]int, len(nums))
	var st stack
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	for i := 0; i < 2*len(nums); i++ {
		num := nums[i%len(nums)]
		for !st.Empty() && nums[st.Top()] < num {
			res[st.Pop()] = num
		}
		st.Push(i % len(nums))
	}
	return res
}
