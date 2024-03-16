package list

import (
	"errors"
	"fmt"
	"generalization_tool/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		index   int
		newVal  int
		wantRes []int
		wantErr error
	}{
		{
			name:    "add num to index left",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   0,
			newVal:  0,
			wantRes: []int{0, 1, 2, 3},
		},
		{
			name:    "add num to index right",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   3,
			newVal:  4,
			wantRes: []int{1, 2, 3, 4},
		},
		{
			name:    "add num to index mid",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   1,
			newVal:  4,
			wantRes: []int{1, 4, 2, 3},
		},
		{
			name:    "add num to index -1",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   -1,
			newVal:  4,
			wantErr: errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name:    "add num to index OutOfRange",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   4,
			newVal:  4,
			wantErr: errs.NewErrIndexOutOfRange(3, 4),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Add(tc.index, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
			// 因为返回了 error，所以我们不用继续往下比较了
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.values)
		})
	}
}

func TestArrayList_Cap(t *testing.T) {
	testCases := []struct {
		name    string
		wantCap int
		list    *ArrayList[int]
	}{
		{
			name:    "the same as actual",
			wantCap: 5,
			list: &ArrayList[int]{
				values: make([]int, 5),
			},
		},
		{
			name:    "nil",
			wantCap: 0,
			list: &ArrayList[int]{
				values: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.list.Cap()
			assert.Equal(t, tc.wantCap, res)
		})
	}
}

func BenchmarkArrayList_Cap(b *testing.B) {
	list := &ArrayList[int]{
		values: make([]int, 0),
	}
	b.Run("Cap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Cap()
		}
	})

	b.Run("Runtime Cap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = cap(list.values)
		}
	})
}

func TestArrayList_Append(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		newVal  []int
		wantRes []int
	}{
		{
			name:    "append non-empty values to non-empty list",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			newVal:  []int{4, 5, 6, 7},
			wantRes: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:    "append empty values to non-empty list",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			newVal:  []int{},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "append nil to non-empty list",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "append non-empty values to empty list",
			list:    NewArrayListOf[int]([]int{}),
			newVal:  []int{4, 5, 6, 7},
			wantRes: []int{4, 5, 6, 7},
		},
		{
			name:    "append empty values to empty list",
			list:    NewArrayListOf[int]([]int{}),
			newVal:  []int{},
			wantRes: []int{},
		},
		{
			name:    "append nil values to empty list",
			list:    NewArrayListOf[int]([]int{}),
			newVal:  nil,
			wantRes: []int{},
		},
		{
			name:    "append non-empty values to nil list",
			list:    NewArrayListOf[int](nil),
			newVal:  []int{4, 5, 6, 7},
			wantRes: []int{4, 5, 6, 7},
		},
		{
			name:    "append empty values to nil list",
			list:    NewArrayListOf[int](nil),
			newVal:  []int{},
			wantRes: []int{},
		},
		{
			name:    "append nil to nil list",
			list:    NewArrayListOf[int](nil),
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

func TestArrayList_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		index   int
		wantVal int
		wantRes []int
		wantErr error
	}{
		{
			name: "deleted",
			list: &ArrayList[int]{
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			index:   8,
			wantVal: 9,
			wantRes: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "index out of range",
			list: &ArrayList[int]{
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
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
			assert.Equal(t, tc.wantRes, tc.list.values)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestArrayList_Delete_Shrink(t *testing.T) {
	testCases := []struct {
		name    string
		preCap  int //切片的原始容量
		number  int //切片中元素的个数
		wantCap int //缩容后的容量
	}{
		//#----- #阶段一 逻辑测试 -----#
		//测试时：会默认删除一个元素，loop需要+1
		//阶段一：case1 - case5
		{
			name:    "case1：cap<=64,不进行缩容",
			preCap:  64,
			number:  2,
			wantCap: 64,
		},
		{
			name:    "case2：cap<=2048，len<=1/4cap，缩容至1/2",
			preCap:  2000,
			number:  401,
			wantCap: 1000,
		},
		{
			name:    "case3：cap<=2048，len>1/4cap，不满足条件，不进行缩容",
			preCap:  2000,
			number:  701,
			wantCap: 2000,
		},
		{
			name:    "case4：cap>2048，len<=1/2cap，缩容至5/8",
			preCap:  4000,
			number:  1801,
			wantCap: 2500,
		},
		{
			name:    "case5：cap>2048，len>1/2cap，不满足条件，不进行缩容",
			preCap:  4000,
			number:  2201,
			wantCap: 4000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewArrayList[int](tc.preCap)
			for i := 0; i < tc.number; i++ {
				_ = list.Append(i)
			}
			_, _ = list.Delete(0)
			assert.Equal(t, tc.wantCap, list.Cap())
		})
	}
}

func TestArrayList_Len(t *testing.T) {
	testCases := []struct {
		name    string
		wantLen int
		list    *ArrayList[int]
	}{
		{
			name:    "与实际元素数相等",
			wantLen: 5,
			list: &ArrayList[int]{
				values: make([]int, 5),
			},
		},
		{
			name:    "用户传入nil",
			wantLen: 0,
			list: &ArrayList[int]{
				values: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.list.Cap()
			assert.Equal(t, tc.wantLen, actual)
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		index   int
		wantVal int
		wantErr error
	}{
		{
			name:    "index 0",
			list:    NewArrayListOf[int]([]int{123, 100}),
			index:   0,
			wantVal: 123,
		},
		{
			name:    "index 2",
			list:    NewArrayListOf[int]([]int{123, 100}),
			index:   2,
			wantVal: 0,
			wantErr: errs.NewErrIndexOutOfRange(2, 2),
		},
		{
			name:    "index -1",
			list:    NewArrayListOf[int]([]int{123, 100}),
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

func TestArrayList_Range(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		wantVal int
		wantErr error
	}{
		{
			name: "计算全部元素的和",
			list: &ArrayList[int]{
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			wantVal: 55,
			wantErr: nil,
		},
		{
			name: "测试中断",
			list: &ArrayList[int]{
				values: []int{1, 2, 3, 4, -5, 6, 7, 8, -9, 10},
			},
			wantVal: 41,
			wantErr: errors.New("index 4 is error"),
		},
		{
			name: "测试数组为nil",
			list: &ArrayList[int]{
				values: nil,
			},
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

func TestArrayList_AsSlice(t *testing.T) {
	values := []int{1, 2, 3}
	a := NewArrayListOf[int](values)
	newSlice := a.AsSlice()
	// 内容相同
	assert.Equal(t, newSlice, values)
	Addr := fmt.Sprintf("%p", values)
	newAddr := fmt.Sprintf("%p", newSlice)
	// 但是地址不同，也就是意味着 newSlice 必须是一个新创建的
	assert.NotEqual(t, Addr, newAddr)
}

func TestArrayList_Set(t *testing.T) {
	testCases := []struct {
		name      string
		list      *ArrayList[int]
		index     int
		newVal    int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "set 15 by index 1",
			list:      NewArrayListOf[int]([]int{0, 1, 2, 3, 4}),
			index:     1,
			newVal:    15,
			wantSlice: []int{0, 15, 2, 3, 4},
			wantErr:   nil,
		},
		{
			name:      "index -1",
			list:      NewArrayListOf[int]([]int{0, 1, 2, 3, 4}),
			index:     -1,
			newVal:    5,
			wantSlice: []int{},
			wantErr:   errs.NewErrIndexOutOfRange(5, -1),
		},
		{
			name:      "index 10",
			list:      NewArrayListOf[int]([]int{0, 1, 2, 3, 4}),
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
			assert.Equal(t, tc.wantSlice, tc.list.values)
		})
	}
}
