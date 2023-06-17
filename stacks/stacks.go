package stacks

import "errors"

type Stack struct {
	Elements []int
}

func (s *Stack) Push(elem int) {
	s.Elements = append(s.Elements, elem)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("empty stack")
	}

	var lastElement int

	index := len(s.Elements) - 1
	lastElement, s.Elements = s.Elements[index], s.Elements[:index]

	return lastElement, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("empty stack")
	}

	index := len(s.Elements) - 1

	return s.Elements[index], nil
}

func (s *Stack) Length() int {
	return len(s.Elements)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}
