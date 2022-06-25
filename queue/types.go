package queue

//Stack represent a stack
type Stack struct {
	elements []Element
	size     int
}

//Stack represent a stack
type Queue struct {
	elements []Element
	size     int
}

//Element is a value in the stack
type Element struct {
	value int
}

type MultiStack struct {
	elements   []*Element
	sizes      []int
	sizeOfEach int
	numStacks  int
}
