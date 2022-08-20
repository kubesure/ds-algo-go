package queue

func NewMinStack() *MinStack {

	e := make([]Element, 0)
	s := new(Stack)
	s.elements = e
	me := make([]Element, 0)
	m := new(Stack)
	m.elements = me
	ms := MinStack{minstack: m, stack: s}
	return &ms
}

func (m *MinStack) Push(value int) *Element {
	if m.minstack.Peek() == nil {
		m.minstack.Push(value)
	} else if value < m.minstack.Peek().value {
		m.minstack.Push(value)
	}
	return m.stack.Push(value)
}

func (m *MinStack) Pop() *Element {
	if m.stack.Peek().value == m.minstack.Peek().value {
		m.minstack.Pop()
	}
	return m.stack.Pop()
}

func (m *MinStack) Peek() *Element {
	return m.minstack.Peek()
}

func (m *MinStack) Min() *Element {
	return m.minstack.Peek()
}
