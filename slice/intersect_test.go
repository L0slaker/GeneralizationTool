package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     []int
		wantRes []int
	}{
		{
			name:    "continuous two",
			src:     []int{1, 3, 4, 5},
			dst:     []int{2, 4, 6, 8},
			wantRes: []int{4},
		},
		{
			name:    "the length of src is 0",
			src:     []int{},
			dst:     []int{2, 4, 6, 8},
			wantRes: []int{},
		},
		{
			name:    "src nil",
			src:     nil,
			dst:     []int{2, 4, 6, 8},
			wantRes: []int{},
		},
		{
			name:    "exist the same ele in src",
			src:     []int{1, 2, 3, 4, 5, 5},
			dst:     []int{1, 2, 3, 4, 5},
			wantRes: []int{1, 2, 3, 4, 5},
		},
		{
			name:    "dst empty",
			src:     []int{1, 2, 3, 4, 5, 5},
			dst:     []int{},
			wantRes: []int{},
		},
		{
			name:    "dst nil",
			src:     []int{1, 2, 3, 4, 5, 5},
			dst:     nil,
			wantRes: []int{},
		},
		{
			name:    "exist the same ele in src and dst",
			src:     []int{1, 1, 2, 3, 4, 5, 6},
			dst:     []int{1, 2, 3, 4, 4},
			wantRes: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Intersect[int](tc.src, tc.dst)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIntersectSetFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     []int
		wantRes []int
	}{
		{
			name:    "continuous two",
			src:     []int{1, 3, 4, 5},
			dst:     []int{2, 4, 6, 8},
			wantRes: []int{4},
		},
		{
			name:    "the length of src is 0",
			src:     []int{},
			dst:     []int{2, 4, 6, 8},
			wantRes: []int{},
		},
		{
			name:    "src nil",
			src:     nil,
			dst:     []int{2, 4, 6, 8},
			wantRes: []int{},
		},
		{
			name:    "exist the same ele in src",
			src:     []int{1, 2, 3, 4, 5, 5},
			dst:     []int{1, 2, 3, 4, 5},
			wantRes: []int{1, 2, 3, 4, 5},
		},
		{
			name:    "dst empty",
			src:     []int{1, 2, 3, 4, 5, 5},
			dst:     []int{},
			wantRes: []int{},
		},
		{
			name:    "dst nil",
			src:     []int{1, 2, 3, 4, 5, 5},
			dst:     nil,
			wantRes: []int{},
		},
		{
			name:    "exist the same ele in src and dst",
			src:     []int{1, 1, 2, 3, 4, 5, 6},
			dst:     []int{1, 2, 3, 4, 4},
			wantRes: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IntersectSetFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
