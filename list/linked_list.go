package list

import (
	"generalization_tool/internal/errs"
)

var (
	_ List[any] = &LinkedList[any]{}
)

type node[T any] struct {
	prev *node[T]
	next *node[T]
	val  T
}

// LinkedList 双向链表
type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewLinkedList[T any]() *LinkedList[T] {
	head := &node[T]{}
	tail := &node[T]{prev: head, next: head}
	head.prev, head.next = tail, tail
	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}

func NewLinkedListOf[T any](values []T) *LinkedList[T] {
	list := NewLinkedList[T]()
	if err := list.Append(values...); err != nil {
		panic(err)
	}
	return list
}

// Get 返回对应的下标元素
func (l *LinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.length {
		var zero T
		return zero, errs.NewErrIndexOutOfRange(l.length, index)
	}
	n := l.findNode(index)
	return n.val, nil
}

func (l *LinkedList[T]) findNode(index int) *node[T] {
	var res *node[T]
	if index <= l.Len()/2 {
		res = l.head
		for i := -1; i < index; i++ {
			res = res.next
		}
	} else {
		res = l.tail
		for i := l.Len(); i > index; i-- {
			res = res.prev
		}
	}
	return res
}

// Append 在末尾追加元素
func (l *LinkedList[T]) Append(values ...T) error {
	for _, v := range values {
		n := &node[T]{
			prev: l.tail.prev,
			next: l.tail,
			val:  v,
		}
		n.prev.next, n.next.prev = n, n
		l.length++
	}
	return nil
}

// Add 在指定位置添加新元素
func (l *LinkedList[T]) Add(index int, t T) error {
	if index < 0 || index > l.length {
		return errs.NewErrIndexOutOfRange(l.length, index)
	}
	if index == l.length {
		return l.Append(t)
	}
	dst := l.findNode(index)
	n := &node[T]{
		prev: dst.prev,
		next: dst,
		val:  t,
	}
	n.prev.next, n.next.prev = n, n
	//dst.prev.next,dst.prev = n,n
	l.length++
	return nil
}

// Set 更新 index 位置的值
func (l *LinkedList[T]) Set(index int, t T) error {
	if index < 0 || index >= l.length {
		return errs.NewErrIndexOutOfRange(l.length, index)
	}
	n := l.findNode(index)
	n.val = t
	return nil
}

// Delete 删除指定位置的元素
func (l *LinkedList[T]) Delete(index int) (T, error) {
	if index < 0 || index >= l.length {
		var zero T
		return zero, errs.NewErrIndexOutOfRange(l.length, index)
	}
	n := l.findNode(index)
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev, n.next = nil, nil
	l.length--
	return n.val, nil
}

// Len 返回长度
func (l *LinkedList[T]) Len() int {
	return l.length
}

// Cap 返回容量
func (l *LinkedList[T]) Cap() int {
	return l.Len()
}

// Range 遍历 List 的所有元素
func (l *LinkedList[T]) Range(fn func(index int, t T) error) error {
	res := l.head.next
	for i := 0; i < l.length; i++ {
		if err := fn(i, res.val); err != nil {
			return err
		}
		res = res.next
	}
	return nil
}

// AsSlice 将 List转化为一个切片，没有元素的情况下不允许返回nil，
// 必须返回一个len和cap为0的切片；每次调用都必须都必须返回一个新切片
func (l *LinkedList[T]) AsSlice() []T {
	newSlice := make([]T, l.length)
	res := l.head.next
	for i := 0; i < l.length; i++ {
		newSlice[i] = res.val
		res = res.next
	}
	return newSlice
}
