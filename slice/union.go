package slice

// UnionSet 求并集，只支持comparable，返回顺序不一定
func UnionSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)
	dstMap := toMap[T](dst)

	for k := range srcMap {
		dstMap[k] = struct{}{}
	}

	res := make([]T, 0, len(dstMap))
	for k := range dstMap {
		res = append(res, k)
	}

	return res
}

// UnionSetFunc 并集，支持任意类型
// 你应该优先使用 UnionSet
// 已去重
func UnionSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src)+len(dst))
	res = append(res, src...)
	res = append(res, dst...)

	return deduplicateFunc[T](res, equal)
}
