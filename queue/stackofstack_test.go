package queue

import (
	"testing"
)

func TestSOSPush(t *testing.T) {
	sos := NewStackOfStack(3, 1)
	err := sos.Push(1)

	if err != nil {
		t.Errorf("should have not received an error for pushing 1")
	}

	err = sos.Push(2)

	if err != nil {
		t.Errorf("should have not received an error for pushing 2")
	}

	err = sos.Push(3)

	if err != nil {
		t.Errorf("should have not received an error for pushing 3")
	}

	err = sos.Push(4)

	if err == nil {
		t.Errorf("should have received an error for pushing 4")
	}
}

func TestSOSPop(t *testing.T) {
	sos := NewStackOfStack(3, 1)
	sos.Push(1)
	sos.Push(2)
	sos.Push(3)
	e, err := sos.Pop()
	if err != nil {
		t.Errorf("should not have received an error popping 3")
	} else if err == nil && e.value != 3 {
		t.Errorf("should have poped 3")
	}

	e, err = sos.Pop()
	if err != nil {
		t.Errorf("should not have received an error popping 2")
	} else if err == nil && e.value != 2 {
		t.Errorf("should have poped 2")
	}

	e, err = sos.Pop()
	if err != nil {
		t.Errorf("should not have received an error popping 1")
	} else if err == nil && e.value != 1 {
		t.Errorf("should have poped 1")
	}

	_, err = sos.Pop()
	if err == nil {
		t.Errorf("should have received an error")
	}

}

func TestSOSPeek(t *testing.T) {
	sos := NewStackOfStack(3, 1)
	sos.Push(1)
	sos.Push(2)
	sos.Push(3)
	e, err := sos.Peek()
	if err != nil {
		t.Errorf("should not have received an error")
	} else if err == nil && e.value != 3 {
		t.Errorf("should have peeked 3")
	}
}

func TestSOSPopAt(t *testing.T) {
	sos := NewStackOfStack(3, 1)
	sos.Push(1)
	sos.Push(2)
	sos.Push(3)
	e, err := sos.PopAt(2)
	if err != nil {
		t.Errorf("should not have received an error")
	} else if err == nil && e.value != 3 {
		t.Errorf("should not poped 3")
	}

	e, err = sos.PopAt(1)
	if err != nil {
		t.Errorf("should not have received an error")
	} else if err == nil && e.value != 2 {
		t.Errorf("should not poped 2")
	}

	e, err = sos.PopAt(0)
	if err != nil {
		t.Errorf("should not have received an error")
	} else if err == nil && e.value != 1 {
		t.Errorf("should not poped 1")
	}

	_, err = sos.PopAt(4)
	if err == nil {
		t.Errorf("should have received an error")
	}

	_, err = sos.PopAt(-1)
	if err == nil {
		t.Errorf("should have received an error")
	}
}
