package slice

import (
	"generalization_tool/internal/errs"
	"generalization_tool/internal/slice"
)

// Delete 删除 index 处的元素
func Delete[T any](src []T, index int) (t []T, e error) {
	// 新建一个切片将原切片装进去,不指定容量
	if index < 0 || index >= len(src) {
		return t, errs.NewErrIndexOutOfRange(len(src), index)
	}

	for i := index; i < len(src)-1; i++ {
		src[i] = src[i+1]
	}

	src = slice.Shrink[T](src[:len(src)-1])

	return src, e
}

// FilterDelete 删除符合条件的元素
// 考虑到性能问题，所有操作都会在原切片上进行
// 被删除元素之后的元素会往前移动，有且只会移动一次
func FilterDelete[T any](src []T, m func(index int, src T) bool) []T {
	//记录被删除的元素位置，即空缺位置
	pos := 0
	for idx := range src {
		// 判断是否满足删除条件
		if m(idx, src[idx]) {
			continue
		}
		//移动元素
		src[pos] = src[idx]
		pos++
	}
	src = src[:pos]

	// 缩容
	src = slice.Shrink[T](src)

	return src
}
