package queue

import "context"

// Queue 普通队列
type Queue[T any] interface {
	// Enqueue 将元素放入队列，如果此时队列已满，那么返回错误
	Enqueue(t T) error
	// Dequeue 从队首获得一个元素
	Dequeue() (T, error)
}

// BlockingQueue 阻塞队列
type BlockingQueue[T any] interface {
	// Enqueue
	// 1.将元素放入队列。如果在 ctx 超时之前，队列有空闲位置，那么元素会被放入队列；否则 err
	// 2.在超时或主动 cancel 的情况下，所有的实现都必须返回 ctx
	// 3.调用者可以通过检查 error 是否为 context.DeadlineExceeded
	//   或 context.Canceled 来判断入队失败的原因
	// 4.调用者必须使用 errors.Is 来判断，而不能直接使用 ==
	Enqueue(ctx context.Context, t T) error

	// Dequeue
	// 1.从队首获得一个元素
	// 2.如果在 ctx 超时之前，队列中有元素，那么会返回对首位的元素，否则err
	// 3.在超时或调用者主动 cancel 的情况下，所有的实现都必须返回 ctx
	// 4.调用者可以通过检查 error 是否为 context.DeadlineExceeded
	//   或 context.Canceled 来判断入队失败的原因
	// 5.调用者必须使用 errors.Is 来判断，而不能直接使用 ==
	Dequeue(ctx context.Context) (T, error)
}
