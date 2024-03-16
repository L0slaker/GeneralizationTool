package mapx

import (
	"generalization_tool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTreeMapWithMap(t *testing.T) {
	testCases := []struct {
		name       string
		m          map[int]int
		comparable generalization_tool.Comparator[int]
		wantKeys   []int
		wantValues []int
		wantErr    error
	}{
		{
			name:       "nil",
			m:          nil,
			comparable: nil,
			wantKeys:   nil,
			wantValues: nil,
			wantErr:    errTreeMapComparatorIsNull,
		},
		{
			name:       "empty",
			m:          map[int]int{},
			comparable: compare(),
			wantKeys:   nil,
			wantValues: nil,
			wantErr:    nil,
		},
		{
			name: "single",
			m: map[int]int{
				0: 1,
			},
			comparable: compare(),
			wantKeys:   []int{0},
			wantValues: []int{1},
			wantErr:    nil,
		},
		{
			name: "multiple",
			m: map[int]int{
				0: 1,
				1: 2,
				2: 3,
			},
			comparable: compare(),
			wantKeys:   []int{0, 1, 2},
			wantValues: []int{1, 2, 3},
			wantErr:    nil,
		},
		{
			name: "disorder",
			m: map[int]int{
				1: 2,
				2: 3,
				0: 1,
				3: 4,
				5: 6,
				4: 5,
			},
			comparable: compare(),
			wantKeys:   []int{0, 1, 2, 3, 5, 4},
			wantValues: []int{1, 2, 3, 4, 6, 5},
			wantErr:    nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treeMap, err := NewTreeMapWithMap[int, int](tc.comparable, tc.m)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			for k, v := range tc.m {
				value, _ := treeMap.Get(k)
				assert.Equal(t, true, v == value)
			}
		})
	}
}

func TestTreeMap_Get(t *testing.T) {
	testCases := []struct {
		name      string
		m         map[int]int
		target    int
		wantValue int
		isFound   bool
	}{
		{
			name:      "empty treeMap",
			m:         map[int]int{},
			target:    0,
			wantValue: 0,
			isFound:   false,
		},
		{
			name: "found",
			m: map[int]int{
				1: 2,
				2: 3,
				0: 1,
				3: 4,
				5: 6,
				4: 5,
			},
			target:    3,
			wantValue: 4,
			isFound:   true,
		},
		{
			name: "not found",
			m: map[int]int{
				1: 2,
				2: 3,
				0: 1,
				3: 4,
				5: 6,
				4: 5,
			},
			target:    7,
			wantValue: 0,
			isFound:   false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, int](compare())
			putAll(treeMap, tc.m)
			value, found := treeMap.Get(tc.target)
			assert.Equal(t, tc.isFound, found)
			assert.Equal(t, tc.wantValue, value)
		})
	}
}

func TestTreeMap_Put(t *testing.T) {
	testCases := []struct {
		name       string
		keys       []int
		values     []string
		wantKeys   []int
		wantValues []string
		wantErr    error
	}{
		{
			name:       "single",
			keys:       []int{0},
			values:     []string{"a"},
			wantKeys:   []int{0},
			wantValues: []string{"a"},
			wantErr:    nil,
		},
		{
			name:       "multiple",
			keys:       []int{0, 1, 2},
			values:     []string{"a", "b", "c"},
			wantKeys:   []int{0, 1, 2},
			wantValues: []string{"a", "b", "c"},
			wantErr:    nil,
		},
		{
			name:       "same key",
			keys:       []int{0, 0},
			values:     []string{"a", "b"},
			wantKeys:   []int{0},
			wantValues: []string{"b"},
			wantErr:    nil,
		},
		{
			name:       "same keys",
			keys:       []int{0, 1, 0, 1, 2},
			values:     []string{"a", "b", "c", "d", "e"},
			wantKeys:   []int{0, 1, 2},
			wantValues: []string{"c", "d", "e"},
			wantErr:    nil,
		},
		{
			name:       "disorder",
			keys:       []int{1, 2, 0, 5, 3, 4},
			values:     []string{"b", "c", "a", "f", "d", "e"},
			wantKeys:   []int{0, 1, 2, 3, 4, 5},
			wantValues: []string{"a", "b", "c", "d", "e", "f"},
			wantErr:    nil,
		},
		{
			name:       "disorder-same",
			keys:       []int{1, 3, 2, 0, 2, 3},
			values:     []string{"a", "c", "b", "d", "e", "f"},
			wantKeys:   []int{0, 1, 2, 3},
			wantValues: []string{"d", "a", "e", "f"},
			wantErr:    nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, string](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := treeMap.Put(tc.keys[i], tc.values[i])
				if err != nil {
					assert.Equal(t, tc.wantErr, err)
					return
				}
			}
		})
	}

	subTests := []struct {
		name       string
		keys       []int
		values     []string
		wantKeys   []int
		wantValues []string
		wantErr    error
	}{
		{
			name:       "nil",
			keys:       []int{0},
			values:     nil,
			wantKeys:   []int{0},
			wantValues: []string(nil),
		},
		{
			name:       "nil",
			keys:       []int{0},
			values:     []string{"0"},
			wantKeys:   []int{0},
			wantValues: []string{"0"},
		},
	}
	for _, st := range subTests {
		t.Run(st.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, []string](compare())
			for i := 0; i < len(st.keys); i++ {
				err := treeMap.Put(st.keys[i], st.values)
				if err != nil {
					assert.Equal(t, st.wantErr, err)
					return
				}
			}
			for i := 0; i < len(st.wantKeys); i++ {
				v, b := treeMap.Get(st.wantKeys[i])
				assert.Equal(t, true, b)
				assert.Equal(t, st.wantValues, v)
			}

		})
	}
}

func TestTreeMap_Keys(t *testing.T) {
	testCases := []struct {
		name     string
		m        map[int]int
		wantKeys []int
	}{
		{
			name:     "empty",
			wantKeys: []int{},
		},
		{
			name: "normal",
			m: map[int]int{
				1: 123,
				2: 234,
				3: 345,
				4: 456,
				5: 567,
			},
			wantKeys: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, int](compare())
			for k, v := range tc.m {
				_ = treeMap.Put(k, v)
			}
			keys := treeMap.Keys()
			assert.Equal(t, tc.wantKeys, keys)
		})
	}
}

func TestTreeMap_Values(t *testing.T) {
	testCases := []struct {
		name       string
		m          map[int]int
		wantValues []int
	}{
		{
			name:       "empty",
			wantValues: []int{},
		},
		{
			name: "normal",
			m: map[int]int{
				1: 123,
				2: 234,
				3: 345,
				4: 456,
				5: 567,
			},
			wantValues: []int{123, 234, 345, 456, 567},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, int](compare())
			for k, v := range tc.m {
				_ = treeMap.Put(k, v)
			}
			values := treeMap.Values()
			assert.Equal(t, tc.wantValues, values)
		})
	}
}

func TestTreeMap_Delete(t *testing.T) {
	testCases := []struct {
		name      string
		m         map[int]int
		target    int
		deleteVal int
		isDelete  bool
	}{
		{
			name:      "empty",
			m:         map[int]int{},
			deleteVal: 0,
		},
		{
			name: "found",
			m: map[int]int{
				1: 123,
				2: 234,
				3: 345,
				4: 456,
				5: 567,
			},
			target:    1,
			deleteVal: 123,
			isDelete:  true,
		},
		{
			name: "not found",
			m: map[int]int{
				1: 123,
				2: 234,
				3: 345,
				4: 456,
				5: 567,
			},
			target:    7,
			deleteVal: 0,
			isDelete:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, int](compare())
			for k, v := range tc.m {
				_ = treeMap.Put(k, v)
			}
			deleteVal, isDelete := treeMap.Delete(tc.target)
			assert.Equal(t, tc.isDelete, isDelete)
			assert.Equal(t, tc.deleteVal, deleteVal)
			_, ok := treeMap.Get(tc.deleteVal)
			assert.False(t, ok)
		})
	}
}

// goos: windows
// goarch: amd64
// pkg: generalization_tool/mapx
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkNewTreeMap
// BenchmarkNewTreeMap/treeMap_put-8                1657771               613.3 ns/ op
// BenchmarkNewTreeMap/map_put-8                    5248268               196.0 ns/ op
// BenchmarkNewTreeMap/hashMap_put-8                4745463               251.6 ns/ op
// BenchmarkNewTreeMap/treeMap_get-8                6568500               266.5 ns/ op
// BenchmarkNewTreeMap/map_get-8                   12050006               114.7 ns/ op
// BenchmarkNewTreeMap/hashMap_get-8                8785602               130.5 ns/ op
func BenchmarkNewTreeMap(b *testing.B) {
	hashMap := NewHashMap[hashInt, int](10)
	treeMap, _ := NewTreeMap[uint64, int](generalization_tool.ComparatorRealNumber[uint64])
	m := make(map[uint64]int, 10)
	b.Run("treeMap_put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = treeMap.Put(uint64(i), i)
		}
	})
	b.Run("map_put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m[uint64(i)] = i
		}
	})
	b.Run("hashMap_put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hashMap.Put(hashInt(uint64(i)), i)
		}
	})
	b.Run("treeMap_get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = treeMap.Get(uint64(i))
		}
	})
	b.Run("map_get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = m[uint64(i)]
		}
	})
	b.Run("hashMap_get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = hashMap.Get(hashInt(uint64(i)))
		}
	})
}

func compare() generalization_tool.Comparator[int] {
	return generalization_tool.ComparatorRealNumber[int]
}
