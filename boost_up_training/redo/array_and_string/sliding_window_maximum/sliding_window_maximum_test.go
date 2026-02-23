package sliding_window_maximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "Example 2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "k equals array length",
			nums:     []int{1, 3, -1, -3, 5},
			k:        5,
			expected: []int{5},
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{9, 7, 5, 3, 1},
			k:        3,
			expected: []int{9, 7, 5},
		},
		{
			name:     "Increasing sequence",
			nums:     []int{1, 3, 5, 7, 9},
			k:        3,
			expected: []int{5, 7, 9},
		},
		{
			name:     "All same elements",
			nums:     []int{5, 5, 5, 5, 5},
			k:        2,
			expected: []int{5, 5, 5, 5},
		},
		{
			name:     "Window size 1",
			nums:     []int{1, -1, 0, 5, -3},
			k:        1,
			expected: []int{1, -1, 0, 5, -3},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -3, -2, -4, -5},
			k:        2,
			expected: []int{-1, -2, -2, -4},
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			k:        4,
			expected: []int{4, 4, 4, 4, 2, 4},
		},
		{
			name:     "Two elements window size 2",
			nums:     []int{1, 2},
			k:        2,
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func TestMaxSlidingWindowEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Large numbers",
			nums:     []int{1000, 2000, 3000, 1500, 2500},
			k:        3,
			expected: []int{3000, 3000, 3000},
		},
		{
			name:     "Zero included",
			nums:     []int{0, 1, 0, 2, 0},
			k:        2,
			expected: []int{1, 1, 2, 2},
		},
		{
			name:     "Peak in middle",
			nums:     []int{1, 2, 10, 2, 1},
			k:        3,
			expected: []int{10, 10, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func BenchmarkMaxSlidingWindow(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i * 2
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSlidingWindow(nums, k)
	}
}

func BenchmarkMaxSlidingWindowLargeK(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i % 100
	}
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSlidingWindow(nums, k)
	}
}
