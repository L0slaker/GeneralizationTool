package slice

// FilterMap 执行过滤并且转化
// 如果 m 的第二个返回值是 false，那么我们会忽略第一个返回值
// 即便第二个返回值是 false，后续的元素依旧会被遍历
func FilterMap[Src any, Dst any](src []Src, m func(idx int, src Src) (Dst, bool)) []Dst {
	res := make([]Dst, 0, len(src))
	for i, s := range src {
		if dst, ok := m(i, s); ok {
			res = append(res, dst)
		}
	}
	return res
}

func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	res := make([]Dst, len(src))
	for i, s := range src {
		res[i] = m(i, s)
	}
	return res
}

// toMap 构造map
func toMap[T comparable](src []T) map[T]struct{} {
	m := make(map[T]struct{})

	for _, v := range src {
		m[v] = struct{}{}
	}
	return m
}

// 去重
func deduplicateFunc[T any](data []T, equal equalFunc[T]) []T {
	var newData = make([]T, 0, len(data))
	for k, v := range data {
		if !ContainsFunc[T](data[k+1:], v, equal) {
			newData = append(newData, v)
		}
	}
	return newData
}

func deduplicate[T comparable](data []T) []T {
	dataMap := toMap[T](data)
	var newData = make([]T, 0, len(dataMap))
	for key := range dataMap {
		newData = append(newData, key)
	}
	return newData
}
