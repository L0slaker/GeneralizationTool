package setx

type Set[T comparable] interface {
	Add(key T)
	Delete(key T)
	Exist(key T) bool
	Keys() []T
}

type MapSet[T comparable] struct {
	m map[T]struct{}
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func (s *MapSet[T]) Add(key T) {
	s.m[key] = struct{}{}
}

func (s *MapSet[T]) Delete(key T) {
	delete(s.m, key)
}

func (s *MapSet[T]) Exist(key T) bool {
	_, exist := s.m[key]
	return exist
}

func (s *MapSet[T]) Keys() []T {
	res := make([]T, 0, len(s.m))
	for k := range s.m {
		res = append(res, k)
	}
	return res
}
