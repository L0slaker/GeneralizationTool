package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiffSet(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     []int
		wantRes []int
	}{
		{
			name:    "diff 1",
			src:     []int{1, 2, 3, 4},
			dst:     []int{1, 3, 4},
			wantRes: []int{2},
		},
		{
			name:    "src less than dst",
			src:     []int{1, 2},
			dst:     []int{1, 2, 4},
			wantRes: []int{},
		},
		{
			name:    "diff deduplicate",
			src:     []int{1, 2, 3, 4, 5, 6, 6},
			dst:     []int{1, 3, 5},
			wantRes: []int{2, 4, 6},
		},
		{
			name:    "dst duplicate ele",
			src:     []int{1, 1, 3, 5, 7},
			dst:     []int{1, 3, 5, 5},
			wantRes: []int{7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := DiffSet[int](tc.src, tc.dst)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestDiffSetFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		dst     []int
		wantRes []int
	}{
		{
			name:    "diff 1",
			src:     []int{1, 2, 3, 4},
			dst:     []int{1, 3, 4},
			wantRes: []int{2},
		},
		{
			name:    "src less than dst",
			src:     []int{1, 2},
			dst:     []int{1, 2, 4},
			wantRes: []int{},
		},
		{
			name:    "diff deduplicate",
			src:     []int{1, 2, 3, 4, 5, 6, 6},
			dst:     []int{1, 3, 5},
			wantRes: []int{2, 4, 6},
		},
		{
			name:    "dst duplicate ele",
			src:     []int{1, 1, 3, 5, 7},
			dst:     []int{1, 3, 5, 5},
			wantRes: []int{7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := DiffSetFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
