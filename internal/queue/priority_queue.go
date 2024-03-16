package queue

import (
	"errors"
	"generalization_tool"
	"generalization_tool/internal/slice"
)

var (
	ErrOutOfCapacity = errors.New("generalization_tool：超出最大容量限制！")
	ErrEmptyQueue    = errors.New("generalization_tool：队列为空！")
)

// PriorityQueue 基于小顶堆的优先级队列
// 当capacity <= 0时，为无界队列，切片容量会动态扩缩容
// 当capacity > 0 时，为有界队列，初始化后就固定容量，不会扩缩容
type PriorityQueue[T any] struct {
	capacity int                               // 队列容量
	data     []T                               // 队列元素
	compare  generalization_tool.Comparator[T] // 比较前一个元素是否小于后一个元素
}

// NewPriorityQueue 创建优先队列 capacity <= 0 时，为无界队列，否则有有界队列
func NewPriorityQueue[T any](capacity int, compare generalization_tool.Comparator[T]) *PriorityQueue[T] {
	sliceCap := capacity + 1
	if capacity < 1 {
		capacity = 0
		sliceCap = 64
	}
	return &PriorityQueue[T]{
		capacity: capacity,
		data:     make([]T, 1, sliceCap),
		compare:  compare,
	}
}

func (p *PriorityQueue[T]) Len() int {
	return len(p.data) - 1
}

func (p *PriorityQueue[T]) Cap() int {
	return p.capacity
}

func (p *PriorityQueue[T]) IsBoundless() bool {
	return p.capacity <= 0
}

func (p *PriorityQueue[T]) isFull() bool {
	return p.capacity > 0 && len(p.data)-1 == p.capacity
}

func (p *PriorityQueue[T]) isEmpty() bool {
	return p.capacity < 2
}

func (p *PriorityQueue[T]) Peek() (T, error) {
	if p.isEmpty() {
		var t T
		return t, ErrEmptyQueue
	}
	return p.data[1], nil
}

func (p *PriorityQueue[T]) Enqueue(t T) error {
	if p.isFull() {
		return ErrOutOfCapacity
	}

	p.data = append(p.data, t)
	node, parent := len(p.data)-1, (len(p.data)-1)/2
	// 如果新元素比父节点小，就交换他们的位置，直到找到合适的位置
	for parent > 0 && p.compare(p.data[node], p.data[parent]) < 0 {
		p.data[node], p.data[parent] = p.data[parent], p.data[node]
		node = parent
		parent = parent / 2
	}
	return nil
}

func (p *PriorityQueue[T]) Dequeue() (T, error) {
	if p.isEmpty() {
		var t T
		return t, ErrEmptyQueue
	}

	pop := p.data[1]
	p.data[1] = p.data[len(p.data)-1]
	p.data = p.data[:len(p.data)-1]
	p.shrinkIfNec()
	p.heapify(p.data, len(p.data)-1, 1)
	return pop, nil
}

func (p *PriorityQueue[T]) shrinkIfNec() {
	if p.IsBoundless() {
		p.data = slice.Shrink[T](p.data)
	}
}

// 维护小顶堆
func (p *PriorityQueue[T]) heapify(data []T, n, i int) {
	minPos := i
	for {
		// 如果左孩子存在并且小于 minPos 索引处的元素，就更新 minPos 为左孩子的索引
		if left := i * 2; left <= n && p.compare(data[left], data[minPos]) < 0 {
			minPos = left
		}
		// 如果右孩子存在并且小于 minPos 索引处的元素，就更新 minPos 为右孩子的索引
		if right := i*2 + 1; right <= n && p.compare(data[right], data[minPos]) < 0 {
			minPos = right
		}
		// 如果 minPos 仍然等于原始索引 i，表示堆的性质已满足，循环退出
		if minPos == i {
			break
		}
		data[i], data[minPos] = data[minPos], data[i]
		i = minPos
	}
}
