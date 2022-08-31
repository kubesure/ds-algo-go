package queue

import "fmt"

func NewQueueStack() *QueueStack {
	return &QueueStack{new: &Stack{}, old: &Stack{}}
}

func (q *QueueStack) EnQueue(e int) {
	q.new.Push(e)
}

func (q *QueueStack) DeQueue() (*Element, error) {
	if q.old.IsEmpty() {
		if q.new.IsEmpty() {
			return nil, fmt.Errorf("cannot dequeue queue empty")
		}
		q.copyNewToOld()
		return q.old.Pop(), nil
	}
	return q.old.Pop(), nil
}

func (q *QueueStack) Peek() (*Element, error) {
	if q.old.IsEmpty() {
		if q.new.IsEmpty() {
			return nil, fmt.Errorf("cannot dequeue queue empty")
		}
		q.copyNewToOld()
		return q.old.Peek(), nil
	}
	return q.old.Pop(), nil
}

func (q *QueueStack) copyNewToOld() {
	for {
		e := q.new.Pop()
		if e == nil {
			break
		}
		q.old.Push(e.value)
	}
}

func (q *QueueStack) Print() {
	q.new.Print()
	q.old.Print()
}
