package list

import (
	"generalization_tool/internal/errs"
	"generalization_tool/internal/slice"
)

var (
	_ List[any] = &ArrayList[any]{}
)

type ArrayList[T any] struct {
	values []T
}

func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{values: make([]T, 0, cap)}
}

func NewArrayListOf[T any](values []T) *ArrayList[T] {
	return &ArrayList[T]{values: values}
}

// Get 返回对应的下标元素
func (a *ArrayList[T]) Get(index int) (t T, e error) {
	length := a.Len()
	if index < 0 || index >= length {
		return t, errs.NewErrIndexOutOfRange(length, index)
	}
	return a.values[index], e
}

// Append 在末尾追加元素
func (a *ArrayList[T]) Append(values ...T) error {
	a.values = append(a.values, values...)
	return nil
}

// Add 在指定位置添加新元素
func (a *ArrayList[T]) Add(index int, t T) error {
	length := a.Len()
	if index < 0 || index > length {
		return errs.NewErrIndexOutOfRange(length, index)
	}
	a.values = append(a.values, t)
	copy(a.values[index+1:], a.values[index:])
	a.values[index] = t
	return nil
}

// Set 更新 index 位置的值
func (a *ArrayList[T]) Set(index int, t T) error {
	length := a.Len()
	if index < 0 || index > length {
		return errs.NewErrIndexOutOfRange(length, index)
	}
	a.values[index] = t
	return nil
}

// Delete 方法必要的时候引起缩容，其缩容规则是：
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func (a *ArrayList[T]) Delete(index int) (T, error) {
	res, t, err := slice.Delete[T](a.values, index)
	if err != nil {
		return t, err
	}
	a.values = res
	a.shrink()
	return t, nil
}

// Len 返回长度
func (a *ArrayList[T]) Len() int {
	return len(a.values)
}

// Cap 返回容量
func (a *ArrayList[T]) Cap() int {
	return cap(a.values)
}

// Range 遍历 List 的所有元素
func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for k, v := range a.values {
		err := fn(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// AsSlice 将 List转化为一个切片，没有元素的情况下不允许返回nil，
// 必须返回一个len和cap为0的切片；每次调用都必须都必须返回一个新切片
func (a *ArrayList[T]) AsSlice() []T {
	res := make([]T, len(a.values))
	copy(res, a.values)
	return res
}

func (a *ArrayList[T]) shrink() {
	a.values = slice.Shrink[T](a.values)
}
