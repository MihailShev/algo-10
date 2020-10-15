package main

type Stack struct {
	stack []int
}

func (s *Stack) Push(item int) {
	s.stack = append(s.stack, item)
}

func (s *Stack) Pop() int  {
	l := len(s.stack)

	if l == 0 {
		panic("stack is empty")
	}

	item := s.stack[l - 1]

	s.stack = s.stack[: l - 1]

	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}