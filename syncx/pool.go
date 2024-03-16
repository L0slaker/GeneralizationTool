package syncx

import "sync"

// Pool 对象池，对sync.Pool的简单封装
type Pool[T any] struct {
	p sync.Pool
}

// NewPool 创建一个实例，factor必须返回T类型的值，不能返回nil
func NewPool[T any](factor func() T) *Pool[T] {
	return &Pool[T]{
		p: sync.Pool{
			New: func() any {
				return factor()
			},
		},
	}
}

// Get 取出一个元素
func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}

// Put 放入一个元素
func (p *Pool[T]) Put(t T) {
	p.p.Put(t)
}
