package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		wantRes int
	}{
		{
			name:    "single val",
			src:     []int{1},
			wantRes: 1,
		},
		{
			name:    "multiple values",
			src:     []int{1, 2, 3},
			wantRes: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Max[int](tc.src)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		wantRes int
	}{
		{
			name:    "single val",
			src:     []int{1},
			wantRes: 1,
		},
		{
			name:    "multiple values",
			src:     []int{1, 2, 3},
			wantRes: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Min[int](tc.src)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		wantRes int
	}{
		{
			name: "nil",
		},
		{
			name: "empty",
			src:  []int{},
		},
		{
			name:    "single val",
			src:     []int{1},
			wantRes: 1,
		},
		{
			name:    "multiple values",
			src:     []int{1, 2, 3},
			wantRes: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum[int](tc.src)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
