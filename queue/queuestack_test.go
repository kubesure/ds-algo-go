package queue

import "testing"

func TestEnQueue(t *testing.T) {
	q := NewQueueStack()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	e, err := q.DeQueue()
	if err != nil {
		t.Errorf("Shold not be nil when 1 is dequeue")
	} else if e.value != 1 {
		t.Errorf("wanted 1 got %v ", e.value)
	}
	e, err = q.DeQueue()
	if err != nil {
		t.Errorf("Shold not be nil 2 is dequeue")
	} else if e.value != 2 {
		t.Errorf("wanted 2 got %v ", e.value)
	}

	q.EnQueue(4)

	e, err = q.DeQueue()
	if err != nil {
		t.Errorf("Shold not be nil when 3 is dequeue")
	} else if e.value != 3 {
		t.Errorf("wanted 3 got %v ", e.value)
	}

	e, err = q.DeQueue()
	if err != nil {
		t.Errorf("Shold not be nil 4 is dequeue")
	} else if e.value != 4 {
		t.Errorf("wanted 4 got %v ", e.value)
	}

	_, err = q.DeQueue()
	if err == nil {
		t.Errorf("Should be an error as there are not elements to dequeue")
	}
}
