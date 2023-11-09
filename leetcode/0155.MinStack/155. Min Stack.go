package leetcode

type MinStack struct {
	data   []int
	min    []int
	length int
}

func Constructor() MinStack {
	return MinStack{
		data:   []int{},
		min:    []int{},
		length: 0,
	}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)
	if s.length == 0 {
		s.min = append(s.min, val)
	} else {
		min := s.GetMin()
		if val < min {
			s.min = append(s.min, val)
		} else {
			s.min = append(s.min, min)
		}
	}
	s.length++
}

func (s *MinStack) Pop() {
	s.length--
	s.data = s.data[:s.length]
	s.min = s.min[:s.length]
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	return s.min[s.length-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
