package utils

import "fmt"

type Item interface{}

type Stack struct {
	items []Item
}

func NewStack() *Stack {
	return &Stack{
		items: []Item{},
	}
}

func (s *Stack) Push(item Item) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() Item {
	if s.IsEmpty() {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() Item {
	if s.IsEmpty() {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = []Item{}
}

func (s *Stack) ToSlice() []Item {
	return s.items
}

func (s *Stack) Copy() *Stack {
	newStack := NewStack()
	newStack.items = append(newStack.items, s.items...)
	return newStack
}

func (s *Stack) Reverse() {
	for i, j := 0, len(s.items)-1; i < j; i, j = i+1, j-1 {
		s.items[i], s.items[j] = s.items[j], s.items[i]
	}
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.items)
}
