package mapx

type mapi[K any, V any] interface {
	Put(key K, value V) error
	Get(key K) (V, bool)
	Delete(key K) (V, bool)
	// Keys 返回所有的键，调用多次拿到的结果不一定相等
	Keys() []K
	// Values 返回所有的值，调用多次拿到的结果不一定相等
	Values() []V
}
