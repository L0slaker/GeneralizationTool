package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionSet(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		res  []int
	}{
		{
			name: "not empty",
			src:  []int{1, 2, 3, 4},
			dst:  []int{4, 5, 6, 1},
			res:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "src is empty",
			src:  []int{},
			dst:  []int{1, 3},
			res:  []int{1, 3},
		},
		{
			name: "dst is empty",
			src:  []int{1, 3},
			dst:  []int{},
			res:  []int{1, 3},
		},
		{
			name: "both empty",
			src:  []int{},
			dst:  []int{},
			res:  []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := UnionSet[int](tc.src, tc.dst)
			assert.Equal(t, tc.res, res)
		})
	}
}

func TestUnionSetFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		res  []int
	}{
		{
			name: "not empty",
			src:  []int{1, 2, 3, 4},
			dst:  []int{4, 5, 6, 1},
			res:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "src is empty",
			src:  []int{},
			dst:  []int{1, 3},
			res:  []int{1, 3},
		},
		{
			name: "dst is empty",
			src:  []int{1, 3},
			dst:  []int{},
			res:  []int{1, 3},
		},
		{
			name: "both empty",
			src:  []int{},
			dst:  []int{},
			res:  []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := UnionSetFunc[int](tc.src, tc.dst, func(src int, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.res, res)
		})
	}
}
