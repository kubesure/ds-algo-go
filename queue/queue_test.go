package queue

import (
	"testing"
)

func TestQueueCreate(t *testing.T) {
	q := NewQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)

	e := q.DeQueue()
	if e.value != 1 {
		t.Errorf("wanted 1 got %v", e.value)
	}

	e = q.DeQueue()
	if e.value != 2 {
		t.Errorf("wanted 2 got %v", e.value)
	}

	e = q.DeQueue()

	if e.value != 3 {
		t.Errorf("wanted 3 got %v", e.value)
	}

	e = q.DeQueue()
	if e.value != 4 {
		t.Errorf("wanted 4 got %v", e.value)
	}

	e = q.DeQueue()
	if e.value != 5 {
		t.Errorf("wanted 5 got %v", e.value)
	}

	val := q.DeQueue()

	if val != nil {
		t.Errorf("should be nil but got %v", val)
	}
}
