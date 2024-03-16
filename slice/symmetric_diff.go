package slice

// SymmetricDiffSet 对称差集
// 已去重
// 返回值的元素顺序是不定的
func SymmetricDiffSet[T comparable](src, dst []T) []T {
	res := make([]T, 0, len(src))
	srcMap := toMap[T](src)
	dstMap := toMap[T](dst)

	for k := range srcMap {
		if _, exist := dstMap[k]; exist {
			// 删除相同元素
			delete(srcMap, k)
			delete(dstMap, k)
		}
	}

	//将两个切片的值添加到一起
	for k, v := range srcMap {
		dstMap[k] = v
	}

	//装入结果集中
	for k := range dstMap {
		res = append(res, k)
	}
	return res
}

// SymmetricDiffSetFunc 对称差集
// 你应该优先使用 SymmetricDiffSet
// 已去重
func SymmetricDiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	intersect := make([]T, 0, min(len(src), len(dst)))

	//计算得到交集
	for _, valSrc := range src {
		for _, valDst := range dst {
			if equal(valSrc, valDst) {
				intersect = append(intersect, valSrc)
			}
		}
	}

	res := make([]T, 0, len(src)+len(dst)-len(intersect)*2)
	//取src与交集的差集，存入结果集中
	for _, v := range src {
		if !ContainsFunc[T](intersect, v, equal) {
			res = append(res, v)
		}
	}

	//取dst与交集的差集，存入结果集中
	for _, v := range dst {
		if !ContainsFunc[T](intersect, v, equal) {
			res = append(res, v)
		}
	}

	return deduplicateFunc[T](res, equal)
}

func min(src, dst int) int {
	if src > dst {
		return dst
	}
	return src
}
