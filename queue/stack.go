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

	e := s.elements[s.size-1]
	s.elements = s.elements[:s.size-1]
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

func NewMultiStack(numStacks, sizeOfEach int) *MultiStack {
	e := make([]*Element, numStacks*sizeOfEach)
	s := make([]int, numStacks)
	ms := MultiStack{elements: e, sizeOfEach: sizeOfEach, numStacks: numStacks, sizes: s}
	return &ms
}

func (m *MultiStack) head(stack int) int {
	head := m.sizeOfEach*stack - m.sizes[stack-1] - 1
	if head < 0 {
		return 0
	}
	return head
}

func (m *MultiStack) checkStack(stack int) error {
	if stack <= 0 || stack > m.numStacks {
		return fmt.Errorf("Stack does not exists")
	}
	return nil
}

func (m *MultiStack) isFull(stack int) error {
	index := stack - 1
	if m.sizes[index] == m.sizeOfEach {
		return fmt.Errorf("Stack is full")
	}
	return nil
}

func (m *MultiStack) Push(stack, value int) error {

	err := m.checkStack(stack)
	if err != nil {
		return err
	}

	err1 := m.isFull(stack)

	if err1 != nil {
		return err1
	}

	m.elements[m.head(stack)] = &Element{value}
	m.sizes[stack-1]++
	return nil
}

func (m *MultiStack) Pop(stack int) (*Element, error) {
	err := m.checkStack(stack)

	if err != nil {
		return nil, err
	}
	e := m.elements[m.head(stack)]
	m.sizes[stack-1]--
	m.elements[m.head(stack)] = nil
	return e, nil
}

func (m *MultiStack) Peek(stack int) (*Element, error) {
	err := m.checkStack(stack)

	if err != nil {
		return nil, err
	}

	if m.Size(stack) == 0 {
		return nil, fmt.Errorf("Stack is empty")
	}

	e := m.elements[m.head(stack)]
	return e, nil
}

func (m *MultiStack) Print() {
	if m.FullSize() > 0 {
		for _, v := range m.elements {
			if v != nil {
				fmt.Printf(" %v ", v.value)
			}
		}
	}
}

func (m *MultiStack) Size(stack int) int {
	return m.sizes[stack-1]
}

func (m *MultiStack) FullSize() int {
	if m.elements != nil || len(m.elements) > 0 {
		return len(m.elements)
	}
	return 0
}
