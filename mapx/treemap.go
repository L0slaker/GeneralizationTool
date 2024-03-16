package mapx

import (
	"errors"
	"generalization_tool"
	"generalization_tool/internal/tree"
)

var _ mapi[any, any] = (*TreeMap[any, any])(nil)

var errTreeMapComparatorIsNull = errors.New("TreeMap：Comparator不能为nil")

// TreeMap 基于红黑树实现的map
type TreeMap[K any, V any] struct {
	tree *tree.RBTree[K, V]
}

// NewTreeMapWithMap 支持通过传入的map构造生成TreeMap
func NewTreeMapWithMap[K comparable, V any](compare generalization_tool.Comparator[K], m map[K]V) (*TreeMap[K, V], error) {
	treeMap, err := NewTreeMap[K, V](compare)
	if err != nil {
		return treeMap, err
	}
	putAll(treeMap, m)
	return treeMap, nil
}

// NewTreeMap 创建一个的TreeMap,需注意比较器compare不能为nil
func NewTreeMap[K any, V any](compare generalization_tool.Comparator[K]) (*TreeMap[K, V], error) {
	if compare == nil {
		return nil, errTreeMapComparatorIsNull
	}
	return &TreeMap[K, V]{
		tree: tree.NewRBTree[K, V](compare),
	}, nil
}

// putAll 将map传入TreeMap，若map的key已存在，value将会被替换
func putAll[K comparable, V any](treeMap *TreeMap[K, V], m map[K]V) {
	for k, v := range m {
		_ = treeMap.Put(k, v)
	}
}

// Put 在TreeMap中插入指定值,若TreeMap已存在该key，value将会被替换
func (t *TreeMap[K, V]) Put(key K, value V) error {
	err := t.tree.Add(key, value)
	if err == tree.ErrRBTreeSameRBNode {
		return t.tree.Set(key, value)
	}
	return nil
}

// Get 在TreeMap找到指定key的节点，返回value；若未找到指定节点则会返回false
func (t *TreeMap[K, V]) Get(key K) (V, bool) {
	value, err := t.tree.Find(key)
	return value, err == nil
}

// Delete 删除TreeMap中指定key的节点
func (t *TreeMap[K, V]) Delete(key K) (V, bool) {
	return t.tree.Delete(key)
}

// Keys 返回全部的键（中序遍历）
func (t *TreeMap[K, V]) Keys() []K {
	keys, _ := t.tree.KeyValues()
	return keys
}

// Values 返回全部的值（中序遍历）
func (t *TreeMap[K, V]) Values() []V {
	_, values := t.tree.KeyValues()
	return values
}
