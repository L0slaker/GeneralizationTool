package list

import (
	"errors"
	"fmt"
	"generalization_tool/internal/errs"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	testCases := []struct {
		name     string
		list     *LinkedList[int]
		index    int
		newVal   int
		wantList *LinkedList[int]
		wantErr  error
	}{
		{
			name:     "add num to index left",
			list:     NewLinkedListOf[int]([]int{1, 2, 3}),
			index:    0,
			newVal:   0,
			wantList: NewLinkedListOf[int]([]int{0, 1, 2, 3}),
		},
		{
			name:     "add num to index right",
			list:     NewLinkedListOf[int]([]int{1, 2, 3}),
			index:    3,
			newVal:   4,
			wantList: NewLinkedListOf[int]([]int{1, 2, 3, 4}),
		},
		{
			name:     "add num to index mid",
			list:     NewLinkedListOf[int]([]int{1, 2, 3}),
			index:    1,
			newVal:   4,
			wantList: NewLinkedListOf[int]([]int{1, 4, 2, 3}),
		},
		{
			name:    "add num to index -1",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   -1,
			newVal:  4,
			wantErr: errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name:    "add num to index OutOfRange",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
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
			assert.Equal(t, tc.wantList.AsSlice(), tc.list.AsSlice())
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	testCases := []struct {
		name     string
		list     *LinkedList[int]
		index    int
		delVal   int
		wantList *LinkedList[int]
		wantErr  error
	}{
		{
			name:    "deleted num to index -1",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(9, -1),
		},
		{
			name:    "deleted beyond length index 99",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			index:   99,
			wantErr: errs.NewErrIndexOutOfRange(9, 99),
		},
		{
			name:    "deleted beyond length index 9",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			index:   9,
			wantErr: errs.NewErrIndexOutOfRange(9, 9),
		},
		{
			name:    "delete empty node",
			list:    NewLinkedListOf[int]([]int{}),
			index:   0,
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
		},
		{
			name:     "delete num to index 0",
			list:     NewLinkedListOf[int]([]int{1, 2, 3}),
			index:    0,
			delVal:   1,
			wantList: NewLinkedListOf[int]([]int{2, 3}),
		},
		{
			name:     "delete num to index by tail",
			list:     NewLinkedListOf[int]([]int{1, 2, 3, 4, 5, 6, 7}),
			index:    6,
			delVal:   7,
			wantList: NewLinkedListOf[int]([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:     "delete num to index 1",
			list:     NewLinkedListOf[int]([]int{1, 2, 3, 4, 5, 6, 7}),
			index:    1,
			delVal:   2,
			wantList: NewLinkedListOf[int]([]int{1, 3, 4, 5, 6, 7}),
		},
		{
			name:     "deleting an element with only one",
			list:     NewLinkedListOf[int]([]int{100}),
			index:    0,
			delVal:   100,
			wantList: NewLinkedListOf[int]([]int{}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.list.Delete(tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantList.AsSlice(), tc.list.AsSlice())
			assert.Equal(t, tc.delVal, val)
		})
	}
}

func TestLinkedList_Append(t *testing.T) {
	testCases := []struct {
		name     string
		list     *LinkedList[int]
		newVal   []int
		wantList *LinkedList[int]
		wantErr  error
	}{
		{
			name:     "append non-empty values to non-empty list",
			list:     NewLinkedListOf[int]([]int{1, 2, 3}),
			newVal:   []int{4, 5, 6, 7},
			wantList: NewLinkedListOf[int]([]int{1, 2, 3, 4, 5, 6, 7}),
		},
		{
			name:     "append empty values to non-empty list",
			list:     NewLinkedListOf[int]([]int{1, 2, 3}),
			newVal:   []int{},
			wantList: NewLinkedListOf[int]([]int{1, 2, 3}),
		},
		{
			name:     "append nil to non-empty list",
			list:     NewLinkedListOf[int]([]int{1, 2, 3}),
			newVal:   nil,
			wantList: NewLinkedListOf[int]([]int{1, 2, 3}),
		},
		{
			name:     "append non-empty values to empty list",
			list:     NewLinkedListOf[int]([]int{}),
			newVal:   []int{4, 5, 6, 7},
			wantList: NewLinkedListOf[int]([]int{4, 5, 6, 7}),
		},
		{
			name:     "append empty values to empty list",
			list:     NewLinkedListOf[int]([]int{}),
			newVal:   []int{},
			wantList: NewLinkedListOf[int]([]int{}),
		},
		{
			name:     "append nil values to empty list",
			list:     NewLinkedListOf[int]([]int{}),
			newVal:   nil,
			wantList: NewLinkedListOf[int]([]int{}),
		},
		{
			name:     "append non-empty values to nil list",
			list:     NewLinkedListOf[int](nil),
			newVal:   []int{4, 5, 6, 7},
			wantList: NewLinkedListOf[int]([]int{4, 5, 6, 7}),
		},
		{
			name:     "append empty values to nil list",
			list:     NewLinkedListOf[int](nil),
			newVal:   []int{},
			wantList: NewLinkedListOf[int]([]int{}),
		},
		{
			name:     "append nil to nil list",
			list:     NewLinkedListOf[int](nil),
			newVal:   nil,
			wantList: NewLinkedListOf[int]([]int{}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Append(tc.newVal...)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantList.AsSlice(), tc.list.AsSlice())
		})
	}
}

func TestNewLinkedListOf(t *testing.T) {
	testCases := []struct {
		name        string
		slice       []int
		wantedSlice []int
	}{
		{
			name:        "nil",
			slice:       nil,
			wantedSlice: []int{},
		},
		{
			name:        "vacant",
			slice:       []int{},
			wantedSlice: []int{},
		},
		{
			name:        "single",
			slice:       []int{1},
			wantedSlice: []int{1},
		},
		{
			name:        "normal",
			slice:       []int{1, 2, 3},
			wantedSlice: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		list := NewLinkedListOf(tc.slice)
		assert.Equal(t, tc.wantedSlice, list.AsSlice())
	}
}

func TestLinkedList_AsSlice(t *testing.T) {
	values := []int{1, 2, 3}
	a := NewLinkedListOf[int](values)
	newSlice := a.AsSlice()
	// 内容相同
	assert.Equal(t, newSlice, values)
	Addr := fmt.Sprintf("%p", values)
	newAddr := fmt.Sprintf("%p", newSlice)
	// 但是地址不同，也就是意味着 newSlice 必须是一个新创建的
	assert.NotEqual(t, Addr, newAddr)
}

func TestLinkedList_Cap(t *testing.T) {
	list := NewLinkedList[int]()
	assert.Equal(t, 0, list.Cap())
	if err := list.Append(1); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, list.Cap())
}

func TestLinkedList_Get(t *testing.T) {
	testCases := []struct {
		name    string
		list    *LinkedList[int]
		index   int
		wantVal int
		wantErr error
	}{
		{
			name:    "get left",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, 5}),
			index:   0,
			wantVal: 1,
		},
		{
			name:    "get right",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, 5}),
			index:   4,
			wantVal: 5,
		},
		{
			name:    "get middle",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, 5}),
			index:   2,
			wantVal: 3,
		},
		{
			name:    "over left",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, 5}),
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(5, -1),
		},
		{
			name:    "over right",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, 5}),
			index:   5,
			wantErr: errs.NewErrIndexOutOfRange(5, 5),
		},
		{
			name:    "empty list",
			list:    NewLinkedListOf[int]([]int{}),
			index:   0,
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
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

func TestLinkedList_Range(t *testing.T) {
	testCases := []struct {
		name    string
		list    *LinkedList[int]
		wantVal int
		wantErr error
	}{
		{
			name:    "计算全部元素的和",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			wantVal: 55,
			wantErr: nil,
		},
		{
			name:    "测试中断",
			list:    NewLinkedListOf[int]([]int{1, 2, 3, 4, -5, 6, 7, 8, -9, 10}),
			wantVal: 41,
			wantErr: errors.New("index 4 is error"),
		},
		{
			name:    "测试数组为nil",
			list:    NewLinkedListOf[int](nil),
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

func TestLinkedList_Set(t *testing.T) {
	testCases := []struct {
		name     string
		list     *LinkedList[int]
		index    int
		newVal   int
		wantList *LinkedList[int]
		wantErr  error
	}{
		{
			name:    "set num to index -1",
			list:    NewLinkedListOf[int]([]int{0, 1, 2, 3, 4}),
			index:   -1,
			newVal:  15,
			wantErr: errs.NewErrIndexOutOfRange(5, -1),
		},
		{
			name:    "set beyond length index 99",
			list:    NewLinkedListOf[int]([]int{0, 1, 2, 3, 4}),
			index:   99,
			newVal:  15,
			wantErr: errs.NewErrIndexOutOfRange(5, 99),
		},
		{
			name:    "set empty node",
			list:    NewLinkedListOf[int]([]int{}),
			index:   1,
			newVal:  15,
			wantErr: errs.NewErrIndexOutOfRange(0, 1),
		},
		{
			name:     "set num to index 3",
			list:     NewLinkedListOf[int]([]int{0, 1, 2, 3, 4}),
			index:    2,
			newVal:   15,
			wantList: NewLinkedListOf[int]([]int{0, 1, 15, 3, 4}),
		},
		{
			name:     "set num to head",
			list:     NewLinkedListOf[int]([]int{0, 1, 2, 3, 4}),
			index:    0,
			newVal:   15,
			wantList: NewLinkedListOf[int]([]int{15, 1, 2, 3, 4}),
		},
		{
			name:     "set num to tail",
			list:     NewLinkedListOf[int]([]int{0, 1, 2, 3, 4}),
			index:    4,
			newVal:   15,
			wantList: NewLinkedListOf[int]([]int{0, 1, 2, 3, 15}),
		},
		{
			name:    "index == len(*node)",
			list:    NewLinkedListOf[int]([]int{0, 1, 2, 3, 4}),
			index:   5,
			newVal:  15,
			wantErr: errs.NewErrIndexOutOfRange(5, 5),
		},
		{
			name:    "len(*node) == 0",
			list:    NewLinkedListOf[int]([]int{}),
			index:   0,
			newVal:  15,
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Set(tc.index, tc.newVal)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantList.AsSlice(), tc.list.AsSlice())
		})
	}
}

func BenchmarkLinkedList_Add(b *testing.B) {
	l := NewLinkedListOf[int]([]int{1, 2, 3})
	testCase := make([]int, 0, b.N)
	for i := 1; i <= b.N; i++ {
		testCase = append(testCase, rand.Intn(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = l.Add(testCase[i], testCase[i])
	}
}

func BenchmarkLinkedList_Get(b *testing.B) {
	l := NewLinkedListOf[int]([]int{1, 2, 3})
	for i := 1; i <= b.N; i++ {
		err := l.Add(i, i)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = l.Get(i)
	}
}
