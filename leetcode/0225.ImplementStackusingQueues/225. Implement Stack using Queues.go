package leetcode

type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{q1: []int{}, q2: []int{}}

}

func (s *MyStack) Push(x int) {
	if len(s.q1) != 0 {
		s.q1 = append(s.q1, x)
	} else {
		s.q2 = append(s.q2, x)
	}
}

func (s *MyStack) Pop() int {
	if s.Empty() {
		return 0
	}
	if len(s.q1) != 0 {
		n := len(s.q1)
		res := s.q1[n-1]
		for i := 0; i < n-1; i++ {
			s.q2 = append(s.q2, s.q1[i])
		}
		s.q1 = nil
		return res
	} else {
		n := len(s.q2)
		res := s.q2[n-1]
		for i := 0; i < n-1; i++ {
			s.q1 = append(s.q1, s.q2[i])
		}
		s.q2 = nil
		return res
	}
}

func (s *MyStack) Top() int {
	if s.Empty() {
		return 0
	}
	if len(s.q1) != 0 {
		n := len(s.q1)
		return s.q1[n-1]
	} else {
		n := len(s.q2)
		return s.q2[n-1]
	}
}

func (s *MyStack) Empty() bool {
	return len(s.q1)+len(s.q2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
