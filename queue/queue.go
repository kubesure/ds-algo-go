package queue

import "fmt"

/*func NewQueue(size int) *Queue {
	e := make([]Element, size)
	q := new(Queue)
	q.elements = e
	return q
}*/

func NewQueue() *Queue {
	q := new(Queue)
	return q
}

func (q *Queue) EnQueue(value int) *Element {

	e := Element{value: value}
	q.elements = append(q.elements, e)
	q.size++
	return &e
}

func (q *Queue) DeQueue() *Element {
	if q.size <= 0 {
		return nil
	}
	//e := q.elements[q.size-1]
	e := q.elements[0]
	//q.elements = q.elements[0 : q.size-1]
	q.elements = q.elements[1:]
	q.size--
	return &e
}

func (q *Queue) Print() {
	for _, v := range q.elements {
		fmt.Printf(" %v ", v.value)
	}
}
