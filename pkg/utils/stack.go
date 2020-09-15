package utils

import "errors"

type Stack struct {
	s []rune
}

func NewStack() *Stack {
	return &Stack{
		s: make([]rune, 0),
	}
}

func (s *Stack) Size() int {
	return len(s.s)
}

func (s *Stack) Push(v rune) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() (rune, error) {
	l := len(s.s)
	if l == 0 {
		return -1, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *Stack) Top() (rune, error) {
	l := len(s.s)
	if l == 0 {
		return -1, errors.New("Empty Stack")
	}

	return s.s[l-1], nil
}
