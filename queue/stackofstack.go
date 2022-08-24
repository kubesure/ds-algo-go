package queue

import "fmt"

func NewStackOfStack(numStack, stackSize int) *StackOfStack {

	sizes := make([]int, numStack)
	stacks := make([]*Stack, 0)

	stacks = addStack(numStack, stacks)

	return &StackOfStack{
		stacks:     stacks,
		sizeof:     sizes,
		capacity:   numStack * stackSize,
		sizeOfEach: stackSize,
	}
}

func addStack(numStack int, stacks []*Stack) []*Stack {
	for i := 0; i < numStack; i++ {
		s := Stack{}
		stacks = append(stacks, &s)
	}
	return stacks
}

func (sos *StackOfStack) Push(value int) error {

	if sos.capacity == 0 {
		return fmt.Errorf("Stack capacity is full pop some")
	}

	for stack, size := range sos.sizeof {
		if size != sos.sizeOfEach {
			sos.stacks[stack].Push(value)
			sos.capacity--
			sos.sizeof[stack]++
			sos.currStack = stack
			break
		}
	}

	return nil
}

func (sos *StackOfStack) Pop() (*Element, error) {

	if sos.IsEmpty() {
		return nil, fmt.Errorf("Stack is empty push some")
	}

	e := sos.stacks[sos.currStack].Pop()
	sos.capacity++

	sos.sizeof[sos.currStack]--
	if sos.sizeof[sos.currStack] == 0 {
		if sos.currStack != 0 {
			sos.currStack--
		}

	}

	return e, nil
}

func (sos *StackOfStack) Peek() (*Element, error) {

	if sos.IsEmpty() {
		return nil, fmt.Errorf("Stack is empty push some")
	}

	return sos.stacks[sos.currStack].Peek(), nil
}

func (sos *StackOfStack) PopAt(sindex int) (*Element, error) {

	if sindex < 0 || sindex > len(sos.stacks) {
		return nil, fmt.Errorf("stack index incorrect")
	}

	if sos.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	e := sos.stacks[sindex].Pop()
	sos.capacity++
	sos.sizeof[sindex]--

	if sos.sizeof[sindex] == 0 {
		if sos.currStack != 0 {
			sos.currStack--
		}
	}

	return e, nil
}

func (sos *StackOfStack) IsEmpty() bool {
	var elements int
	for _, stack := range sos.sizeof {
		elements += stack
	}
	return elements == 0
}

func (sos *StackOfStack) IsFull() bool {
	return sos.capacity == 0
}

func (sos *StackOfStack) Print() {
	for _, stack := range sos.stacks {
		stack.Print()
	}
}
