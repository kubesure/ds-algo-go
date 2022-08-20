package queue

func NewStackOfStack(numStack int) *StackOfStack {
	return &StackOfStack{
		stacks:    []*Stack{{}},
		sizeof:    []int{0},
		stackMax:  numStack,
		currStack: 0,
	}
}

func (sos StackOfStack) Push(value int) {

}

func (sos StackOfStack) Pop() int {
	return 0
}

func (sos StackOfStack) PopAt(index int) int {
	return 0
}

func (sos StackOfStack) Peek() int {
	return 0
}

func (sos StackOfStack) IsEmpty() bool {
	return sos.sizeof[sos.currStack] == 0
}
