package queue

import (
	"generalization_tool"
	"generalization_tool/internal/queue"
)

type PriorityQueue[T any] struct {
	priorityQueue *queue.PriorityQueue[T]
}

func NewPriorityQueue[T any](capacity int, compare generalization_tool.Comparator[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	pq.priorityQueue = queue.NewPriorityQueue[T](capacity, compare)
	return pq
}

func (p *PriorityQueue[T]) Len() int {
	return p.priorityQueue.Len()
}

func (p *PriorityQueue[T]) Peek() (T, error) {
	return p.priorityQueue.Peek()
}

func (p *PriorityQueue[T]) Enqueue(t T) error {
	return p.priorityQueue.Enqueue(t)
}

func (p *PriorityQueue[T]) Dequeue() (T, error) {
	return p.priorityQueue.Dequeue()
}
