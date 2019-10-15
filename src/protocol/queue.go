package protocol

type queue struct {
	first *element
	last  *element
}

func (q *queue) Enqueue(value interface{}) {
	newElement := &element{
		next:  nil,
		value: value,
	}
	if q.IsEmpty() {
		q.first = newElement
		q.last = newElement
	} else {
		q.last.next = newElement
		q.last = newElement
	}
}

func (q *queue) PushFront(value interface{}) {
	newElement := &element{
		next:  q.first,
		value: value,
	}
	if q.IsEmpty() {
		q.first = newElement
		q.last = newElement
	} else {
		q.first = newElement
	}
}

func (q *queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	elem := q.first
	q.first = elem.next
	return elem.value
}

func (q *queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.first.value
}

func (q *queue) IsEmpty() bool {
	return q.first == nil
}

type element struct {
	next  *element
	value interface{}
}
