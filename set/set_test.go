package setx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSet_Add(t *testing.T) {
	testCases := []struct {
		name    string
		values  []int
		wantSet map[int]struct{}
	}{
		{
			name:   "add",
			values: []int{1, 2, 3, 1, 2},
			wantSet: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewMapSet[int](10)
			for _, val := range tc.values {
				s.Add(val)
			}
			assert.Equal(t, tc.wantSet, s.m)
		})
	}
}

func TestMapSet_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		deleteValue int
		data        map[int]struct{}
		wantSet     map[int]struct{}
		isExist     bool
	}{
		{
			name:        "not found",
			deleteValue: 3,
			data: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
			},
			wantSet: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
			},
			isExist: false,
		},
		{
			name:        "delete",
			deleteValue: 1,
			data: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
			},
			wantSet: map[int]struct{}{
				2: struct{}{},
			},
			isExist: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewMapSet[int](10)
			s.m = tc.data
			s.Delete(tc.deleteValue)
			assert.Equal(t, tc.wantSet, s.m)
		})
	}
}

func TestMapSet_Exist(t *testing.T) {
	s := NewMapSet[int](10)
	s.Add(1)
	testCases := []struct {
		name    string
		value   int
		isExist bool
	}{
		{
			name:    "not found",
			value:   2,
			isExist: false,
		},
		{
			name:    "found",
			value:   1,
			isExist: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := s.Exist(tc.value)
			assert.Equal(t, tc.isExist, ok)
		})
	}
}

func TestMapSet_Keys(t *testing.T) {
	s := NewMapSet[int](10)
	testCases := []struct {
		name      string
		data      map[int]struct{}
		wantValue map[int]struct{}
	}{
		{
			name: "found",
			data: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
			wantValue: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s.m = tc.data
			values := s.Keys()
			isEqual := equal(values, tc.wantValue)
			assert.Equal(t, true, isEqual)
		})
	}
}

func equal(values []int, m map[int]struct{}) bool {
	for _, val := range values {
		_, ok := m[val]
		if !ok {
			return false
		}
		delete(m, val)
	}
	return true && len(m) == 0
}

// goos: windows
// goarch: amd64
// pkg: generalization_tool/set
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkSet
// BenchmarkSet/set_add-8            242268              4213 ns/op
// BenchmarkSet/map_add-8            302539              4087 ns/op
// BenchmarkSet/set_del-8            407031              3238 ns/op
// BenchmarkSet/map_del-8            395287              2997 ns/op
// BenchmarkSet/set_exist-8          447367              2279 ns/op
// BenchmarkSet/map_exist-8          604509              1978 ns/op
func BenchmarkSet(b *testing.B) {
	b.Run("set_add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			s := NewMapSet[int](100)
			b.StartTimer()
			setAdd(s)
		}
	})
	b.Run("map_add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := make(map[int]struct{}, 100)
			b.StartTimer()
			mapAdd(m)
		}
	})
	b.Run("set_del", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			s := NewMapSet[int](100)
			setAdd(s)
			b.StartTimer()
			setDel(s)
		}
	})
	b.Run("map_del", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := make(map[int]struct{}, 100)
			mapAdd(m)
			b.StartTimer()
			mapDel(m)
		}
	})
	b.Run("set_exist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			s := NewMapSet[int](100)
			setAdd(s)
			b.StartTimer()
			setGet(s)
		}
	})
	b.Run("map_exist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := make(map[int]struct{}, 100)
			mapAdd(m)
			b.StartTimer()
			mapGet(m)
		}
	})

}

func setAdd(s Set[int]) {
	for i := 0; i < 100; i++ {
		s.Add(i)
	}
}

func mapAdd(m map[int]struct{}) {
	for i := 0; i < 100; i++ {
		m[i] = struct{}{}
	}
}

func setDel(s Set[int]) {
	for i := 0; i < 100; i++ {
		s.Delete(i)
	}
}

func mapDel(m map[int]struct{}) {
	for i := 0; i < 100; i++ {
		delete(m, i)
	}
}
func setGet(s Set[int]) {
	for i := 0; i < 100; i++ {
		_ = s.Exist(i)
	}
}

func mapGet(s map[int]struct{}) {
	for i := 0; i < 100; i++ {
		_ = s[i]
	}
}
