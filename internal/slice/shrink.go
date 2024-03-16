package slice

// Shrink 缩容
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度小于容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func Shrink[T any](src []T) []T {
	capacity, length := cap(src), len(src)
	caps, change := calculateCap(capacity, length)
	if !change {
		return src
	}

	res := make([]T, 0, caps)
	res = append(res, src...)
	return res
}

func calculateCap(capacity int, length int) (int, bool) {
	if capacity <= 64 {
		return capacity, false
	}
	if capacity <= 2048 && 4*length <= capacity {
		capacity = capacity / 2
		return capacity, true
	}
	if capacity > 2048 && 2*length <= capacity {
		capacity = int(float32(0.625) * float32(capacity))
		return capacity, true
	}
	return capacity, false
}
