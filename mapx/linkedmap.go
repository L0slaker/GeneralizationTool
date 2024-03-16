package mapx

import "generalization_tool"

type linkedKeyValue[K any, V any] struct {
	key   K
	value V
	prev  *linkedKeyValue[K, V]
	next  *linkedKeyValue[K, V]
}

type LinkedMap[K any, V any] struct {
	m      mapi[K, *linkedKeyValue[K, V]]
	head   *linkedKeyValue[K, V]
	tail   *linkedKeyValue[K, V]
	length int
}

func NewLinkedHashMap[K Hashable, V any](size int) *LinkedMap[K, V] {
	hashmap := NewHashMap[K, *linkedKeyValue[K, V]](size)
	head := &linkedKeyValue[K, V]{}
	tail := &linkedKeyValue[K, V]{prev: head, next: head}
	head.prev, head.next = tail, tail
	return &LinkedMap[K, V]{
		m:    hashmap,
		head: head,
		tail: tail,
	}
}

func NewLinkedTreeMap[K any, V any](comparator generalization_tool.Comparator[K]) (*LinkedMap[K, V], error) {
	treeMap, err := NewTreeMap[K, *linkedKeyValue[K, V]](comparator)
	if err != nil {
		return nil, err
	}
	head := &linkedKeyValue[K, V]{}
	tail := &linkedKeyValue[K, V]{prev: head, next: head}
	head.prev, head.next = tail, tail
	return &LinkedMap[K, V]{
		m:    treeMap,
		head: head,
		tail: tail,
	}, nil
}

func (l *LinkedMap[K, V]) Put(key K, value V) error {
	if lkMap, ok := l.m.Get(key); ok {
		lkMap.value = value
		return nil
	}
	lkMap := &linkedKeyValue[K, V]{
		key:   key,
		value: value,
		prev:  l.tail.prev,
		next:  l.tail,
	}
	if err := l.m.Put(key, lkMap); err != nil {
		return err
	}
	lkMap.prev.next, lkMap.next.prev = lkMap, lkMap
	l.length++
	return nil
}

func (l *LinkedMap[K, V]) Get(key K) (V, bool) {
	if lkMap, ok := l.m.Get(key); ok {
		return lkMap.value, ok
	}
	var zero V
	return zero, false
}

func (l *LinkedMap[K, V]) Delete(key K) (V, bool) {
	if lkMap, ok := l.m.Delete(key); ok {
		lkMap.prev.next, lkMap.next.prev = lkMap.next, lkMap.prev
		return lkMap.value, ok
	}
	var zero V
	return zero, false
}

func (l *LinkedMap[K, V]) Keys() []K {
	res := make([]K, 0, l.length)
	for current := l.head.next; current != l.tail; {
		res = append(res, current.key)
		current = current.next
	}
	return res
}

func (l *LinkedMap[K, V]) Values() []V {
	res := make([]V, 0, l.length)
	for current := l.head.next; current != l.tail; {
		res = append(res, current.value)
		current = current.next
	}
	return res
}
