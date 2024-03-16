package mapx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	testCases := []struct {
		name    string
		input   map[int]string
		wantRes []int
	}{
		{
			name:    "nil",
			input:   nil,
			wantRes: []int{},
		},
		{
			name:    "empty",
			input:   map[int]string{},
			wantRes: []int{},
		},
		{
			name: "single",
			input: map[int]string{
				1: "lisa",
			},
			wantRes: []int{1},
		},
		{
			name: "multiple",
			input: map[int]string{
				1: "mike",
				2: "jane",
				3: "kanye",
			},
			wantRes: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Keys[int, string](tc.input)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestValues(t *testing.T) {
	testCases := []struct {
		name    string
		input   map[int]string
		wantRes []string
	}{
		{
			name:    "nil",
			input:   nil,
			wantRes: []string{},
		},
		{
			name:    "empty",
			input:   map[int]string{},
			wantRes: []string{},
		},
		{
			name: "single",
			input: map[int]string{
				1: "lisa",
			},
			wantRes: []string{"lisa"},
		},
		{
			name: "multiple",
			input: map[int]string{
				1: "mike",
				2: "jane",
				3: "kanye",
			},
			wantRes: []string{"mike", "jane", "kanye"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Values[int, string](tc.input)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestKeysValues(t *testing.T) {
	testCases := []struct {
		name       string
		input      map[int]string
		wantKeys   []int
		wantValues []string
	}{
		{
			name:       "nil",
			input:      nil,
			wantKeys:   []int{},
			wantValues: []string{},
		},
		{
			name:       "empty",
			input:      map[int]string{},
			wantKeys:   []int{},
			wantValues: []string{},
		},
		{
			name: "single",
			input: map[int]string{
				1: "lisa",
			},
			wantKeys:   []int{1},
			wantValues: []string{"lisa"},
		},
		{
			name: "multiple",
			input: map[int]string{
				1: "mike",
				2: "jane",
				3: "kanye",
			},
			wantKeys:   []int{1, 2, 3},
			wantValues: []string{"mike", "jane", "kanye"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keys, values := KeysValues[int, string](tc.input)
			assert.Equal(t, tc.wantKeys, keys)
			assert.Equal(t, tc.wantValues, values)
		})
	}
}
