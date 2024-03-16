package slice

// DiffSet 求差集，只支持comparable，返回顺序不一定
func DiffSet[T comparable](src, dst []T) []T {
	res := make([]T, 0, len(src))
	srcMap := toMap[T](src)

	for _, v := range dst {
		delete(srcMap, v)
	}

	for v := range srcMap {
		res = append(res, v)
	}

	return res
}

// DiffSetFunc 差集，已去重
// 你应该优先使用 DiffSet
func DiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	// 优化容量估计
	res := make([]T, 0, len(src))
	for _, v := range src {
		if !ContainsFunc[T](dst, v, equal) {
			res = append(res, v)
		}
	}
	return deduplicateFunc[T](res, equal)
}
