package mapx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var _ mapi[testData, int] = &HashMap[testData, int]{}

func TestHashMap_Get_PUT(t *testing.T) {
	cases := []struct {
		key testData
		val int
	}{
		{
			key: testData{id: 1},
			val: 1,
		},
		{
			key: testData{id: 2},
			val: 2,
		},
		{
			key: testData{id: 3},
			val: 3,
		},
		{
			key: testData{id: 11},
			val: 11,
		},
		{
			key: testData{id: 1},
			val: 101,
		},
	}
	mymap := NewHashMap[testData, int](10)
	for _, cs := range cases {
		if err := mymap.Put(cs.key, cs.val); err != nil {
			require.NoError(t, err)
		}
	}

	wantMap := NewHashMap[testData, int](10)
	wantMap.hashmap = map[uint64]*node[testData, int]{
		1: &node[testData, int]{
			key:   testData{id: 1},
			value: 101,
			next: &node[testData, int]{
				key:   testData{id: 11},
				value: 11,
			},
		},
		2: wantMap.newNode(newTestData(2), 2),
		3: wantMap.newNode(newTestData(3), 3),
	}
	assert.Equal(t, wantMap.hashmap, mymap.hashmap)

	testCase := []struct {
		name    string
		key     testData
		wantVal any
		found   bool
	}{
		{
			name:    "get val successfully",
			key:     testData{id: 1},
			wantVal: 101,
			found:   true,
		},
		{
			name:    "hash conflicts",
			key:     testData{id: 11},
			wantVal: 11,
			found:   true,
		},
		{
			name:  "hash not Found",
			key:   testData{id: 9},
			found: false,
		},
		{
			name:  "val not Found",
			key:   testData{id: 31},
			found: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			val, ok := mymap.Get(tc.key)
			assert.Equal(t, tc.found, ok)
			if !ok {
				return
			}
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestHashMap_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		key         testData
		setHashMap  func() map[uint64]*node[testData, int]
		wantHashMap func() map[uint64]*node[testData, int]
		wantVal     int
		found       bool
	}{
		{
			name: "hash not found",
			setHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{}
			},
			found: false,
			key:   testData{id: 1},
		},
		{
			name: "key not found",
			setHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{
					1: &node[testData, int]{
						key:   testData{id: 1},
						value: 1,
					},
				}
			},
			found: false,
			key:   testData{id: 11},
		},
		{
			name: "delete only one link element",
			setHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{
					1: &node[testData, int]{
						key:   testData{id: 1},
						value: 1,
					},
				}
			},
			found:   true,
			key:     testData{id: 1},
			wantVal: 1,
			wantHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{}
			},
		},
		{
			name: "many link elements delete first",
			setHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{
					1: &node[testData, int]{
						key:   testData{id: 1},
						value: 1,
						next: &node[testData, int]{
							key:   testData{id: 11},
							value: 11,
							next: &node[testData, int]{
								key:   testData{id: 21},
								value: 21,
							},
						},
					},
				}
			},
			found:   true,
			key:     testData{id: 1},
			wantVal: 1,
			wantHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{
					11: &node[testData, int]{
						key:   testData{id: 11},
						value: 11,
						next: &node[testData, int]{
							key:   testData{id: 21},
							value: 21,
						},
					},
				}
			},
		},
		{
			name: "delete link elements delete middle",
			setHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{
					1: &node[testData, int]{
						key:   testData{id: 1},
						value: 1,
						next: &node[testData, int]{
							key:   testData{id: 11},
							value: 11,
							next: &node[testData, int]{
								key:   testData{id: 21},
								value: 21,
							},
						},
					},
				}
			},
			found:   true,
			wantVal: 11,
			key:     testData{id: 11},
			wantHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{
					1: &node[testData, int]{
						key:   testData{id: 1},
						value: 1,
						next: &node[testData, int]{
							key:   testData{id: 21},
							value: 21,
						},
					},
				}
			},
		},
		{
			name: "delete link elements delete end",
			setHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{
					1: &node[testData, int]{
						key:   testData{id: 1},
						value: 1,
						next: &node[testData, int]{
							key:   testData{id: 11},
							value: 11,
							next: &node[testData, int]{
								key:   testData{id: 21},
								value: 21,
							},
						},
					},
				}
			},
			found:   true,
			key:     testData{id: 21},
			wantVal: 21,
			wantHashMap: func() map[uint64]*node[testData, int] {
				return map[uint64]*node[testData, int]{
					1: &node[testData, int]{
						key:   testData{id: 1},
						value: 1,
						next: &node[testData, int]{
							key:   testData{id: 11},
							value: 11,
						},
					},
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mymap := NewHashMap[testData, int](10)
			mymap.hashmap = tc.setHashMap()
			val, ok := mymap.Delete(tc.key)
			assert.Equal(t, tc.found, ok)
			if !ok {
				return
			}
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestHashMap_Keys_Values(t *testing.T) {
	testCases := []struct {
		name       string
		getHashMap func() *HashMap[testData, int]
		wantKeys   []Hashable
		wantValues []int
	}{
		{
			name: "empty",
			getHashMap: func() *HashMap[testData, int] {
				return &HashMap[testData, int]{}
			},
			wantKeys:   []Hashable{},
			wantValues: []int{},
		},
		{
			name: "size is zero empty",
			getHashMap: func() *HashMap[testData, int] {
				return NewHashMap[testData, int](0)
			},
			wantKeys:   []Hashable{},
			wantValues: []int{},
		},
		{
			name: "single key",
			getHashMap: func() *HashMap[testData, int] {
				testHashMap := NewHashMap[testData, int](10)
				err := testHashMap.Put(newTestData(1), 1)
				require.NoError(t, err)
				return testHashMap
			},
			wantKeys:   []Hashable{newTestData(1)},
			wantValues: []int{1},
		},
		{
			name: "multiple keys",
			getHashMap: func() *HashMap[testData, int] {
				testHashMap := NewHashMap[testData, int](10)
				for _, v := range []int{1, 2} {
					err := testHashMap.Put(newTestData(v), v)
					require.NoError(t, err)
				}
				return testHashMap
			},
			wantKeys:   []Hashable{newTestData(1), newTestData(2)},
			wantValues: []int{1, 2},
		},
		{
			name: "same key",
			getHashMap: func() *HashMap[testData, int] {
				testHashMap := NewHashMap[testData, int](10)
				err := testHashMap.Put(newTestData(1), 1)
				require.NoError(t, err)
				err = testHashMap.Put(newTestData(1), 11)
				require.NoError(t, err)
				return testHashMap
			},
			wantKeys:   []Hashable{newTestData(1)},
			wantValues: []int{11},
		},
		{
			name: "multiple with same key",
			getHashMap: func() *HashMap[testData, int] {
				testHashMap := NewHashMap[testData, int](10)
				for _, v := range []int{1, 2, 3} {
					err := testHashMap.Put(newTestData(v), v*10)
					require.NoError(t, err)
				}
				err := testHashMap.Put(newTestData(1), 11)
				require.NoError(t, err)
				return testHashMap
			},
			wantKeys:   []Hashable{newTestData(1), newTestData(2), newTestData(3)},
			wantValues: []int{11, 20, 30},
		},
		{
			name: "single key collision",
			getHashMap: func() *HashMap[testData, int] {
				testHashMap := NewHashMap[testData, int](10)
				err := testHashMap.Put(newTestData(1), 11)
				require.NoError(t, err)
				err = testHashMap.Put(newTestData(11), 111)
				require.NoError(t, err)
				err = testHashMap.Put(newTestData(111), 1111)
				require.NoError(t, err)
				return testHashMap
			},
			wantKeys:   []Hashable{newTestData(1), newTestData(11), newTestData(111)},
			wantValues: []int{11, 111, 1111},
		},
		{
			name: "multiple key collision",
			getHashMap: func() *HashMap[testData, int] {
				testHashMap := NewHashMap[testData, int](10)
				for _, v := range []int{1, 2, 3} {
					err := testHashMap.Put(newTestData(v), v)
					require.NoError(t, err)
					err = testHashMap.Put(newTestData(v*10+v), v*10)
					require.NoError(t, err)
				}
				return testHashMap
			},
			wantKeys: []Hashable{
				newTestData(1), newTestData(11),
				newTestData(2), newTestData(22),
				newTestData(3), newTestData(33),
			},
			wantValues: []int{1, 10, 2, 20, 3, 30},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keys := tc.getHashMap().Keys()
			values := tc.getHashMap().Values()
			// 判断长度
			assert.Equal(t, len(tc.wantKeys), len(keys))
			assert.Equal(t, len(tc.wantValues), len(values))
			// 判断元素
			assert.ElementsMatch(t, tc.wantKeys, keys)
			assert.ElementsMatch(t, tc.wantValues, values)
		})
	}
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

type hashInt uint64

func (h hashInt) Code() uint64 {
	return uint64(h)
}

func (h hashInt) Equals(key any) bool {
	switch keyVal := key.(type) {
	case hashInt:
		return keyVal == h
	default:
		return false
	}
}

func newHashInt(i int) hashInt {
	return hashInt(i)
}
