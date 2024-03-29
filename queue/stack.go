package queue

import (
	"fmt"
)

func NewStack() *Stack {
	e := make([]Element, 0)
	s := new(Stack)
	s.elements = e
	return s
}

//Push pushes value on top of stack
func (s *Stack) Push(value int) *Element {
	e := Element{value: value}

	s.elements = append(s.elements, e)
	s.size++
	return &e
}

//Pop pops the top element
func (s *Stack) Pop() *Element {
	if s.size <= 0 {
		return nil
	}

	index := s.size - 1
	e := s.elements[index]
	s.elements = s.elements[:index]
	s.size--
	return &e
}

//Peek returns top to stack element
func (s *Stack) Peek() *Element {
	if s.size <= 0 {
		return nil
	}

	l := s.size
	e := s.elements[l-1]
	return &e
}

//Size return size of stack
func (s *Stack) Size() int {
	return s.size
}

//IsEmpty return size of stack
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Print method on Element class
func (s *Stack) Print() {
	for _, v := range s.elements {
		fmt.Printf(" %v ", v.value)
	}
}
