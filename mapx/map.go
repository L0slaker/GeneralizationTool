package mapx

// Keys 返回 map 里面所有的key。key的顺序是随机的
func Keys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

// Values 返回 map 里面所有的values。values的顺序是随机的
func Values[K comparable, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))
	for k := range m {
		res = append(res, m[k])
	}
	return res
}

// KeysValues 返回 map 里面所有的 key，values
func KeysValues[K comparable, V any](m map[K]V) ([]K, []V) {
	KeyRes := make([]K, 0, len(m))
	ValRes := make([]V, 0, len(m))
	for k := range m {
		KeyRes = append(KeyRes, k)
		ValRes = append(ValRes, m[k])
	}
	return KeyRes, ValRes
}
