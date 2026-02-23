package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: []int{1},
		},
		{
			name:     "Two elements",
			nums:     []int{2, 3},
			expected: []int{3, 2},
		},
		{
			name:     "Contains zero at start",
			nums:     []int{0, 1, 2, 3},
			expected: []int{6, 0, 0, 0},
		},
		{
			name:     "Contains zero at end",
			nums:     []int{1, 2, 3, 0},
			expected: []int{0, 0, 0, 6},
		},
		{
			name:     "Contains zero in middle",
			nums:     []int{1, 0, 3, 4},
			expected: []int{0, 12, 0, 0},
		},
		{
			name:     "All negative",
			nums:     []int{-1, -2, -3},
			expected: []int{6, 3, 2},
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-2, 3, -4},
			expected: []int{-12, 8, -6},
		},
		{
			name:     "Large numbers",
			nums:     []int{10, 20, 30, 40},
			expected: []int{24000, 12000, 8000, 6000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestProductExceptSelfEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Multiple zeros",
			nums:     []int{0, 0, 2, 3},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "One is present",
			nums:     []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		productExceptSelf(nums)
	}
}
