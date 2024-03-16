package errs

import "fmt"

func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("generalization_tool: 下表超出范围，长度 %d, 下标 %d", length, index)
}
