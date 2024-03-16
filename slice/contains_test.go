package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     int
		wantRes bool
	}{
		{
			name:    "dst exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     3,
			wantRes: true,
		},
		{
			name:    "dst not exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     6,
			wantRes: false,
		},
		{
			name:    "length of src is 0",
			src:     []int{},
			dst:     6,
			wantRes: false,
		},
		{
			name:    "src nil",
			dst:     6,
			wantRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Contains[int](tc.src, tc.dst)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     int
		wantRes bool
	}{
		{
			name:    "dst exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     3,
			wantRes: true,
		},
		{
			name:    "dst not exist",
			src:     []int{1, 2, 3, 4, 5},
			dst:     6,
			wantRes: false,
		},
		{
			name:    "length of src is 0",
			src:     []int{},
			dst:     6,
			wantRes: false,
		},
		{
			name:    "src nil",
			dst:     6,
			wantRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAny(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     []int
		wantRes bool
	}{
		{
			name:    "exist two elements",
			src:     []int{1, 2, 3, 4, 5, 6},
			dst:     []int{1, 3},
			wantRes: true,
		},
		{
			name:    "exist two same elements",
			src:     []int{1, 1, 2, 3, 4, 5, 6},
			dst:     []int{1, 1},
			wantRes: true,
		},
		{
			name:    "not exists the same",
			src:     []int{1, 2, 3, 4, 5, 6},
			dst:     []int{7, 8},
			wantRes: false,
		},
		{
			name:    "length of src is 0",
			src:     []int{},
			dst:     []int{7, 8},
			wantRes: false,
		},
		{
			name:    "src nil",
			dst:     []int{7, 8},
			wantRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAny[int](tc.src, tc.dst)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAnyFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     []int
		wantRes bool
	}{
		{
			name:    "exist two elements",
			src:     []int{1, 2, 3, 4, 5, 6},
			dst:     []int{1, 3},
			wantRes: true,
		},
		{
			name:    "exist two same elements",
			src:     []int{1, 1, 2, 3, 4, 5, 6},
			dst:     []int{1, 1},
			wantRes: true,
		},
		{
			name:    "not exists the same",
			src:     []int{1, 2, 3, 4, 5, 6},
			dst:     []int{7, 8},
			wantRes: false,
		},
		{
			name:    "length of src is 0",
			src:     []int{},
			dst:     []int{7, 8},
			wantRes: false,
		},
		{
			name:    "src nil",
			dst:     []int{7, 8},
			wantRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAnyFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAll(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     []int
		wantRes bool
	}{
		{
			name:    "src exist one not in dist",
			src:     []int{1, 2, 3, 4, 4},
			dst:     []int{1, 2, 3, 4},
			wantRes: true,
		},
		{
			name:    "src not include the whole elements",
			src:     []int{1, 2, 3, 4, 4},
			dst:     []int{1, 2, 3, 4, 4, 5},
			wantRes: false,
		},
		{
			name:    "length of src is 0",
			src:     []int{},
			dst:     []int{1},
			wantRes: false,
		},
		{
			name:    "src nil dst empty",
			src:     nil,
			dst:     []int{},
			wantRes: true,
		},
		{
			name:    "scr and dst nil",
			src:     nil,
			dst:     nil,
			wantRes: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAll[int](tc.src, tc.dst)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     []int
		wantRes bool
	}{
		{
			name:    "src exist one not in dist",
			src:     []int{1, 2, 3, 4, 4},
			dst:     []int{1, 2, 3, 4},
			wantRes: true,
		},
		{
			name:    "src not include the whole elements",
			src:     []int{1, 2, 3, 4, 4},
			dst:     []int{1, 2, 3, 4, 4, 5},
			wantRes: false,
		},
		{
			name:    "length of src is 0",
			src:     []int{},
			dst:     []int{1},
			wantRes: false,
		},
		{
			name:    "src nil dst empty",
			src:     nil,
			dst:     []int{},
			wantRes: true,
		},
		{
			name:    "scr and dst nil",
			src:     nil,
			dst:     nil,
			wantRes: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAllFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
