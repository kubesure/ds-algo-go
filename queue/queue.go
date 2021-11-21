package queue

import "fmt"

func NewQueue(size int) *Queue {
	e := make([]Element, size)
	q := new(Queue)
	q.elements = e
	return q
}

func (q *Queue) EnQueue(value int) *Element {
	if q.size == cap(q.elements) {
		els := make([]Element, 10)
		q.elements = append(q.elements, els...)
	}
	e := Element{value: value}
	q.elements = append(q.elements[:q.size], e)
	q.size++
	return &e
}

func (q *Queue) DeQueue() *Element {
	if q.size <= 0 {
		return nil
	}
	e := q.elements[q.size-1]
	q.elements = q.elements[q.size-1:]
	q.size--
	return &e
}

func (q *Queue) Print() {
	for _, v := range q.elements {
		fmt.Printf(" %v ", v.value)
	}
}
