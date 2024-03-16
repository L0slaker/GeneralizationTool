package slice

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		wantRes []string
	}{
		{
			name:    "src nil",
			src:     nil,
			wantRes: []string{},
		},
		{
			name:    "src empty",
			src:     []int{},
			wantRes: []string{},
		},
		{
			name:    "src has element",
			src:     []int{1, 2, 3},
			wantRes: []string{"1", "2", "3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Map[int, string](tc.src, func(idx int, src int) string {
				return strconv.Itoa(src)
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestFilterMap(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		wantRes []string
	}{
		{
			name:    "src nil",
			src:     nil,
			wantRes: []string{},
		},
		{
			name:    "src empty",
			src:     []int{},
			wantRes: []string{},
		},
		{
			name:    "src has element",
			src:     []int{1, -2, 3},
			wantRes: []string{"1", "3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterMap[int, string](tc.src, func(idx int, src int) (string, bool) {
				return strconv.Itoa(src), src >= 0
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
