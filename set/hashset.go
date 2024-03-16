package setx

import "generalization_tool/mapx"

type HashSet[T mapx.Hashable] struct {
	hashMap *mapx.HashMap[T, any]
}

func NewHashSet[T mapx.Hashable](size int) *HashSet[T] {
	hashMap := mapx.NewHashMap[T, any](size)
	return &HashSet[T]{
		hashMap: hashMap,
	}
}

func (h *HashSet[T]) Add(key T) {
	_ = h.hashMap.Put(key, nil)
}

func (h *HashSet[T]) Delete(key T) {
	h.hashMap.Delete(key)
}

func (h *HashSet[T]) Exist(key T) bool {
	_, isExist := h.hashMap.Get(key)
	return isExist
}

func (h *HashSet[T]) Keys() []T {
	return h.hashMap.Keys()
}
