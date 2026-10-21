package main

type LNode struct {
	Data string
	Next *LNode
	Prev *LNode
}

type DoubleLinkedList struct {
	Head *LNode
	Tail *LNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (d *DoubleLinkedList) Empty() bool {
	return d.Head == nil
}

func (d *DoubleLinkedList) Size() int {
	n := 0
	for c := d.Head; c != nil; c = c.Next {
		n++
	}
	return n
}

func (d *DoubleLinkedList) Clear() {
	d.Head = nil
	d.Tail = nil
}

func (d *DoubleLinkedList) PushFront(str string) {
	node := &LNode{Data: str, Next: d.Head}
	if d.Head == nil {
		d.Tail = node
	} else {
		d.Head.Prev = node
	}
	d.Head = node
}

func (d *DoubleLinkedList) PushBack(str string) {
	node := &LNode{Data: str, Prev: d.Tail}
	if d.Tail == nil {
		d.Head = node
	} else {
		d.Tail.Next = node
	}
	d.Tail = node
}

func (d *DoubleLinkedList) PopFront() (string, bool) {
	if d.Empty() {
		return "", false
	}
	out := d.Head.Data
	d.Head = d.Head.Next
	if d.Head == nil {
		d.Tail = nil
	} else {
		d.Head.Prev = nil
	}
	return out, true
}

func (d *DoubleLinkedList) PopBack() (string, bool) {
	if d.Empty() {
		return "", false
	}
	out := d.Tail.Data
	d.Tail = d.Tail.Prev
	if d.Tail == nil {
		d.Head = nil
	} else {
		d.Tail.Next = nil
	}
	return out, true
}

func (d *DoubleLinkedList) Index(idx int) *LNode {
	if idx < 0 || d.Empty() {
		return nil
	}
	curr := d.Head
	for n := 0; n != idx && curr != nil; n++ {
		curr = curr.Next
	}
	return curr
}

func (d *DoubleLinkedList) InsertBeforeValue(value, str string) bool {
	if d.Empty() {
		return false
	}
	node := d.Head
	for node != nil && node.Data != value {
		node = node.Next
	}
	if node == nil {
		return false
	}
	if node == d.Head {
		d.PushFront(str)
		return true
	}
	left := node.Prev
	newNode := &LNode{Data: str, Next: node, Prev: left}
	left.Next = newNode
	node.Prev = newNode
	return true
}

func (d *DoubleLinkedList) InsertAfterValue(value, str string) bool {
	if d.Empty() {
		return false
	}
	node := d.Head
	for node != nil && node.Data != value {
		node = node.Next
	}
	if node == nil {
		return false
	}
	if node == d.Tail {
		d.PushBack(str)
		return true
	}
	right := node.Next
	newNode := &LNode{Data: str, Next: right, Prev: node}
	node.Next = newNode
	right.Prev = newNode
	return true
}

func (d *DoubleLinkedList) RemoveValue(value string) bool {
	if d.Empty() {
		return false
	}
	node := d.Head
	for node != nil && node.Data != value {
		node = node.Next
	}
	if node == nil {
		return false
	}
	if node == d.Head {
		_, ok := d.PopFront()
		return ok
	}
	if node == d.Tail {
		_, ok := d.PopBack()
		return ok
	}
	left := node.Prev
	right := node.Next
	left.Next = right
	right.Prev = left
	return true
}

func (d *DoubleLinkedList) RemoveBeforeValue(value string) bool {
	if d.Empty() || d.Head.Data == value {
		return false
	}
	node := d.Head
	for node != nil && node.Data != value {
		node = node.Next
	}
	if node == nil {
		return false
	}
	toRemove := node.Prev
	if toRemove == d.Head {
		_, ok := d.PopFront()
		return ok
	}
	left := toRemove.Prev
	right := toRemove.Next
	left.Next = right
	right.Prev = left
	return true
}

func (d *DoubleLinkedList) RemoveAfterValue(value string) bool {
	if d.Empty() {
		return false
	}
	node := d.Head
	for node != nil && node.Data != value {
		node = node.Next
	}
	if node == nil || node.Next == nil {
		return false
	}
	toRemove := node.Next
	if toRemove == d.Tail {
		_, ok := d.PopBack()
		return ok
	}
	left := toRemove.Prev
	right := toRemove.Next
	left.Next = right
	right.Prev = left
	return true
}

func (d *DoubleLinkedList) Find(value string) int {
	n := 0
	for c := d.Head; c != nil; c = c.Next {
		if c.Data == value {
			return n
		}
		n++
	}
	return -1
}

func (d *DoubleLinkedList) Get() []string {
	var result []string
	for c := d.Head; c != nil; c = c.Next {
		result = append(result, c.Data)
	}
	return result
}

func (d *DoubleLinkedList) GetReverse() []string {
	var result []string
	for c := d.Tail; c != nil; c = c.Prev {
		result = append(result, c.Data)
	}
	return result
}

func (d *DoubleLinkedList) GetAfterValue(value string) []string {
	var result []string
	node := d.Head
	for node != nil && node.Data != value {
		node = node.Next
	}
	if node == nil {
		return result
	}
	for c := node.Next; c != nil; c = c.Next {
		result = append(result, c.Data)
	}
	return result
}

func (d *DoubleLinkedList) GetBeforeValue(value string) []string {
	var result []string
	curr := d.Head
	for curr != nil && curr.Data != value {
		result = append(result, curr.Data)
		curr = curr.Next
	}
	if curr == nil {
		return nil
	}
	return result
}
