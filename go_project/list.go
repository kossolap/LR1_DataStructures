package main

type FNode struct {
	Data string
	Next *FNode
}

type LinkedList struct {
	Head *FNode
	Tail *FNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Empty() bool {
	return l.Head == nil
}

func (l *LinkedList) Size() int {
	n := 0
	for c := l.Head; c != nil; c = c.Next {
		n++
	}
	return n
}

func (l *LinkedList) Clear() {
	l.Head = nil
	l.Tail = nil
}

func (l *LinkedList) PushFront(str string) {
	node := &FNode{Data: str}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

func (l *LinkedList) PushBack(str string) {
	node := &FNode{Data: str}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *LinkedList) PopFront() (string, bool) {
	if l.Empty() {
		return "", false
	}
	out := l.Head.Data
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	}
	return out, true
}

func (l *LinkedList) PopBack() (string, bool) {
	if l.Empty() {
		return "", false
	}
	out := l.Tail.Data
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		return out, true
	}
	prev := l.Head
	for prev.Next != l.Tail {
		prev = prev.Next
	}
	l.Tail = prev
	l.Tail.Next = nil
	return out, true
}

func (l *LinkedList) Index(idx int) *FNode {
	if idx < 0 || l.Empty() {
		return nil
	}
	curr := l.Head
	for n := 0; n != idx && curr != nil; n++ {
		curr = curr.Next
	}
	return curr
}

func (l *LinkedList) InsertBeforeValue(value, str string) bool {
	if l.Empty() {
		return false
	}
	if l.Head.Data == value {
		l.PushFront(str)
		return true
	}
	prev := l.Head
	curr := l.Head.Next
	for curr != nil && curr.Data != value {
		prev = curr
		curr = curr.Next
	}
	if curr == nil {
		return false
	}
	prev.Next = &FNode{Data: str, Next: curr}
	return true
}

func (l *LinkedList) InsertAfterValue(value, str string) bool {
	if l.Empty() {
		return false
	}
	curr := l.Head
	for curr != nil && curr.Data != value {
		curr = curr.Next
	}
	if curr == nil {
		return false
	}
	newNode := &FNode{Data: str, Next: curr.Next}
	curr.Next = newNode
	if newNode.Next == nil {
		l.Tail = newNode
	}
	return true
}

func (l *LinkedList) RemoveValue(value string) bool {
	if l.Empty() {
		return false
	}
	if l.Head.Data == value {
		_, ok := l.PopFront()
		return ok
	}
	prev := l.Head
	curr := l.Head.Next
	for curr != nil {
		if curr.Data == value {
			if curr == l.Tail {
				_, ok := l.PopBack()
				return ok
			}
			prev.Next = curr.Next
			return true
		}
		prev = curr
		curr = curr.Next
	}
	return false
}

func (l *LinkedList) RemoveBeforeValue(value string) bool {
	if l.Empty() || l.Head == l.Tail {
		return false
	}
	if l.Head.Data == value {
		return false
	}
	var prevPrev *FNode
	prev := l.Head
	curr := l.Head.Next
	for curr != nil && curr.Data != value {
		prevPrev = prev
		prev = curr
		curr = curr.Next
	}
	if curr == nil {
		return false
	}
	if prevPrev == nil {
		_, ok := l.PopFront()
		return ok
	}
	prevPrev.Next = prev.Next
	return true
}

func (l *LinkedList) RemoveAfterValue(value string) bool {
	if l.Empty() {
		return false
	}
	curr := l.Head
	for curr != nil && curr.Data != value {
		curr = curr.Next
	}
	if curr == nil || curr == l.Tail {
		return false
	}
	toDelete := curr.Next
	if toDelete == l.Tail {
		l.Tail = curr
		l.Tail.Next = nil
		return true
	}
	curr.Next = toDelete.Next
	return true
}

func (l *LinkedList) Find(value string) int {
	n := 0
	for c := l.Head; c != nil; c = c.Next {
		if c.Data == value {
			return n
		}
		n++
	}
	return -1
}

func (l *LinkedList) Get() []string {
	var result []string
	for c := l.Head; c != nil; c = c.Next {
		result = append(result, c.Data)
	}
	return result
}

func (l *LinkedList) GetAfterValue(value string) []string {
	var result []string
	if l.Empty() {
		return result
	}
	curr := l.Head
	for curr != nil && curr.Data != value {
		curr = curr.Next
	}
	if curr == nil {
		return result
	}
	for c := curr.Next; c != nil; c = c.Next {
		result = append(result, c.Data)
	}
	return result
}

func (l *LinkedList) GetBeforeValue(value string) []string {
	var result []string
	if l.Empty() {
		return result
	}
	curr := l.Head
	for curr != nil && curr.Data != value {
		result = append(result, curr.Data)
		curr = curr.Next
	}
	if curr == nil {
		return nil
	}
	return result
}
