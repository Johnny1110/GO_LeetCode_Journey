package quick_sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "two elements unsorted",
			input:    []int{3, 1},
			expected: []int{1, 3},
		},
		{
			name:     "two elements sorted",
			input:    []int{1, 3},
			expected: []int{1, 3},
		},
		{
			name:     "basic unsorted array",
			input:    []int{3, 6, 8, 10, 1, 2, 1},
			expected: []int{1, 1, 2, 3, 6, 8, 10},
		},
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all same elements",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:     "negative numbers",
			input:    []int{-3, -1, -7, 2, -5},
			expected: []int{-7, -5, -3, -1, 2},
		},
		{
			name:     "large array",
			input:    []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42},
			expected: []int{11, 12, 22, 25, 34, 42, 50, 64, 76, 88, 90},
		},
		{
			name:     "example - 1",
			input:    []int{3, 6, 2, 7, 1, 4},
			expected: []int{1, 2, 3, 4, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying the original test data
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			quickSort(input)

			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("quickSort() = %v, want %v", input, tt.expected)
			}
		})
	}
}
