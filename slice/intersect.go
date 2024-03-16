package slice

// Intersect 求交集
func Intersect[T comparable](src, dst []T) []T {
	srcMap := toMap(src)
	res := make([]T, 0, len(src))

	for _, v := range dst {
		if _, ok := srcMap[v]; ok {
			delete(srcMap, v)
			res = append(res, v)
		}
	}

	return deduplicate[T](res)
}

// IntersectSetFunc 支持任意类型
// 你应该优先使用 Intersect
// 已去重
func IntersectSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src))
	for _, v1 := range src {
		for _, v2 := range dst {
			if equal(v1, v2) {
				res = append(res, v2)
				break
			}
		}
	}
	return deduplicateFunc[T](res, equal)
}
