package GeneralizationTool

// Comparator 用于比较两个对象的大小
type Comparator[T any] func(src T, dst T) int

func ComparatorRealNumber[T RealNumber](src T, dst T) int {
	if src < dst {
		return -1
	} else if src == dst {
		return 0
	} else {
		return 1
	}
}
