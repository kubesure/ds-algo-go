package queue

import "testing"

func TestQueueCreate(t *testing.T) {
	q := NewQueue(2)
	q.EnQueue(1)
	q.EnQueue(2)
	q.Print()
	q.DeQueue()
	q.Print()
}
