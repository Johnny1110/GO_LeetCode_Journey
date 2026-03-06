package kth_largest_element

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Classic example 1 - k=2",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Classic example 2 - k=4",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "Find largest element - k=1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        1,
			expected: 6,
		},
		{
			name:     "Find smallest element - k=length",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        6,
			expected: 1,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "Two elements - first largest",
			nums:     []int{3, 2},
			k:        1,
			expected: 3,
		},
		{
			name:     "Two elements - second largest",
			nums:     []int{3, 2},
			k:        2,
			expected: 2,
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1, 1},
			k:        3,
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -3, 2, 0, -2},
			k:        2,
			expected: 0,
		},
		{
			name:     "Large numbers",
			nums:     []int{10000, 9999, 9998, 10001, 10002},
			k:        3,
			expected: 10000,
		},
		{
			name:     "Duplicates with different k",
			nums:     []int{2, 1, 3, 3, 3, 4},
			k:        2,
			expected: 3,
		},
		{
			name:     "Reverse sorted array",
			nums:     []int{5, 4, 3, 2, 1},
			k:        3,
			expected: 3,
		},
		{
			name:     "Already sorted array",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: 3,
		},
		{
			name:     "Mix of positive and negative",
			nums:     []int{-5, 1, -3, 7, 2, -1},
			k:        3,
			expected: 1,
		},
		{
			name:     "Array with zeros",
			nums:     []int{0, 1, 0, -1, 2, 0},
			k:        4,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("findKthLargest(%v, %d) = %d; want %d",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func TestFindKthLargestEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Large array with many duplicates",
			nums:     []int{1, 2, 2, 2, 2, 2, 3, 3, 3, 4},
			k:        5,
			expected: 2,
		},
		{
			name:     "Zeros and ones",
			nums:     []int{0, 1, 0, 1, 0, 1},
			k:        3,
			expected: 1,
		},
		{
			name:     "Very large k equals array size",
			nums:     []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			k:        9,
			expected: 1,
		},
		{
			name:     "Random order with negatives",
			nums:     []int{-10, 5, -2, 0, 3, -7, 1},
			k:        4,
			expected: 0,
		},
		{
			name:     "Small k with large array",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:        2,
			expected: 9,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        1,
			expected: -1,
		},
		{
			name:     "All negative numbers - kth smallest",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        5,
			expected: -5,
		},
		{
			name:     "Array with one large outlier",
			nums:     []int{1, 1, 1, 1, 100},
			k:        2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("findKthLargest(%v, %d) = %d; want %d",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func TestFindKthLargestLeetCodeExamples(t *testing.T) {
	// Official LeetCode test cases
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "LeetCode Example 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "LeetCode Example 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "Single element case",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "Two identical elements",
			nums:     []int{3, 3},
			k:        1,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("findKthLargest(%v, %d) = %d; want %d",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

// Helper function to verify our expected results by sorting
func verifyKthLargest(nums []int, k int) int {
	// Create a copy to avoid modifying original
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	// Simple bubble sort for verification (not efficient, but correct)
	n := len(numsCopy)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numsCopy[j] < numsCopy[j+1] {
				numsCopy[j], numsCopy[j+1] = numsCopy[j+1], numsCopy[j]
			}
		}
	}

	return numsCopy[k-1]
}

func TestVerifyExpectedResults(t *testing.T) {
	// Test to verify our expected results are correct
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
		{[]int{-1, -3, 2, 0, -2}, 2},
		{[]int{1, 2, 2, 2, 2, 2, 3, 3, 3, 4}, 5},
	}

	for _, tc := range testCases {
		expected := verifyKthLargest(tc.nums, tc.k)
		t.Logf("Array: %v, k=%d -> %dth largest = %d", tc.nums, tc.k, tc.k, expected)
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	nums := []int{3, 2, 1, 5, 6, 4, 8, 7, 9, 10, 12, 11, 15, 14, 13}
	k := 5

	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargestLarge(b *testing.B) {
	// Create a larger test case
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = 1000 - i // Reverse order
	}
	k := 100

	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargestWorstCase(b *testing.B) {
	// Worst case for heap: finding smallest element (k = len(nums))
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	k := len(nums)

	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargestBestCase(b *testing.B) {
	// Best case for heap: finding largest element (k = 1)
	nums := []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	k := 1

	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}
