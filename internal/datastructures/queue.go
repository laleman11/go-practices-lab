package datastructures

import "errors"

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: []T{}}
}

func (q *Queue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.data) == 0 {
		var zero T
		return zero, errors.New("there's no elements in the queue")
	} else {
		response := q.data[0]
		q.data = append(q.data[1:len(q.data)])
		return response, nil
	}
}

func (q *Queue[T]) Peek() (T, error) {
	if len(q.data) == 0 {
		var zero T
		return zero, errors.New("there's no elements in the queue")
	} else {
		response := q.data[0]
		return response, nil
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}
