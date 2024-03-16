package setx

import (
	"generalization_tool"
	"generalization_tool/mapx"
)

var _ Set[int] = (*TreeSet[int])(nil)

type TreeSet[T any] struct {
	treeMap *mapx.TreeMap[T, any]
}

func NewTreeSet[T any](compare generalization_tool.Comparator[T]) (*TreeSet[T], error) {
	treeMap, err := mapx.NewTreeMap[T, any](compare)
	if err != nil {
		return nil, err
	}
	return &TreeSet[T]{
		treeMap: treeMap,
	}, nil
}

func (s *TreeSet[T]) Add(key T) {
	_ = s.treeMap.Put(key, nil)
}

func (s *TreeSet[T]) Delete(key T) {
	s.treeMap.Delete(key)
}

func (s *TreeSet[T]) Exist(key T) bool {
	_, isExist := s.treeMap.Get(key)
	return isExist
}

func (s *TreeSet[T]) Keys() []T {
	return s.treeMap.Keys()
}
