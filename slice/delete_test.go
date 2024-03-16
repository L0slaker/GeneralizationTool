package slice

import (
	"fmt"
	"generalization_tool/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		index   int
		wantRes []int
		wantErr error
	}{
		{
			name:    "index 0",
			src:     []int{1, 2, 3},
			index:   0,
			wantRes: []int{2, 3},
		},
		{
			name:    "index -1",
			src:     []int{1, 2, 3},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(3, -1),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Delete[int](tc.src, tc.index)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestFilterDelete(t *testing.T) {
	testCases := []struct {
		name            string
		src             []int
		deleteCondition func(index, src int) bool
		wantRes         []int
		capacity        int
	}{
		{
			name: "空切片",
			src:  []int{},
			deleteCondition: func(index, src int) bool {
				return false
			},
			wantRes: []int{},
		},
		{
			name: "不删除元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return false
			},
			wantRes:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			capacity: 8,
		},
		{
			name: "删除首位元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return index == 0
			},
			wantRes:  []int{2, 3, 4, 5, 6, 7, 8},
			capacity: 8,
		},
		{
			name: "删除前面两个元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return index == 0 || index == 1
			},
			wantRes:  []int{3, 4, 5, 6, 7, 8},
			capacity: 8,
		},
		{
			name: "删除中间单个元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return index == 3
			},
			wantRes:  []int{1, 2, 3, 5, 6, 7, 8},
			capacity: 8,
		},
		{
			name: "删除中间多个不连续元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return index == 2 || index == 4 || index == 6
			},
			wantRes:  []int{1, 2, 4, 6, 8},
			capacity: 8,
		},
		{
			name: "删除中间多个连续元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return index == 3 || index == 4 || index == 5
			},
			wantRes:  []int{1, 2, 3, 7, 8},
			capacity: 8,
		},
		{
			name: "删除中间多个元素，第一部分为一个元素，第二部分为连续元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return index == 2 || index == 4 || index == 5
			},
			wantRes:  []int{1, 2, 4, 7, 8},
			capacity: 8,
		},
		{
			name: "删除中间多个元素，第一部分为连续元素，第二部分为一个元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return index == 2 || index == 3 || index == 5
			},
			wantRes:  []int{1, 2, 5, 7, 8},
			capacity: 8,
		},
		{
			name: "删除后面两个元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return index == 6 || index == 7
			},
			wantRes:  []int{1, 2, 3, 4, 5, 6},
			capacity: 8,
		},
		{
			name: "删除末尾元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return index == 7
			},
			wantRes:  []int{1, 2, 3, 4, 5, 6, 7},
			capacity: 8,
		},
		{
			name: "删除所有元素",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			deleteCondition: func(index, src int) bool {
				return true
			},
			wantRes:  []int{},
			capacity: 8,
		},
		{
			name: "容量大于256，触发缩容",
			src:  testCase(),
			deleteCondition: func(index, src int) bool {
				return index >= 0 && index <= 277
			},
			wantRes:  []int{278, 279},
			capacity: 140,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterDelete[int](tc.src, tc.deleteCondition)
			fmt.Println("tc.wantRes:", tc.wantRes)
			fmt.Println("res:", res)
			assert.Equal(t, tc.wantRes, res)
			fmt.Println("tc.capacity:", tc.capacity)
			fmt.Println("cap(res):", cap(res))
			assert.Equal(t, tc.capacity, cap(res))
		})
	}
}

func testCase() []int {
	res := make([]int, 0, 280)
	for i := 0; i < 280; i++ {
		res = append(res, i)
	}
	return res
}
