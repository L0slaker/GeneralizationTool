package setx

import (
	"generalization_tool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewTreeSet(t *testing.T) {
	testCases := []struct {
		name       string
		keys       []int
		comparator generalization_tool.Comparator[int]
		wantKeys   []int
		wantErr    string
	}{
		{
			name:       "nil comparator",
			comparator: nil,
			keys:       nil,
			wantKeys:   []int{},
			wantErr:    "TreeMap：Comparator不能为nil",
		},
		{
			name:       "nil key",
			keys:       nil,
			comparator: compare(),
			wantKeys:   []int{},
		},
		{
			name:       "duplicate key",
			keys:       []int{0, 1, 2, 3, 3, 2},
			comparator: compare(),
			wantKeys:   []int{0, 1, 2, 3},
		},
		{
			name:       "disorder key",
			keys:       []int{0, 2, 1, 4, 3},
			comparator: compare(),
			wantKeys:   []int{0, 1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treeSet, err := NewTreeSet[int](tc.comparator)
			if err != nil {
				assert.Equal(t, tc.wantErr, err.Error())
				return
			}
			for i := 0; i < len(tc.keys); i++ {
				treeSet.Add(tc.keys[i])
			}
			keys := treeSet.Keys()
			assert.ElementsMatch(t, tc.wantKeys, keys)
		})
	}
}

func TestTreeSet_Delete(t *testing.T) {
	testCases := []struct {
		name     string
		keys     []int
		target   int
		wantKeys []int
	}{
		{
			name:     "nil",
			keys:     nil,
			target:   0,
			wantKeys: []int{},
		},
		{
			name:     "not found",
			keys:     []int{1, 2, 3},
			target:   7,
			wantKeys: []int{1, 2, 3},
		},
		{
			name:     "found",
			keys:     []int{1, 2, 3, 4, 5, 6, 7},
			target:   7,
			wantKeys: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treeSet, err := NewTreeSet[int](compare())
			require.NoError(t, err)
			for i := 0; i < len(tc.keys); i++ {
				treeSet.Add(tc.keys[i])
			}
			treeSet.Delete(tc.target)
			keys := treeSet.Keys()
			assert.Equal(t, tc.wantKeys, keys)
		})
	}
}

func TestTreeSet_Exist(t *testing.T) {
	testCases := []struct {
		name    string
		keys    []int
		target  int
		isExist bool
	}{
		{
			name:    "nil",
			keys:    nil,
			target:  0,
			isExist: false,
		},
		{
			name:    "not found",
			keys:    []int{1, 2, 3},
			target:  7,
			isExist: false,
		},
		{
			name:    "found",
			keys:    []int{1, 2, 3, 4, 5, 6, 7},
			target:  7,
			isExist: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treeSet, err := NewTreeSet[int](compare())
			require.NoError(t, err)
			for i := 0; i < len(tc.keys); i++ {
				treeSet.Add(tc.keys[i])
			}
			isExist := treeSet.Exist(tc.target)
			assert.Equal(t, tc.isExist, isExist)
		})
	}
}

// goos: windows
// goarch: amd64
// pkg: generalization_tool/set
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkNewTreeSet
// BenchmarkNewTreeSet/treeSet_add-8                3415324               372.3 ns/op
// BenchmarkNewTreeSet/set_add-8                    9298821               157.5 ns/op
// BenchmarkNewTreeSet/map_add-8                    9710834               162.9 ns/op
// BenchmarkNewTreeSet/treeSet_exist-8              8682106               173.4 ns/op
// BenchmarkNewTreeSet/set_exist-8                 30738885                62.75 ns/op
// BenchmarkNewTreeSet/set_exist#01-8              29999324                62.39 ns/op
// BenchmarkNewTreeSet/treeSet_del-8               294773630                4.095 ns/op
// BenchmarkNewTreeSet/set_del-8                   431158255                2.756 ns/op
// BenchmarkNewTreeSet/set_del#01-8                774206042                1.603 ns/op
func BenchmarkNewTreeSet(b *testing.B) {
	treeSet, err := NewTreeSet[uint64](generalization_tool.ComparatorRealNumber[uint64])
	require.NoError(b, err)
	s := NewMapSet[uint64](100)
	m := make(map[uint64]int, 100)
	b.Run("treeSet_add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			treeSet.Add(uint64(i))
		}
	})
	b.Run("set_add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Add(uint64(i))
		}
	})
	b.Run("map_add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m[uint64(i)] = i
		}
	})

	b.Run("treeSet_exist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			treeSet.Exist(uint64(i))
		}
	})
	b.Run("set_exist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Exist(uint64(i))
		}
	})
	b.Run("set_exist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = m[uint64(i)]
		}
	})

	b.Run("treeSet_del", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			treeSet.Delete(uint64(i))
		}
	})
	b.Run("set_del", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Delete(uint64(i))
		}
	})
	b.Run("set_del", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			delete(m, uint64(i))
		}
	})
}

func compare() generalization_tool.Comparator[int] {
	return generalization_tool.ComparatorRealNumber[int]
}
