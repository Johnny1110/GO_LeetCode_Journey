package first_missing_positive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 0},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 4, -1, 1},
			expected: 2,
		},
		{
			name:     "Example 3",
			nums:     []int{7, 8, 9, 11, 12},
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 1,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 2,
		},
		{
			name:     "Single element: [0]",
			nums:     []int{0},
			expected: 1,
		},
		{
			name:     "Single element: [-1]",
			nums:     []int{-1},
			expected: 1,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3},
			expected: 1,
		},
		{
			name:     "Consecutive from 1: [1,2,3,4,5]",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "Duplicates: [1,1,1,1]",
			nums:     []int{1, 1, 1, 1},
			expected: 2,
		},
		{
			name:     "Mixed with duplicates: [1,2,2,3]",
			nums:     []int{1, 2, 2, 3},
			expected: 4,
		},
		{
			name:     "Large gap: [1,1000]",
			nums:     []int{1, 1000},
			expected: 2,
		},
		{
			name:     "Unordered: [3,1,2]",
			nums:     []int{3, 1, 2},
			expected: 4,
		},
		{
			name:     "Mixed positive and negative: [-1,4,2,1,9,10]",
			nums:     []int{-1, 4, 2, 1, 9, 10},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := firstMissingPositive(tc.nums)
			if result != tc.expected {
				t.Errorf("firstMissingPositive(%v) = %d, expected %d", tc.nums, result, tc.expected)
			}
		})
	}
}
