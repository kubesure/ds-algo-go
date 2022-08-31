package queue

import "fmt"

func NewSortedStack() *SortedStack {
	return &SortedStack{sorted: &Stack{}, temp: &Stack{}}
}

func (s *SortedStack) Push(value int) {
	//create a temp stack of values greater than pushed value

	if !s.sorted.IsEmpty() {
		for {
			peek := s.sorted.Peek()
			if peek == nil {
				break
			}
			if peek.value > value {
				pop := s.sorted.Pop()
				s.temp.Push(pop.value)
			} else {
				break
			}
		}
	}
	//push the current value at the position which makes it top
	s.sorted.Push(value)
	// push values greater than pushed value from temp to sorted stack
	if !s.temp.IsEmpty() {
		for {
			pop := s.temp.Pop()
			if pop == nil {
				break
			}
			s.sorted.Push(pop.value)
		}
	}
}

func (s *SortedStack) Pop() (*Element, error) {
	if !s.sorted.IsEmpty() {
		return s.sorted.Pop(), nil
	} else {
		return nil, fmt.Errorf("Stack is empty")
	}
}

func (s *SortedStack) Peek() (*Element, error) {
	if !s.sorted.IsEmpty() {
		return s.sorted.Peek(), nil
	} else {
		return nil, fmt.Errorf("Stack is empty")
	}
}

func (s *SortedStack) Print() {
	s.sorted.Print()
}
