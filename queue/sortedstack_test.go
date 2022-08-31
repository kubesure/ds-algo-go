package queue

import "testing"

func TestSortedStack(t *testing.T) {
	ss := NewSortedStack()

	ss.Push(1)
	ss.Push(10)

	e, err := ss.Peek()
	if err != nil {
		t.Errorf("Should not be nil to peek value 10")
	} else if e.value != 10 {
		t.Errorf("wanted 10 got %v ", e.value)
	}

	ss.Push(12)
	e, err = ss.Peek()
	if err != nil {
		t.Errorf("Should not be nil to peek value 12")
	} else if e.value != 12 {
		t.Errorf("wanted 12 got %v ", e.value)
	}

	ss.Push(4)

	e, err = ss.Pop()
	if err != nil {
		t.Errorf("Should not be nil to peek value 12")
	} else if e.value != 12 {
		t.Errorf("wanted 12 got %v ", e.value)
	}

	ss.Print()
}
