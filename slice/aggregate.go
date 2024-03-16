package slice

import "github.com/L0slaker/GeneralizationTool"

// Max 返回最大值。
// 该方法假设你至少会传入一个值。
// 在使用 float32 或者 float64 的时候要小心精度问题
func Max[T GeneralizationTool.RealNumber](val []T) T {
	res := val[0]
	for i := 1; i < len(val); i++ {
		if val[i] > res {
			res = val[i]
		}
	}
	return res
}

// Min 返回最小值
// 该方法会假设你至少会传入一个值
// 在使用 float32 或者 float64 的时候要小心精度问题
func Min[T GeneralizationTool.RealNumber](val []T) T {
	res := val[0]
	for i := 1; i < len(val); i++ {
		if val[i] < res {
			res = val[i]
		}
	}
	return res
}

// Sum 求和
// 在使用 float32 或者 float64 的时候要小心精度问题
func Sum[T GeneralizationTool.RealNumber](val []T) T {
	var res T
	for _, v := range val {
		res += v
	}
	return res
}
