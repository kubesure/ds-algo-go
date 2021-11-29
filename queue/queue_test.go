package queue

import "testing"

func TestQueueCreate(t *testing.T) {
	q := NewQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	//q.Print()
	q.DeQueue()
	q.Print()
}