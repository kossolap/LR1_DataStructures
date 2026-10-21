package main

type QNode struct {
	Data string
	Next *QNode
}

type Queue struct {
	Head *QNode
	Tail *QNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Empty() bool {
	return q.Head == nil
}

func (q *Queue) Push(str string) {
	node := &QNode{Data: str}
	if q.Tail == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

func (q *Queue) Pop() (string, bool) {
	if q.Empty() {
		return "", false
	}
	out := q.Head.Data
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return out, true
}

func (q *Queue) Get() []string {
	var result []string
	curr := q.Head
	for curr != nil {
		result = append(result, curr.Data)
		curr = curr.Next
	}
	return result
}

func (q *Queue) Clear() {
	q.Head = nil
	q.Tail = nil
}
