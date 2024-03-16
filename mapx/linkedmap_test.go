package mapx

import (
	"errors"
	"generalization_tool"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testErr = errors.New("testMap: put error")

func TestLinkedMap_NewLinkedHashMap(t *testing.T) {
	testCases := []struct {
		name string
		size int
	}{
		{
			name: "illegal size",
			size: -1,
		},
		{
			name: "size is 0",
			size: 0,
		},
		{
			name: "normal",
			size: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			linkedMap := NewLinkedHashMap[testData, int](tc.size)
			assert.NotNil(t, linkedMap)
			assert.Equal(t, linkedMap.Keys(), []testData{})
			assert.Equal(t, linkedMap.Values(), []int{})
		})
	}
}

func TestLinkedMap_NewLinkedTreeMap(t *testing.T) {
	testCases := []struct {
		name       string
		comparator generalization_tool.Comparator[int]
		wantErr    error
	}{
		{
			name:       "no error",
			comparator: generalization_tool.ComparatorRealNumber[int],
			wantErr:    nil,
		},
		{
			name:       "errLinkedTreeMapComparatorIsNull",
			comparator: nil,
			wantErr:    errTreeMapComparatorIsNull,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			linkedTreeMap, err := NewLinkedTreeMap[int, int](tt.comparator)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				assert.Nil(t, linkedTreeMap)
			} else {
				assert.NotNil(t, linkedTreeMap)
				assert.Equal(t, linkedTreeMap.Keys(), []int{})
				assert.Equal(t, linkedTreeMap.Values(), []int{})
			}
		})
	}
}

func TestLinkedMap_Put(t *testing.T) {
	testCases := []struct {
		name       string
		linkedMap  func(t *testing.T) *LinkedMap[int, int]
		keys       []int
		values     []int
		wantKeys   []int
		wantValues []int
		wantErrs   []error
	}{
		{
			name: "put single key",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				return linkedTreeMap
			},
			keys:       []int{1},
			values:     []int{123},
			wantKeys:   []int{1},
			wantValues: []int{123},
			wantErrs:   []error{nil},
		},
		{
			name: "put multiple keys",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				return linkedTreeMap
			},
			keys:       []int{1, 2, 3},
			values:     []int{123, 234, 345},
			wantKeys:   []int{1, 2, 3},
			wantValues: []int{123, 234, 345},
			wantErrs:   []error{nil, nil, nil},
		},
		{
			name: "change value of single key",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				return linkedTreeMap
			},
			keys:       []int{1, 2, 3, 1},
			values:     []int{123, 234, 345, 111},
			wantKeys:   []int{1, 2, 3},
			wantValues: []int{111, 234, 345},
			wantErrs:   []error{nil, nil, nil, nil},
		},
		{
			name: "change value of multiple keys",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				return linkedTreeMap
			},
			keys:       []int{1, 2, 3, 1, 2, 3},
			values:     []int{123, 234, 345, 111, 222, 333},
			wantKeys:   []int{1, 2, 3},
			wantValues: []int{111, 222, 333},
			wantErrs:   []error{nil, nil, nil, nil, nil, nil},
		},
		{
			name: "get error when put single key",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := newLinkedTestMap[int, int](true, generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				return linkedTreeMap
			},
			keys:       []int{1},
			values:     []int{123},
			wantKeys:   []int{},
			wantValues: []int{},
			wantErrs:   []error{testErr},
		},
		{
			name: "get multiple errors when put single key",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := newLinkedTestMap[int, int](true, generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				return linkedTreeMap
			},
			keys:       []int{1, 2, 3, 4, 5, 6},
			values:     []int{123, 234, 345, 456, 567, 678},
			wantKeys:   []int{2, 4, 6},
			wantValues: []int{234, 456, 678},
			wantErrs:   []error{testErr, nil, testErr, nil, testErr, nil},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errs := make([]error, 0)
			lkMap := tc.linkedMap(t)
			for k := range tc.keys {
				err := lkMap.Put(tc.keys[k], tc.values[k])
				errs = append(errs, err)
			}

			for k := range tc.wantKeys {
				value, isFound := lkMap.Get(tc.wantKeys[k])
				assert.Equal(t, true, isFound)
				assert.Equal(t, tc.wantValues[k], value)
			}

			assert.Equal(t, tc.wantKeys, lkMap.Keys())
			assert.Equal(t, tc.wantValues, lkMap.Values())
			assert.Equal(t, tc.wantErrs, errs)
		})
	}
}

func TestLinkedMap_Get(t *testing.T) {
	testCases := []struct {
		name      string
		linkedMap func(t *testing.T) *LinkedMap[int, int]
		target    int
		wantValue int
		isFound   bool
	}{
		{
			name: "cannot find value in empty linked_map",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				return linkedTreeMap
			},
			target:    1,
			wantValue: 0,
			isFound:   false,
		},
		{
			name: "cannot find value in linked_map",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				err = linkedTreeMap.Put(1, 123)
				assert.NoError(t, err)
				return linkedTreeMap
			},
			target:    7,
			wantValue: 0,
			isFound:   false,
		},
		{
			name: "find value in linked_map",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				err = linkedTreeMap.Put(1, 123)
				assert.NoError(t, err)
				return linkedTreeMap
			},
			target:    1,
			wantValue: 123,
			isFound:   true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lkMap := tc.linkedMap(t)
			value, isFound := lkMap.Get(tc.target)
			assert.Equal(t, tc.isFound, isFound)
			assert.Equal(t, tc.wantValue, value)
		})
	}
}

func TestLinkedMap_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		linkedMap   func(t *testing.T) *LinkedMap[int, int]
		target      int
		deleteValue int
		isDeleted   bool
		wantKeys    []int
		wantValues  []int
	}{
		{
			name: "delete key in empty linked_map",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				return linkedTreeMap
			},
			target:      1,
			deleteValue: 0,
			isDeleted:   false,
			wantKeys:    []int{},
			wantValues:  []int{},
		},
		{
			name: "delete unknown key in linked_map",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				err = linkedTreeMap.Put(1, 123)
				assert.NoError(t, err)
				return linkedTreeMap
			},
			target:      2,
			deleteValue: 0,
			isDeleted:   false,
			wantKeys:    []int{1},
			wantValues:  []int{123},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lkMap := tc.linkedMap(t)
			deleteValue, isDeleted := lkMap.Delete(tc.target)
			assert.Equal(t, tc.isDeleted, isDeleted)
			assert.Equal(t, tc.deleteValue, deleteValue)

			assert.Equal(t, tc.wantKeys, lkMap.Keys())
			assert.Equal(t, tc.wantValues, lkMap.Values())
		})
	}
}

func TestLinkedMap_PutAndDelete(t *testing.T) {
	testCases := []struct {
		name       string
		linkedMap  func(t *testing.T) *LinkedMap[int, int]
		wantKeys   []int
		wantValues []int
	}{
		{
			name: "put k1 && delete k1",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				assert.NoError(t, linkedTreeMap.Put(1, 123))
				value, ok := linkedTreeMap.Delete(1)
				assert.Equal(t, 123, value)
				assert.Equal(t, true, ok)
				return linkedTreeMap
			},
			wantKeys:   []int{},
			wantValues: []int{},
		},
		{
			name: "put k1,k2 && delete k1",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				assert.NoError(t, linkedTreeMap.Put(1, 123))
				assert.NoError(t, linkedTreeMap.Put(2, 234))
				value, ok := linkedTreeMap.Delete(1)
				assert.Equal(t, 123, value)
				assert.Equal(t, true, ok)
				return linkedTreeMap
			},
			wantKeys:   []int{2},
			wantValues: []int{234},
		},
		{
			name: "put k1,k2 && delete k2",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				assert.NoError(t, linkedTreeMap.Put(1, 123))
				assert.NoError(t, linkedTreeMap.Put(2, 234))
				value, ok := linkedTreeMap.Delete(2)
				assert.Equal(t, 234, value)
				assert.Equal(t, true, ok)
				return linkedTreeMap
			},
			wantKeys:   []int{1},
			wantValues: []int{123},
		},
		{
			name: "put k1 && delete k1 && put k2,k3",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				assert.NoError(t, linkedTreeMap.Put(1, 123))
				value, ok := linkedTreeMap.Delete(1)
				assert.Equal(t, 123, value)
				assert.Equal(t, true, ok)
				assert.NoError(t, linkedTreeMap.Put(2, 234))
				assert.NoError(t, linkedTreeMap.Put(3, 345))
				return linkedTreeMap
			},
			wantKeys:   []int{2, 3},
			wantValues: []int{234, 345},
		},
		{
			name: "put k1,k2,k3 && delete k2",
			linkedMap: func(t *testing.T) *LinkedMap[int, int] {
				linkedTreeMap, err := NewLinkedTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
				assert.NoError(t, err)
				assert.NoError(t, linkedTreeMap.Put(1, 123))
				assert.NoError(t, linkedTreeMap.Put(2, 234))
				assert.NoError(t, linkedTreeMap.Put(3, 345))
				value, ok := linkedTreeMap.Delete(2)
				assert.Equal(t, 234, value)
				assert.Equal(t, true, ok)
				return linkedTreeMap
			},
			wantKeys:   []int{1, 3},
			wantValues: []int{123, 345},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lkMap := tc.linkedMap(t)
			for k := range tc.wantKeys {
				value, isFound := lkMap.Get(tc.wantKeys[k])
				assert.Equal(t, true, isFound)
				assert.Equal(t, tc.wantValues[k], value)
			}
			assert.Equal(t, tc.wantKeys, lkMap.Keys())
			assert.Equal(t, tc.wantValues, lkMap.Values())
		})
	}
}

type testMap[K any, V any] struct {
	*LinkedMap[K, V]
	count          int
	activeFirstErr bool
}

func newLinkedTestMap[K any, V any](activeFirstErr bool, comparator generalization_tool.Comparator[K]) (*LinkedMap[K, V], error) {
	treeMap, err := NewLinkedTreeMap[K, *linkedKeyValue[K, V]](comparator)
	if err != nil {
		return nil, err
	}
	m := &testMap[K, *linkedKeyValue[K, V]]{
		LinkedMap:      treeMap,
		activeFirstErr: activeFirstErr,
	}
	head := &linkedKeyValue[K, V]{}
	tail := &linkedKeyValue[K, V]{prev: head, next: head}
	head.prev, head.next = tail, tail
	return &LinkedMap[K, V]{
		m:    m,
		head: head,
		tail: tail,
	}, nil
}

func (tm *testMap[K, V]) Put(key K, value V) error {
	tm.count++
	if tm.activeFirstErr {
		tm.activeFirstErr = false
		return testErr
	}
	if tm.count == 3 {
		return testErr
	}
	if tm.count == 5 {
		return testErr
	}
	return tm.LinkedMap.Put(key, value)
}
