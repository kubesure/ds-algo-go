package queue

import "testing"

func TestMinStack(t *testing.T) {
	ms := NewMinStack()
	ms.Push(50)
	ms.Push(25)
	ms.Push(10)
	if ms.Min().value != 10 {
		t.Errorf("Should have been 10 after push")
	}

	ms.Pop()
	if ms.Peek().value != 25 {
		t.Errorf("Should have been 25 after 1 pop")
	}

	ms.Pop()

	if ms.Min().value != 50 {
		t.Errorf("Should have been 50 after 2 pops")
	}

}
