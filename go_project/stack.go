package main

type SNode struct {
	Data string
	Next *SNode
}

type Stack struct {
	Head *SNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Empty() bool {
	return s.Head == nil
}

func (s *Stack) Push(str string) {
	s.Head = &SNode{Data: str, Next: s.Head}
}

func (s *Stack) Pop() (string, bool) {
	if s.Empty() {
		return "", false
	}
	out := s.Head.Data
	s.Head = s.Head.Next
	return out, true
}

func (s *Stack) Get() []string {
	var result []string
	curr := s.Head
	for curr != nil {
		result = append(result, curr.Data)
		curr = curr.Next
	}
	return result
}

func (s *Stack) Clear() {
	s.Head = nil
}
