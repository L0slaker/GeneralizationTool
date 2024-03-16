package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		wantRes []int
	}{
		{
			name:    "length of src is 0",
			src:     []int{},
			wantRes: []int{},
		},
		{
			name:    "src nil",
			src:     nil,
			wantRes: []int{},
		},
		{
			name:    "normal",
			src:     []int{1, 2, 3, 4},
			wantRes: []int{4, 3, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Reverse[int](tc.src)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestReverseSelf(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		wantRes []int
	}{
		{
			name:    "length of src is 0",
			src:     []int{},
			wantRes: []int{},
		},
		{
			name:    "src nil",
			src:     nil,
			wantRes: nil,
		},
		{
			name:    "normal",
			src:     []int{1, 2, 3, 4},
			wantRes: []int{4, 3, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ReverseSelf[int](tc.src)
			assert.Equal(t, tc.wantRes, tc.src)
		})
	}
}
