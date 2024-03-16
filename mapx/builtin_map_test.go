package mapx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuiltinMap_Delete(t *testing.T) {
	testCases := []struct {
		name      string
		m         map[string]string
		target    string
		wantValue string
		isDeleted bool
	}{
		{
			name:      "nil",
			m:         map[string]string{},
			target:    "book-1",
			isDeleted: false,
		},
		{
			name: "not exist key",
			m: map[string]string{
				"book-1": "nature",
			},
			target:    "book-2",
			isDeleted: false,
		},
		{
			name: "delete successfully",
			m: map[string]string{
				"book-1": "nature",
			},
			target:    "book-1",
			wantValue: "nature",
			isDeleted: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := newBuiltinMapOf[string, string](tc.m)
			value, isDelete := m.Delete(tc.target)
			assert.Equal(t, tc.isDeleted, isDelete)
			assert.Equal(t, tc.wantValue, value)
			_, ok := m.Get(tc.target)
			assert.False(t, ok)
		})
	}
}

func TestBuiltinMap_Get(t *testing.T) {
	testCases := []struct {
		name      string
		m         map[string]string
		target    string
		wantValue string
		isFound   bool
	}{
		{
			name:    "nil",
			m:       map[string]string{},
			target:  "book-2",
			isFound: false,
		},
		{
			name: "not exist key",
			m: map[string]string{
				"book-1": "nature",
			},
			target:  "book-2",
			isFound: false,
		},
		{
			name: "found",
			m: map[string]string{
				"book-1": "nature",
			},
			target:    "book-1",
			wantValue: "nature",
			isFound:   true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := newBuiltinMapOf[string, string](tc.m)
			value, isFound := m.Get(tc.target)
			assert.Equal(t, tc.isFound, isFound)
			assert.Equal(t, tc.wantValue, value)
		})
	}
}

func TestBuiltinMap_Put(t *testing.T) {
	testCases := []struct {
		name    string
		key     string
		value   string
		cap     int
		wantErr error
	}{
		{
			name:  "put",
			key:   "book-1",
			value: "nature",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := newBuiltinMap[string, string](tc.cap)
			err := m.Put(tc.key, tc.value)
			assert.Equal(t, tc.wantErr, err)
			value, ok := m.data[tc.key]
			assert.True(t, ok)
			assert.Equal(t, tc.value, value)
		})
	}
}

func TestBuiltinMap_Keys(t *testing.T) {
	testCases := []struct {
		name     string
		m        map[string]string
		wantKeys []string
	}{
		{
			name:     "nil",
			wantKeys: []string{},
		},
		{
			name:     "empty",
			m:        map[string]string{},
			wantKeys: []string{},
		},
		{
			name: "keys",
			m: map[string]string{
				"book-1": "nature",
				"book-2": "world",
				"book-3": "peace",
			},
			wantKeys: []string{"book-1", "book-2", "book-3"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := newBuiltinMapOf[string, string](tc.m)
			keys := m.Keys()
			assert.ElementsMatch(t, tc.wantKeys, keys)
		})
	}
}

func TestBuiltinMap_Values(t *testing.T) {
	testCases := []struct {
		name       string
		m          map[string]string
		wantValues []string
	}{
		{
			name:       "nil",
			wantValues: []string{},
		},
		{
			name:       "empty",
			m:          map[string]string{},
			wantValues: []string{},
		},
		{
			name: "values",
			m: map[string]string{
				"book-1": "nature",
				"book-2": "world",
				"book-3": "peace",
			},
			wantValues: []string{"nature", "world", "peace"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := newBuiltinMapOf[string, string](tc.m)
			values := m.Values()
			assert.ElementsMatch(t, tc.wantValues, values)
		})
	}
}

func newBuiltinMapOf[K comparable, V any](data map[K]V) *builtinMap[K, V] {
	return &builtinMap[K, V]{data: data}
}
