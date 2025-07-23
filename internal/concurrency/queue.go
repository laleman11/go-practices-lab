package concurrency

import "sync"

type ConcurrentQueue[T comparable] struct {
	mu   sync.Mutex
	data []T
}

func NewConcurrentQueue[T comparable]() *ConcurrentQueue[T] {
	return &ConcurrentQueue[T]{}
}

func (q *ConcurrentQueue[T]) Enqueue(value T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, value)
}

func (q *ConcurrentQueue[T]) Dequeue() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	var response T
	if len(q.data) != 0 {
		response = q.data[0]
		q.data = q.data[1:]
		return response, true
	}
	return response, false
}
