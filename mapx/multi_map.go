package mapx

import "generalization_tool"

// MultiMap 多映射的map，可以将一个键映射到多个值上
type MultiMap[K any, V any] struct {
	m mapi[K, []V]
}

// NewMultiTreeMap 创建一个基于TreeMap的MultiMap。comparator不能为nil
func NewMultiTreeMap[K any, V any](comparator generalization_tool.Comparator[K]) (*MultiMap[K, V], error) {
	treeMap, err := NewTreeMap[K, []V](comparator)
	if err != nil {
		return nil, err
	}
	return &MultiMap[K, V]{
		m: treeMap,
	}, nil
}

// NewMultiHashMap 创建一个基于HashMap的MultiMap。comparator不能为nil
func NewMultiHashMap[K Hashable, V any](size int) *MultiMap[K, V] {
	var m mapi[K, []V] = NewHashMap[K, []V](size)
	return &MultiMap[K, V]{
		m: m,
	}
}

// NewMultiBuiltinMap 创建一个基于HashMap的MultiMap。comparator不能为nil
func NewMultiBuiltinMap[K comparable, V any](size int) *MultiMap[K, V] {
	var m mapi[K, []V] = newBuiltinMap[K, []V](size)
	return &MultiMap[K, V]{
		m: m,
	}
}

// Put 往MultiMap添加键值对或向已有key的值追加数据
func (m *MultiMap[K, V]) Put(key K, value V) error {
	return m.PutMany(key, value)
}

// PutMany 往MultiMap添加键值对或向已有key的值追加数据
func (m *MultiMap[K, V]) PutMany(key K, values ...V) error {
	val, _ := m.Get(key)
	val = append(val, values...)
	return m.m.Put(key, val)
}

// Get 从MultiMap获取已有key的值，若key不存在，返回的bool值为false
func (m *MultiMap[K, V]) Get(key K) ([]V, bool) {
	if val, ok := m.m.Get(key); ok {
		return append([]V{}, val...), ok
	}
	return nil, false
}

// Delete 从MultiMap中删除指定的key
func (m *MultiMap[K, V]) Delete(key K) ([]V, bool) {
	return m.m.Delete(key)
}

// Keys 获取MultiMap所有的key
func (m *MultiMap[K, V]) Keys() []K {
	return m.m.Keys()
}

// Values 获取MultiMap所有的value
func (m *MultiMap[K, V]) Values() [][]V {
	values := m.m.Values()
	copiedValues := make([][]V, 0, len(values))
	for k := range values {
		copiedValues = append(copiedValues, append([]V{}, values[k]...))
	}
	return copiedValues
}
