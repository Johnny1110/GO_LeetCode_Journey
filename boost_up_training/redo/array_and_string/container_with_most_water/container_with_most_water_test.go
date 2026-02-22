package container_with_most_water

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1: [1,8,6,2,5,4,8,3,7]",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2: [1,1]",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "Decreasing heights",
			height:   []int{5, 4, 3, 2, 1},
			expected: 6,
		},
		{
			name:     "Increasing heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "Same heights",
			height:   []int{3, 3, 3, 3, 3},
			expected: 12,
		},
		{
			name:     "Tallest at both ends",
			height:   []int{10, 1, 1, 1, 10},
			expected: 40,
		},
		{
			name:     "Tallest in the middle",
			height:   []int{1, 1, 10, 10, 1, 1},
			expected: 10,
		},
		{
			name:     "Two elements - different heights",
			height:   []int{1, 100},
			expected: 1,
		},
		{
			name:     "Large gap with small walls",
			height:   []int{2, 0, 0, 0, 0, 0, 0, 0, 2},
			expected: 16,
		},
		{
			name:     "One very tall line",
			height:   []int{1, 1, 1, 1, 100},
			expected: 4,
		},
		{
			name:     "Symmetric V-shape",
			height:   []int{5, 3, 1, 3, 5},
			expected: 20,
		},
		{
			name:     "All ones",
			height:   []int{1, 1, 1, 1, 1, 1},
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxArea(tc.height)
			if result != tc.expected {
				t.Errorf("maxArea(%v) = %d, expected %d", tc.height, result, tc.expected)
			}
		})
	}
}
