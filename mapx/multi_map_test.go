package mapx

import (
	"generalization_tool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiMap_NewMultiHashMap(t *testing.T) {
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
			multiMap := NewMultiHashMap[testData, int](tc.size)
			assert.NotNil(t, multiMap)
		})
	}
}

func TestMultiMap_NewMultiTreeMap(t *testing.T) {
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
			name:       "errTreeMapComparatorIsNull",
			comparator: nil,
			wantErr:    errTreeMapComparatorIsNull,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			multiMap, err := NewMultiTreeMap[int, int](tc.comparator)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				assert.Nil(t, multiMap)
			} else {
				assert.NotNil(t, multiMap)
			}
		})
	}
}

func TestNewMultiBuiltinMap(t *testing.T) {
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
			multiMap := NewMultiBuiltinMap[testData, int](tc.size)
			assert.NotNil(t, multiMap)
		})
	}
}

func TestMultiMap_Keys(t *testing.T) {
	testCases := []struct {
		name                 string
		multiTreeMap         *MultiMap[int, int]
		multiHashMap         *MultiMap[testData, int]
		wantMultiTreeMapKeys []int
		wantMultiHashMapKeys []testData
	}{
		{
			name: "empty",
			multiTreeMap: func() *MultiMap[int, int] {
				return getMultiTreeMap()
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				return getMultiHashMap()
			}(),
			wantMultiTreeMapKeys: []int{},
			wantMultiHashMapKeys: []testData{},
		},
		{
			name: "single one",
			multiTreeMap: func() *MultiMap[int, int] {
				multiTreeMap := getMultiTreeMap()
				_ = multiTreeMap.Put(1, 1)
				return multiTreeMap
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				multiHashMap := getMultiHashMap()
				_ = multiHashMap.Put(testData{id: 1}, 1)
				return multiHashMap
			}(),
			wantMultiTreeMapKeys: []int{1},
			wantMultiHashMapKeys: []testData{{id: 1}},
		},
		{
			name: "multiple",
			multiTreeMap: func() *MultiMap[int, int] {
				multiTreeMap := getMultiTreeMap()
				_ = multiTreeMap.Put(1, 1)
				_ = multiTreeMap.Put(2, 2)
				_ = multiTreeMap.Put(3, 3)
				return multiTreeMap
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				multiHashMap := getMultiHashMap()
				_ = multiHashMap.Put(testData{id: 1}, 1)
				_ = multiHashMap.Put(testData{id: 2}, 2)
				_ = multiHashMap.Put(testData{id: 3}, 3)
				return multiHashMap
			}(),
			wantMultiTreeMapKeys: []int{1, 2, 3},
			wantMultiHashMapKeys: []testData{
				{id: 1},
				{id: 2},
				{id: 3},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.wantMultiTreeMapKeys, tc.multiTreeMap.Keys())
		})
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.wantMultiHashMapKeys, tc.multiHashMap.Keys())
		})
	}
}

func TestMultiMap_Values(t *testing.T) {
	testCases := []struct {
		name         string
		multiTreeMap *MultiMap[int, int]
		multiHashMap *MultiMap[testData, int]
		wantValues   [][]int
	}{
		{
			name: "empty",
			multiTreeMap: func() *MultiMap[int, int] {
				return getMultiTreeMap()
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				return getMultiHashMap()
			}(),
			wantValues: [][]int{},
		},
		{
			name: "single one",
			multiTreeMap: func() *MultiMap[int, int] {
				multiTreeMap := getMultiTreeMap()
				_ = multiTreeMap.Put(1, 1)
				return multiTreeMap
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				multiHashMap := getMultiHashMap()
				_ = multiHashMap.Put(testData{id: 1}, 1)
				return multiHashMap
			}(),
			wantValues: [][]int{{1}},
		},
		{
			name: "multiple",
			multiTreeMap: func() *MultiMap[int, int] {
				multiTreeMap := getMultiTreeMap()
				_ = multiTreeMap.Put(1, 1)
				_ = multiTreeMap.Put(2, 2)
				_ = multiTreeMap.Put(3, 3)
				return multiTreeMap
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				multiHashMap := getMultiHashMap()
				_ = multiHashMap.Put(testData{id: 1}, 1)
				_ = multiHashMap.Put(testData{id: 2}, 2)
				_ = multiHashMap.Put(testData{id: 3}, 3)
				return multiHashMap
			}(),
			wantValues: [][]int{
				{1},
				{2},
				{3},
			},
		},
	}
	t.Run("MultiTreeMap", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				assert.ElementsMatch(t, tc.wantValues, tc.multiTreeMap.Values())
			})
		}
	})
	t.Run("MultiHashMap", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				assert.ElementsMatch(t, tc.wantValues, tc.multiHashMap.Values())
			})
		}
	})
}

func TestMultiMap_Put(t *testing.T) {
	testCases := []struct {
		name       string
		keys       []int
		values     []int
		wantKeys   []int
		wantValues [][]int
		wantErr    error
	}{
		{
			name:       "single key with single value",
			keys:       []int{1},
			values:     []int{1},
			wantKeys:   []int{1},
			wantValues: [][]int{{1}},
			wantErr:    nil,
		},
		{
			name:     "multiple keys with single value",
			keys:     []int{1, 2, 3},
			values:   []int{1, 2, 3},
			wantKeys: []int{1, 2, 3},
			wantValues: [][]int{
				{1},
				{2},
				{3},
			},
			wantErr: nil,
		},
		{
			name:     "multiple keys with multiple value",
			keys:     []int{1, 2, 3, 1, 1},
			values:   []int{1, 2, 3, 4, 5},
			wantKeys: []int{1, 2, 3},
			wantValues: [][]int{
				{1, 4, 5},
				{2},
				{3},
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run("MultiTreeMap", func(t *testing.T) {
			multiTreeMap, _ := NewMultiTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
			for k := range tc.keys {
				err := multiTreeMap.Put(tc.keys[k], tc.values[k])
				assert.Equal(t, tc.wantErr, err)
			}
			for k := range tc.wantKeys {
				values, isFound := multiTreeMap.Get(tc.wantKeys[k])
				assert.Equal(t, true, isFound)
				assert.Equal(t, tc.wantValues[k], values)
			}
		})
		t.Run("MultiHashMap", func(t *testing.T) {
			multiHashMap := NewMultiHashMap[testData, int](10)
			for k := range tc.keys {
				err := multiHashMap.Put(testData{id: tc.keys[k]}, tc.values[k])
				assert.Equal(t, tc.wantErr, err)
			}
			for k := range tc.wantKeys {
				values, isFound := multiHashMap.Get(testData{id: tc.keys[k]})
				assert.Equal(t, true, isFound)
				assert.Equal(t, tc.wantValues[k], values)
			}
		})
	}
}

func TestMultiMap_Get(t *testing.T) {
	testCases := []struct {
		name         string
		multiTreeMap *MultiMap[int, int]
		multiHashMap *MultiMap[testData, int]
		target       int
		wantValues   []int
		isFound      bool
	}{
		{
			name: "not found in empty data",
			multiTreeMap: func() *MultiMap[int, int] {
				return getMultiTreeMap()
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				return getMultiHashMap()
			}(),
			target:     1,
			wantValues: nil,
			isFound:    false,
		},
		{
			name: "not found in data",
			multiTreeMap: func() *MultiMap[int, int] {
				multiTreeMap := getMultiTreeMap()
				_ = multiTreeMap.Put(1, 1)
				_ = multiTreeMap.Put(2, 2)
				return multiTreeMap
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				multiHashMap := getMultiHashMap()
				_ = multiHashMap.Put(testData{id: 1}, 1)
				_ = multiHashMap.Put(testData{id: 2}, 2)
				return multiHashMap
			}(),
			target:     3,
			wantValues: nil,
			isFound:    false,
		},
		{
			name: "found",
			multiTreeMap: func() *MultiMap[int, int] {
				multiTreeMap := getMultiTreeMap()
				_ = multiTreeMap.Put(1, 1)
				_ = multiTreeMap.Put(2, 2)
				return multiTreeMap
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				multiHashMap := getMultiHashMap()
				_ = multiHashMap.Put(testData{id: 1}, 1)
				_ = multiHashMap.Put(testData{id: 2}, 2)
				return multiHashMap
			}(),
			target:     1,
			wantValues: []int{1},
			isFound:    true,
		},
	}
	for _, tc := range testCases {
		t.Run("MultiTreeMap", func(t *testing.T) {
			values, isFound := tc.multiTreeMap.Get(tc.target)
			assert.Equal(t, tc.isFound, isFound)
			assert.ElementsMatch(t, tc.wantValues, values)
		})

		t.Run("MultiHashMap", func(t *testing.T) {

		})
	}
}

func TestMultiMap_Delete(t *testing.T) {
	testCases := []struct {
		name         string
		multiTreeMap *MultiMap[int, int]
		multiHashMap *MultiMap[testData, int]
		target       int
		deleteValues []int
		isDeleted    bool
	}{
		{
			name: "not found in empty data",
			multiTreeMap: func() *MultiMap[int, int] {
				return getMultiTreeMap()
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				return getMultiHashMap()
			}(),
			target:       1,
			deleteValues: nil,
			isDeleted:    false,
		},
		{
			name: "not found in data",
			multiTreeMap: func() *MultiMap[int, int] {
				multiTreeMap := getMultiTreeMap()
				_ = multiTreeMap.Put(1, 1)
				_ = multiTreeMap.Put(2, 2)
				return multiTreeMap
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				multiHashMap := getMultiHashMap()
				_ = multiHashMap.Put(testData{id: 1}, 1)
				_ = multiHashMap.Put(testData{id: 2}, 2)
				return multiHashMap
			}(),
			target:       3,
			deleteValues: nil,
			isDeleted:    false,
		},
		{
			name: "found and deleted",
			multiTreeMap: func() *MultiMap[int, int] {
				multiTreeMap := getMultiTreeMap()
				_ = multiTreeMap.Put(1, 1)
				_ = multiTreeMap.Put(2, 2)
				return multiTreeMap
			}(),
			multiHashMap: func() *MultiMap[testData, int] {
				multiHashMap := getMultiHashMap()
				_ = multiHashMap.Put(testData{id: 1}, 1)
				_ = multiHashMap.Put(testData{id: 2}, 2)
				return multiHashMap
			}(),
			target:       1,
			deleteValues: []int{1},
			isDeleted:    true,
		},
	}
	for _, tc := range testCases {
		t.Run("MultiTreeMap", func(t *testing.T) {
			deleteValues, isDelete := tc.multiTreeMap.Delete(tc.target)
			assert.Equal(t, tc.isDeleted, isDelete)
			assert.ElementsMatch(t, tc.deleteValues, deleteValues)
			_, ok := tc.multiTreeMap.Get(tc.target)
			assert.False(t, ok)
		})

		t.Run("MultiHashMap", func(t *testing.T) {
			deleteValues, isDelete := tc.multiHashMap.Delete(testData{id: tc.target})
			assert.Equal(t, tc.isDeleted, isDelete)
			assert.ElementsMatch(t, tc.deleteValues, deleteValues)
			_, ok := tc.multiHashMap.Get(testData{id: tc.target})
			assert.False(t, ok)
		})
	}
}

func TestMultiMap_PutMany(t *testing.T) {
	testCases := []struct {
		name       string
		keys       []int
		values     [][]int
		wantKeys   []int
		wantValues [][]int
		wantErr    error
	}{
		{
			name:       "one to one",
			keys:       []int{1},
			values:     [][]int{{1}},
			wantKeys:   []int{1},
			wantValues: [][]int{{1}},
			wantErr:    nil,
		},
		{
			name:       "multiple `one to one`",
			keys:       []int{1, 2, 3},
			values:     [][]int{{1}, {2}, {3}},
			wantKeys:   []int{1, 2, 3},
			wantValues: [][]int{{1}, {2}, {3}},
			wantErr:    nil,
		},
		{
			name:       "one to multiple",
			keys:       []int{1},
			values:     [][]int{{1, 2, 3}},
			wantKeys:   []int{1},
			wantValues: [][]int{{1, 2, 3}},
			wantErr:    nil,
		},
		{
			name: "multiple to multiple",
			keys: []int{1, 2, 3},
			values: [][]int{
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 5},
			},
			wantKeys: []int{1, 2, 3},
			wantValues: [][]int{
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 5},
			},
			wantErr: nil,
		},
		{
			name: "same key append value",
			keys: []int{1, 1},
			values: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			wantKeys: []int{1},
			wantValues: [][]int{
				{1, 2, 3, 4, 5, 6},
			},
			wantErr: nil,
		},
		{
			name: "same key append value",
			keys: []int{1, 2, 1},
			values: [][]int{
				{1, 2, 3},
				{2, 3, 4},
				{4, 5, 6},
			},
			wantKeys: []int{1, 2},
			wantValues: [][]int{
				{1, 2, 3, 4, 5, 6},
				{2, 3, 4},
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run("MultiTreeMap", func(t *testing.T) {
			multiTreeMap, _ := NewMultiTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
			for k := range tc.keys {
				err := multiTreeMap.PutMany(tc.keys[k], tc.values[k]...)
				assert.Equal(t, tc.wantErr, err)
			}
			for k := range tc.wantKeys {
				values, isFound := multiTreeMap.Get(tc.wantKeys[k])
				assert.Equal(t, true, isFound)
				assert.Equal(t, tc.wantValues[k], values)
			}
		})
		t.Run("MultiHashMap", func(t *testing.T) {
			multiHashMap := NewMultiHashMap[testData, int](10)
			for k := range tc.keys {
				err := multiHashMap.PutMany(testData{id: tc.keys[k]}, tc.values[k]...)
				assert.Equal(t, tc.wantErr, err)
			}
			for k := range tc.wantKeys {
				values, isFound := multiHashMap.Get(testData{id: tc.wantKeys[k]})
				assert.Equal(t, true, isFound)
				assert.Equal(t, tc.wantValues[k], values)
			}
		})
	}
}

func getMultiTreeMap() *MultiMap[int, int] {
	multiTreeMap, _ := NewMultiTreeMap[int, int](generalization_tool.ComparatorRealNumber[int])
	return multiTreeMap
}

func getMultiHashMap() *MultiMap[testData, int] {
	return NewMultiHashMap[testData, int](10)
}
