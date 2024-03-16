package list

import "sync"

var (
	_ List[any] = &ConcurrentList[any]{}
)

type ConcurrentList[T any] struct {
	List[T]
	lock sync.RWMutex
}

func NewConcurrentListOfSlice[T any](src []T) *ConcurrentList[T] {
	var list List[T] = NewArrayListOf(src)
	return &ConcurrentList[T]{List: list}
}

// Get 返回对应的下标元素
func (c *ConcurrentList[T]) Get(index int) (T, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Get(index)
}

// Append 在末尾追加元素
func (c *ConcurrentList[T]) Append(values ...T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Append(values...)
}

// Add 在指定位置添加新元素
func (c *ConcurrentList[T]) Add(index int, t T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Add(index, t)
}

// Set 更新 index 位置的值
func (c *ConcurrentList[T]) Set(index int, t T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Set(index, t)
}

// Delete 方法必要的时候引起缩容，其缩容规则是：
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func (c *ConcurrentList[T]) Delete(index int) (T, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Delete(index)
}

// Len 返回长度
func (c *ConcurrentList[T]) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Len()
}

// Cap 返回容量
func (c *ConcurrentList[T]) Cap() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Cap()
}

// Range 遍历 List 的所有元素
func (c *ConcurrentList[T]) Range(fn func(index int, t T) error) error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Range(fn)
}

// AsSlice 将 List转化为一个切片，没有元素的情况下不允许返回nil，
// 必须返回一个len和cap为0的切片；每次调用都必须都必须返回一个新切片
func (c *ConcurrentList[T]) AsSlice() []T {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.AsSlice()
}
