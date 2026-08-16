package queue

import gocontainer "github.com/gtantech/go-container"

type Queue[T any] interface {
	Enqueue(v T)
	Dequeue() T
	gocontainer.Container
}

var _ Queue[int] = (*queue[int])(nil) //ensures queue implements Queue at compile time

type queue[T any] struct {
	items []T
	count int
}

func New[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
	q.count++
}

func (q *queue[T]) Dequeue() T {
	//check non-empty
	if q.count == 0 {
		panic("tried to dequeue an empty queue")
	}
	//get the first value in the slice
	v := q.items[0]

	//remove the reference for GC
	var zero T
	q.items[0] = zero
	q.items = q.items[1:]

	//decrement count
	q.count--

	//return value
	return v
}

func (s *queue[T]) IsEmpty() bool {
	return s.count == 0
}

func (s *queue[T]) Size() int {
	return s.count
}
