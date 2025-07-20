package datastructures

import (
	"errors"
)

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.data) > 0 {
		response := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return response, nil
	} else {
		var zero T
		return zero, errors.New("no elements in the stack")
	}
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.data) == 0 {
		var zero T
		return zero, errors.New("the stack is empty")
	} else {
		return s.data[len(s.data)-1], nil
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}
