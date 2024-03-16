package queue

import (
	"context"
	"fmt"
	"generalization_tool/internal/queue"
	"sync"
	"time"
)

type DelayQueue[T Delayable] struct {
	q             queue.PriorityQueue[T]
	mutex         *sync.Mutex
	dequeueSignal *cond
	enqueueSignal *cond
}

func (d *DelayQueue[T]) Enqueue(ctx context.Context, t T) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		d.mutex.Lock()
		err := d.q.Enqueue(t)
		switch err {
		case nil:
			d.enqueueSignal.broadcast()
			return nil
		case queue.ErrOutOfCapacity:
			signal := d.dequeueSignal.signalCh()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-signal:
			}
		default:
			d.mutex.Unlock()
			return fmt.Errorf("generalization_tool：延时队列入队时遇到未知问题 %w，请上报", err)
		}
	}
}

func (d *DelayQueue[T]) Dequeue(ctx context.Context) (T, error) {
	var timer *time.Timer
	defer func() {
		if timer != nil {
			timer.Stop()
		}
	}()
	for {
		select {
		case <-ctx.Done():
			var t T
			return t, ctx.Err()
		default:
		}
		d.mutex.Lock()
		val, err := d.q.Peek()
		switch err {
		case nil:
			delay := val.Delay()
			if delay <= 0 {
				val, err = d.q.Dequeue()
				d.dequeueSignal.broadcast()
				return val, err
			}
			signal := d.dequeueSignal.signalCh()
			if timer == nil {
				timer = time.NewTimer(delay)
			} else {
				timer.Reset(delay)
			}
			select {
			case <-ctx.Done():
				var t T
				return t, ctx.Err()
			case <-timer.C:
				// 到时间了
				d.mutex.Lock()
				// 原队头可能被其他协程先出队了，检查一下队头
				val, err := d.q.Peek()
				if err != nil || val.Delay() > 0 {
					d.mutex.Unlock()
					continue
				}
				// 出队
				val, err = d.q.Dequeue()
				d.dequeueSignal.broadcast()
				return val, err
			case <-signal:
				// 进入下一循环
			}
		case queue.ErrEmptyQueue:
			signal := d.enqueueSignal.signalCh()
			select {
			case <-ctx.Done():
				var t T
				return t, ctx.Err()
			case <-signal:
			}
		default:
			d.mutex.Unlock()
			var t T
			return t, fmt.Errorf("generalization_tool:延时队列出队的适合遇到未知错误 %w,请上报", err)
		}
	}
}

type Delayable interface {
	Delay() time.Duration
}

type cond struct {
	signal chan struct{}
	l      sync.Locker
}

func newCond(l sync.Locker) *cond {
	return &cond{
		signal: make(chan struct{}),
		l:      l,
	}
}

// broadcast 广播，唤醒等待者如果没人等待，什么也不会发生。
// 必须加锁之后才能调用这个方法，广播之后释放锁，主要是为了确保用户必然是在锁范围内的。
func (c *cond) broadcast() {
	signal := make(chan struct{})
	old := c.signal
	c.signal = signal
	c.l.Unlock()
	close(old)
}

// signalCh 返回一个channel，用于监听广播信号，必须在锁范围内使用
// 调用后，锁会被释放，主要是为了确保用户必然是在锁范围内的
func (c *cond) signalCh() <-chan struct{} {
	res := c.signal
	c.l.Unlock()
	return res
}
