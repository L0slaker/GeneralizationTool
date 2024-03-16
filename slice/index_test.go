package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     int
		wantRes int
	}{
		{
			name:    "first one",
			src:     []int{1, 2, 3, 4, 5},
			dst:     1,
			wantRes: 0,
		},
		{
			name:    "the length of src is 0",
			src:     []int{},
			dst:     1,
			wantRes: -1,
		},
		{
			name:    "src nil",
			src:     nil,
			dst:     1,
			wantRes: -1,
		},
		{
			name:    "dst not exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     7,
			wantRes: -1,
		},
		{
			name:    "last one",
			src:     []int{1, 2, 3, 4, 5},
			dst:     5,
			wantRes: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Index[int](tc.src, tc.dst)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIndexFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     int
		wantRes int
	}{
		{
			name:    "first one",
			src:     []int{1, 2, 3, 4, 5},
			dst:     1,
			wantRes: 0,
		},
		{
			name:    "the length of src is 0",
			src:     []int{},
			dst:     1,
			wantRes: -1,
		},
		{
			name:    "src nil",
			src:     nil,
			dst:     1,
			wantRes: -1,
		},
		{
			name:    "dst not exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     7,
			wantRes: -1,
		},
		{
			name:    "last one",
			src:     []int{1, 2, 3, 4, 5},
			dst:     5,
			wantRes: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestLastIndex(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     int
		wantRes int
	}{
		{
			name:    "first one",
			src:     []int{1, 1, 3, 4, 5},
			dst:     1,
			wantRes: 1,
		},
		{
			name:    "the length of src is 0",
			src:     []int{},
			dst:     1,
			wantRes: -1,
		},
		{
			name:    "src nil",
			src:     nil,
			dst:     1,
			wantRes: -1,
		},
		{
			name:    "dst not exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     7,
			wantRes: -1,
		},
		{
			name:    "last one",
			src:     []int{1, 2, 3, 4, 5, 1},
			dst:     1,
			wantRes: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := LastIndex[int](tc.src, tc.dst)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestLastIndexFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     int
		wantRes int
	}{
		{
			name:    "first one",
			src:     []int{1, 1, 3, 4, 5},
			dst:     1,
			wantRes: 1,
		},
		{
			name:    "the length of src is 0",
			src:     []int{},
			dst:     1,
			wantRes: -1,
		},
		{
			name:    "src nil",
			src:     nil,
			dst:     1,
			wantRes: -1,
		},
		{
			name:    "dst not exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     7,
			wantRes: -1,
		},
		{
			name:    "last one",
			src:     []int{1, 2, 3, 4, 5, 1},
			dst:     1,
			wantRes: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := LastIndexFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIndexAll(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     int
		wantRes []int
	}{
		{
			name:    "continuous two",
			src:     []int{1, 1, 3, 4, 5},
			dst:     1,
			wantRes: []int{0, 1},
		},
		{
			name:    "the length of src is 0",
			src:     []int{},
			dst:     1,
			wantRes: []int{},
		},
		{
			name:    "src nil",
			src:     nil,
			dst:     1,
			wantRes: []int{},
		},
		{
			name:    "dst not exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     7,
			wantRes: []int{},
		},
		{
			name:    "discontinuous two",
			src:     []int{1, 2, 3, 4, 5, 1},
			dst:     1,
			wantRes: []int{0, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexAll[int](tc.src, tc.dst)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIndexAllFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     int
		wantRes []int
	}{
		{
			name:    "continuous two",
			src:     []int{1, 1, 3, 4, 5},
			dst:     1,
			wantRes: []int{0, 1},
		},
		{
			name:    "the length of src is 0",
			src:     []int{},
			dst:     1,
			wantRes: []int{},
		},
		{
			name:    "src nil",
			src:     nil,
			dst:     1,
			wantRes: []int{},
		},
		{
			name:    "dst not exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     7,
			wantRes: []int{},
		},
		{
			name:    "discontinuous two",
			src:     []int{1, 2, 3, 4, 5, 1},
			dst:     1,
			wantRes: []int{0, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexAllFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
