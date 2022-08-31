package queue

//Element is a value in the stack
type Element struct {
	value int
}

//Stack represent a stack
type Stack struct {
	elements []Element
	size     int
}

type MinStack struct {
	stack    *Stack
	minstack *Stack
}

type MultiStack struct {
	elements   []*Element
	sizes      []int
	sizeOfEach int
	numStacks  int
}

type StackOfStack struct {
	stacks     []*Stack
	sizeof     []int
	currStack  int
	capacity   int
	sizeOfEach int
}

//Stack represent a stack
type Queue struct {
	elements []Element
	size     int
}

type QueueStack struct {
	new *Stack
	old *Stack
}
