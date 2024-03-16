package list

import (
	"errors"
	"fmt"
	"generalization_tool/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcurrent_Add(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ConcurrentList[int]
		index   int
		newVal  int
		wantRes []int
		wantErr error
	}{
		{
			name:    "add num to index left",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3}),
			index:   0,
			newVal:  0,
			wantRes: []int{0, 1, 2, 3},
		},
		{
			name:    "add num to index right",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3}),
			index:   3,
			newVal:  4,
			wantRes: []int{1, 2, 3, 4},
		},
		{
			name:    "add num to index mid",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3}),
			index:   1,
			newVal:  4,
			wantRes: []int{1, 4, 2, 3},
		},
		{
			name:    "add num to index -1",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3}),
			index:   -1,
			newVal:  4,
			wantErr: errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name:    "add num to index OutOfRange",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3}),
			index:   4,
			newVal:  4,
			wantErr: errs.NewErrIndexOutOfRange(3, 4),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Add(tc.index, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.AsSlice())
		})
	}
}

func TestConcurrent_Cap(t *testing.T) {
	testCases := []struct {
		name    string
		wantCap int
		list    *ConcurrentList[int]
	}{
		{
			name:    "the same as actual",
			wantCap: 5,
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3, 4, 5}),
		},
		{
			name:    "nil",
			wantCap: 0,
			list:    NewConcurrentListOfSlice[int](nil),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.list.Cap()
			assert.Equal(t, tc.wantCap, res)
		})
	}
}

func TestConcurrent_Append(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ConcurrentList[int]
		newVal  []int
		wantRes []int
	}{
		{
			name:    "append non-empty values to non-empty list",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3}),
			newVal:  []int{4, 5, 6, 7},
			wantRes: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:    "append empty values to non-empty list",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3}),
			newVal:  []int{},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "append nil to non-empty list",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3}),
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "append non-empty values to empty list",
			list:    NewConcurrentListOfSlice[int]([]int{}),
			newVal:  []int{4, 5, 6, 7},
			wantRes: []int{4, 5, 6, 7},
		},
		{
			name:    "append empty values to empty list",
			list:    NewConcurrentListOfSlice[int]([]int{}),
			newVal:  []int{},
			wantRes: []int{},
		},
		{
			name:    "append nil values to empty list",
			list:    NewConcurrentListOfSlice[int]([]int{}),
			newVal:  nil,
			wantRes: []int{},
		},
		{
			name:    "append non-empty values to nil list",
			list:    NewConcurrentListOfSlice[int](nil),
			newVal:  []int{4, 5, 6, 7},
			wantRes: []int{4, 5, 6, 7},
		},
		{
			name:    "append empty values to nil list",
			list:    NewConcurrentListOfSlice[int](nil),
			newVal:  []int{},
			wantRes: []int{},
		},
		{
			name:    "append nil to nil list",
			list:    NewConcurrentListOfSlice[int](nil),
			newVal:  nil,
			wantRes: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Append(tc.newVal...)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.AsSlice())
		})
	}
}

func TestConcurrent_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ConcurrentList[int]
		index   int
		wantVal int
		wantRes []int
		wantErr error
	}{
		{
			name:    "deleted",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			index:   8,
			wantVal: 9,
			wantRes: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:    "index out of range",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			index:   10,
			wantErr: errs.NewErrIndexOutOfRange(9, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.list.Delete(tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.AsSlice())
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestConcurrent_Len(t *testing.T) {
	testCases := []struct {
		name    string
		wantLen int
		list    *ConcurrentList[int]
	}{
		{
			name:    "与实际元素数相等",
			wantLen: 5,
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3, 4, 5}),
		},
		{
			name:    "用户传入nil",
			wantLen: 0,
			list:    NewConcurrentListOfSlice[int]([]int{}),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.list.Cap()
			assert.Equal(t, tc.wantLen, actual)
		})
	}
}

func TestConcurrent_Get(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ConcurrentList[int]
		index   int
		wantVal int
		wantErr error
	}{
		{
			name:    "index 0",
			list:    NewConcurrentListOfSlice[int]([]int{123, 100}),
			index:   0,
			wantVal: 123,
		},
		{
			name:    "index 2",
			list:    NewConcurrentListOfSlice[int]([]int{123, 100}),
			index:   2,
			wantVal: 0,
			wantErr: errs.NewErrIndexOutOfRange(2, 2),
		},
		{
			name:    "index -1",
			list:    NewConcurrentListOfSlice[int]([]int{123, 100}),
			index:   -1,
			wantVal: 0,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.list.Get(tc.index)
			assert.Equal(t, tc.wantErr, err)
			// 因为返回了 error，所以我们不用继续往下比较了
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestConcurrent_Range(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ConcurrentList[int]
		wantVal int
		wantErr error
	}{
		{
			name:    "计算全部元素的和",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			wantVal: 55,
			wantErr: nil,
		},
		{
			name:    "测试中断",
			list:    NewConcurrentListOfSlice[int]([]int{1, 2, 3, 4, -5, 6, 7, 8, -9, 10}),
			wantVal: 41,
			wantErr: errors.New("index 4 is error"),
		},
		{
			name:    "测试数组为nil",
			list:    NewConcurrentListOfSlice[int](nil),
			wantVal: 0,
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := 0
			err := tc.list.Range(func(index int, num int) error {
				if num < 0 {
					return fmt.Errorf("index %d is error", index)
				}
				result += num
				return nil
			})

			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, result)
		})
	}
}

func TestConcurrent_AsSlice(t *testing.T) {
	values := []int{1, 2, 3}
	a := NewConcurrentListOfSlice[int](values)
	newSlice := a.AsSlice()
	// 内容相同
	assert.Equal(t, newSlice, values)
	Addr := fmt.Sprintf("%p", values)
	newAddr := fmt.Sprintf("%p", newSlice)
	// 但是地址不同，也就是意味着 newSlice 必须是一个新创建的
	assert.NotEqual(t, Addr, newAddr)
}

func TestConcurrent_Set(t *testing.T) {
	testCases := []struct {
		name      string
		list      *ConcurrentList[int]
		index     int
		newVal    int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "set 15 by index 1",
			list:      NewConcurrentListOfSlice[int]([]int{0, 1, 2, 3, 4}),
			index:     1,
			newVal:    15,
			wantSlice: []int{0, 15, 2, 3, 4},
			wantErr:   nil,
		},
		{
			name:      "index -1",
			list:      NewConcurrentListOfSlice[int]([]int{0, 1, 2, 3, 4}),
			index:     -1,
			newVal:    5,
			wantSlice: []int{},
			wantErr:   errs.NewErrIndexOutOfRange(5, -1),
		},
		{
			name:      "index 10",
			list:      NewConcurrentListOfSlice[int]([]int{0, 1, 2, 3, 4}),
			index:     10,
			newVal:    6,
			wantSlice: []int{},
			wantErr:   errs.NewErrIndexOutOfRange(5, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Set(tc.index, tc.newVal)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantSlice, tc.list.AsSlice())
		})
	}
}
