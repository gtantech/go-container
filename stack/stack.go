package stack

import "github.com/gtantech/go-container/container"

type Stack[T any] interface {
	Push(v T)
	Pop() T
	Peek() T
	container.Container
}

var _ Stack[int] = (*stack[int])(nil) //ensures stack implements Stack at compile time

type stack[T any] struct {
	items []T
}

func New[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *stack[T]) Pop() T {
	//check non-empty
	if len(s.items) == 0 {
		panic("tried to pop an empty stack")
	}
	//get the last value in the slice
	v := s.items[len(s.items)-1]

	//remove the reference for GC
	var zero T
	s.items[len(s.items)-1] = zero
	s.items = s.items[:len(s.items)-1]

	//return value
	return v
}

func (s *stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *stack[T]) Size() int {
	return len(s.items)
}
