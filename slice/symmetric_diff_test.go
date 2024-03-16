package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymmetricDiffSet(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		res  []int
	}{
		{
			name: "deduplicate",
			src:  []int{1, 1, 2, 3, 4},
			dst:  []int{4, 5, 6, 7, 6},
			res:  []int{1, 2, 3, 5, 6, 7},
		},
		{
			name: "src length is 0",
			src:  []int{},
			dst:  []int{},
			res:  []int{},
		},
		{
			name: "not exist sam ele",
			src:  []int{1, 2, 3, 4},
			dst:  []int{5, 6, 7, 8},
			res:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "both nil",
			res:  []int{},
		},
		{
			name: "normal",
			src:  []int{1, 2, 3, 4},
			dst:  []int{3, 4, 5, 6},
			res:  []int{1, 2, 5, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SymmetricDiffSet[int](tc.src, tc.dst)
			assert.Equal(t, tc.res, res)
		})
	}
}

func TestSymmetricDiffSetFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		res  []int
	}{
		{
			name: "deduplicate",
			src:  []int{1, 1, 2, 3, 4},
			dst:  []int{4, 5, 6, 7, 6},
			res:  []int{1, 2, 3, 5, 6, 7},
		},
		{
			name: "src length is 0",
			src:  []int{},
			dst:  []int{},
			res:  []int{},
		},
		{
			name: "not exist sam ele",
			src:  []int{1, 2, 3, 4},
			dst:  []int{5, 6, 7, 8},
			res:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "both nil",
			res:  []int{},
		},
		{
			name: "normal",
			src:  []int{1, 2, 3, 4},
			dst:  []int{3, 4, 5, 6},
			res:  []int{1, 2, 5, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SymmetricDiffSetFunc[int](tc.src, tc.dst, func(src int, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.res, res)
		})
	}
}
