package slice

import "github.com/L0slaker/GeneralizationTool/internal/errs"

// Delete 区别于切片中的Delete方法，多返回了删除的参数
func Delete[T any](src []T, index int) ([]T, T, error) {
	if index < 0 || index >= len(src) {
		var t T
		return nil, t, errs.NewErrIndexOutOfRange(len(src), index)
	}

	res := src[index]
	for i := index; i < len(src)-1; i++ {
		src[i] = src[i+1]
	}

	return src[:len(src)-1], res, nil
}
