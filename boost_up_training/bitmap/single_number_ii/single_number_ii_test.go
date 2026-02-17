package single_number_ii

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 2, 3, 2},
			expected: 3,
		},
		{
			name:     "Example 2", 
			nums:     []int{0, 1, 0, 1, 0, 1, 99},
			expected: 99,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-2, -2, -3, -2},
			expected: -3,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-4, -4, -4, 5, 5, 5, 7},
			expected: 7,
		},
		{
			name:     "Zero as single number",
			nums:     []int{1, 1, 1, 0, 2, 2, 2},
			expected: 0,
		},
		{
			name:     "Large numbers",
			nums:     []int{43, 16, 45, 89, 45, -2147483648, 45, 2147483646, -2147483647, -2147483648, 43, 2147483647, -2147483646, -2147483648, 89, -2147483646, 89, -2147483646, 2147483646, 16, -2147483647, 2147483646, -2147483647, 16, 43},
			expected: 2147483647,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := singleNumber(tt.nums)
			if result != tt.expected {
				t.Errorf("singleNumber(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkSingleNumber(t *testing.B) {
	nums := []int{2, 2, 3, 2}
	for i := 0; i < t.N; i++ {
		singleNumber(nums)
	}
}
