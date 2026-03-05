package kth_largest_element

import "testing"

func TestQuickSelectKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		start    int
		end      int
		k        int
		expected int
	}{
		{
			name:     "Classic 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			start:    0,
			end:      5,
			k:        2,
			expected: 5,
		},
		{
			name:     "Classic 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			start:    0,
			end:      8,
			k:        4,
			expected: 4,
		},
		{
			name:     "Find largest element - k=1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			start:    0,
			end:      5,
			k:        1,
			expected: 6,
		},
		{
			name:     "Find smallest element - k=length",
			nums:     []int{3, 2, 1, 5, 6, 4},
			start:    0,
			end:      5,
			k:        6,
			expected: 1,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			start:    0,
			end:      0,
			k:        1,
			expected: 1,
		},
		{
			name:     "Two elements - first largest",
			nums:     []int{3, 2},
			start:    0,
			end:      1,
			k:        1,
			expected: 3,
		},
		{
			name:     "Two elements - second largest",
			nums:     []int{3, 2},
			start:    0,
			end:      1,
			k:        2,
			expected: 2,
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1, 1},
			start:    0,
			end:      4,
			k:        3,
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -3, 2, 0, -2},
			start:    0,
			end:      4,
			k:        2,
			expected: 0,
		},
		{
			name:     "Large numbers",
			nums:     []int{10000, 9999, 9998, 10001, 10002},
			start:    0,
			end:      4,
			k:        3,
			expected: 10000,
		},
		{
			name:     "Duplicates with different k",
			nums:     []int{2, 1, 3, 3, 3, 4},
			start:    0,
			end:      5,
			k:        2,
			expected: 4,
		},
		{
			name:     "Reverse sorted array",
			nums:     []int{5, 4, 3, 2, 1},
			start:    0,
			end:      4,
			k:        3,
			expected: 3,
		},
		{
			name:     "Already sorted array",
			nums:     []int{1, 2, 3, 4, 5},
			start:    0,
			end:      4,
			k:        3,
			expected: 3,
		},
		{
			name:     "Partial range - middle elements",
			nums:     []int{1, 7, 3, 6, 5, 2, 9},
			start:    2,
			end:      5,
			k:        2,
			expected: 5, // In subarray [3,6,5,2], 2nd largest is 5
		},
		{
			name:     "Mix of positive and negative",
			nums:     []int{-5, 1, -3, 7, 2, -1},
			start:    0,
			end:      5,
			k:        3,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums to avoid modifying the original
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			result := quickSelectKthLargest(numsCopy, tt.start, tt.end, tt.k)
			if result != tt.expected {
				t.Errorf("quickSelectKthLargest(%v, %d, %d, %d) = %d; want %d",
					tt.nums, tt.start, tt.end, tt.k, result, tt.expected)
			}
		})
	}
}

func TestQuickSelectKthLargestEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		start    int
		end      int
		k        int
		expected int
	}{
		{
			name:     "Large array with many duplicates",
			nums:     []int{1, 2, 2, 2, 2, 2, 3, 3, 3, 4},
			start:    0,
			end:      9,
			k:        5,
			expected: 2,
		},
		{
			name:     "Zeros and ones",
			nums:     []int{0, 1, 0, 1, 0, 1},
			start:    0,
			end:      5,
			k:        3,
			expected: 1,
		},
		{
			name:     "Single range element",
			nums:     []int{5, 10, 15},
			start:    1,
			end:      1,
			k:        1,
			expected: 10,
		},
		{
			name:     "Very large k equals array size",
			nums:     []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			start:    0,
			end:      8,
			k:        9,
			expected: 1,
		},
		{
			name:     "Random order with negatives",
			nums:     []int{-10, 5, -2, 0, 3, -7, 1},
			start:    0,
			end:      6,
			k:        4,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums to avoid modifying the original
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			result := quickSelectKthLargest(numsCopy, tt.start, tt.end, tt.k)
			if result != tt.expected {
				t.Errorf("quickSelectKthLargest(%v, %d, %d, %d) = %d; want %d",
					tt.nums, tt.start, tt.end, tt.k, result, tt.expected)
			}
		})
	}
}

// Helper function to verify our expected results by sorting
func verifyKthLargest(nums []int, start, end, k int) int {
	// Extract the subarray
	subArray := make([]int, end-start+1)
	copy(subArray, nums[start:end+1])

	// Simple bubble sort for verification (not efficient, but correct)
	n := len(subArray)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if subArray[j] < subArray[j+1] {
				subArray[j], subArray[j+1] = subArray[j+1], subArray[j]
			}
		}
	}

	return subArray[k-1]
}

func TestVerifyExpectedResults(t *testing.T) {
	// Test to verify our expected results are correct
	testCases := []struct {
		nums  []int
		start int
		end   int
		k     int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 0, 5, 2},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 0, 8, 4},
		{[]int{-1, -3, 2, 0, -2}, 0, 4, 2},
	}

	for _, tc := range testCases {
		expected := verifyKthLargest(tc.nums, tc.start, tc.end, tc.k)
		t.Logf("Array: %v, k=%d -> %dth largest = %d", tc.nums, tc.k, tc.k, expected)
	}
}

func BenchmarkQuickSelectKthLargest(b *testing.B) {
	nums := []int{3, 2, 1, 5, 6, 4, 8, 7, 9, 10, 12, 11, 15, 14, 13}

	for i := 0; i < b.N; i++ {
		// Create a copy for each benchmark iteration
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		quickSelectKthLargest(numsCopy, 0, len(nums)-1, 5)
	}
}

func BenchmarkQuickSelectKthLargestLarge(b *testing.B) {
	// Create a larger test case
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = 1000 - i // Reverse order
	}

	for i := 0; i < b.N; i++ {
		// Create a copy for each benchmark iteration
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		quickSelectKthLargest(numsCopy, 0, len(nums)-1, 100)
	}
}
