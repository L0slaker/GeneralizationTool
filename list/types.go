package list

// List 该接口定义各个方法的行为和表现
type List[T any] interface {
	// Get 返回对应的下标元素
	Get(index int) (T, error)
	// Append 在末尾追加元素
	Append(values ...T) error
	// Add 在指定位置添加新元素
	Add(index int, t T) error
	// Set 更新 index 位置的值
	Set(index int, t T) error
	// Delete 删除目标元素的位置，并返回该位置的值
	Delete(index int) (T, error)
	// Len 返回长度
	Len() int
	// Cap 返回容量
	Cap() int
	// Range 遍历 List 的所有元素
	Range(fn func(index int, t T) error) error
	// AsSlice 将 List转化为一个切片，没有元素的情况下不允许返回nil，
	//必须返回一个len和cap为0的切片；每次调用都必须都必须返回一个新切片
	AsSlice() []T
}
