package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringSliceToIntSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected []int
		hasError bool
	}{
		{
			input:    []string{"1", "2", "3"},
			expected: []int{1, 2, 3},
			hasError: false,
		},
		{
			input:    []string{},
			expected: []int{},
			hasError: false,
		},
		{
			input:    []string{"1", "invalid", "3"},
			expected: nil,
			hasError: true,
		},
	}

	for _, test := range tests {
		result, err := StringSliceToIntSlice(test.input)

		if test.hasError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, test.expected, result)
		}
	}
}

func TestIntSliceToInt32Slice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int32
	}{
		{
			input:    []int{1, 2, 3},
			expected: []int32{1, 2, 3},
		},
		{
			input:    []int{},
			expected: []int32{},
		},
		{
			input:    []int{-1, -2, -3},
			expected: []int32{-1, -2, -3},
		},
	}

	for _, test := range tests {
		result := IntSliceToInt32Slice(test.input)
		require.Equal(t, test.expected, result)
	}
}
