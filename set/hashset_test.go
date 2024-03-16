package setx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ Set[testData] = &HashSet[testData]{}

func TestNewHashSet(t *testing.T) {
	testCases := []struct {
		name     string
		keys     []int
		wantKeys []testData
	}{
		{
			name:     "nil",
			keys:     nil,
			wantKeys: []testData{},
		},
		{
			name: "normal",
			keys: []int{1, 2, 3},
			wantKeys: []testData{
				{id: 1},
				{id: 2},
				{id: 3},
			},
		},
		{
			name: "duplicate key",
			keys: []int{1, 2, 3, 1, 2},
			wantKeys: []testData{
				{id: 1},
				{id: 2},
				{id: 3},
			},
		},
		{
			name: "disorder key",
			keys: []int{4, 5, 3, 1, 2},
			wantKeys: []testData{
				{id: 1},
				{id: 2},
				{id: 3},
				{id: 4},
				{id: 5},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hashSet := NewHashSet[testData](10)
			for i := 0; i < len(tc.keys); i++ {
				hashSet.Add(testData{id: tc.keys[i]})
			}
			keys := hashSet.Keys()
			assert.ElementsMatch(t, tc.wantKeys, keys)
		})
	}
}

func TestHashSet_Delete(t *testing.T) {
	testCases := []struct {
		name      string
		keys      []int
		addedData []testData
		target    testData
		wantKeys  []testData
	}{
		{
			name:      "nil",
			keys:      nil,
			addedData: []testData{},
			target:    testData{id: 1},
			wantKeys:  []testData{},
		},
		{
			name: "not found",
			keys: []int{1, 2, 3},
			addedData: []testData{
				{id: 1},
				{id: 2},
				{id: 3},
			},
			target: testData{id: 5},
			wantKeys: []testData{
				{id: 1},
				{id: 2},
				{id: 3},
			},
		},
		{
			name: "found",
			keys: []int{1, 2, 3},
			addedData: []testData{
				{id: 1},
				{id: 2},
				{id: 3},
			},
			target: testData{id: 1},
			wantKeys: []testData{
				{id: 2},
				{id: 3},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hashSet := NewHashSet[testData](10)
			for i := 0; i < len(tc.keys); i++ {
				hashSet.Add(testData{id: tc.keys[i]})
			}
			addedKeys := hashSet.Keys()
			assert.ElementsMatch(t, tc.addedData, addedKeys)

			hashSet.Delete(tc.target)
			deletedKeys := hashSet.Keys()
			assert.ElementsMatch(t, tc.wantKeys, deletedKeys)
		})
	}
}

// goos: windows
// goarch: amd64
// pkg: generalization_tool/set
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkNewHashSet
// BenchmarkNewHashSet/hashSet_add-8                 154585            135373 ns/op
// BenchmarkNewHashSet/set_add-8                    8982532               153.7 ns/op
// BenchmarkNewHashSet/map_add-8                    9505281               158.4 ns/op
// BenchmarkNewHashSet/hashSet_exist-8               183991            127410 ns/op
// BenchmarkNewHashSet/set_exist-8                 29882709                64.08 ns/op
// BenchmarkNewHashSet/map_exist-8                 29973024                62.35 ns/op
// BenchmarkNewHashSet/hashSet_del-8                 666644              3353 ns/op
// BenchmarkNewHashSet/set_del-8                   441627805                2.734 ns/op
// BenchmarkNewHashSet/map_del-8                   747382456                1.593 ns/op
func BenchmarkNewHashSet(b *testing.B) {
	hashSet := NewHashSet[testData](100)
	s := NewMapSet[int](100)
	m := make(map[int]int, 100)
	b.Run("hashSet_add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hashSet.Add(testData{id: i})
		}
	})
	b.Run("set_add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Add(i)
		}
	})
	b.Run("map_add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m[i] = i
		}
	})

	b.Run("hashSet_exist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hashSet.Exist(testData{id: i})
		}
	})
	b.Run("set_exist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Exist(i)
		}
	})
	b.Run("map_exist", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = m[i]
		}
	})

	b.Run("hashSet_del", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hashSet.Delete(testData{id: i})
		}
	})
	b.Run("set_del", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Delete(i)
		}
	})
	b.Run("map_del", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			delete(m, i)
		}
	})
}

type testData struct {
	id int
}

func (t testData) Code() uint64 {
	hash := t.id % 10
	return uint64(hash)
}

func (t testData) Equals(key any) bool {
	val, ok := key.(testData)
	if !ok {
		return false
	}
	return t.id == val.id
}

func newTestData(id int) testData {
	return testData{id: id}
}
