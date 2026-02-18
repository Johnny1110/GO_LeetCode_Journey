package three_sum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Example 2: [0,1,1]",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "Example 3: [0,0,0]",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			expected: [][]int{},
		},
		{
			name:     "Three elements - no sum zero",
			nums:     []int{1, 2, 3},
			expected: [][]int{},
		},
		{
			name:     "Three elements - sum zero",
			nums:     []int{-1, 0, 1},
			expected: [][]int{{-1, 0, 1}},
		},
		{
			name:     "All positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: [][]int{},
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			expected: [][]int{},
		},
		{
			name:     "Multiple valid triplets",
			nums:     []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			expected: [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Duplicates with valid triplet",
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
		{
			name:     "Many duplicates",
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "Large numbers",
			nums:     []int{-100000, -99999, 199999},
			expected: [][]int{{-100000, -99999, 199999}},
		},
		{
			name:     "Single zero with positive and negative",
			nums:     []int{-1, 0, 1, 0},
			expected: [][]int{{-1, 0, 1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Make a copy since threeSum might modify the input
			numsCopy := make([]int, len(tc.nums))
			copy(numsCopy, tc.nums)

			result := threeSum(numsCopy)

			// Sort both result and expected for comparison
			sortTriplets(result)
			sortTriplets(tc.expected)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("threeSum(%v) = %v, expected %v", tc.nums, result, tc.expected)
			}
		})
	}
}

// Helper function to sort triplets for consistent comparison
func sortTriplets(triplets [][]int) {
	// Sort each triplet internally
	for _, triplet := range triplets {
		sort.Ints(triplet)
	}
	// Sort the array of triplets
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Unsorted array",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: []int{-4, -1, -1, 0, 1, 2},
		},
		{
			name:     "Already sorted",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			nums:     []int{42},
			expected: []int{42},
		},
		{
			name:     "Duplicates",
			nums:     []int{3, 1, 3, 1, 2},
			expected: []int{1, 1, 2, 3, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			numsCopy := make([]int, len(tc.nums))
			copy(numsCopy, tc.nums)

			if len(numsCopy) > 0 {
				quickSort(numsCopy, 0, len(numsCopy)-1)
			}

			if !reflect.DeepEqual(numsCopy, tc.expected) {
				t.Errorf("quickSort(%v) = %v, expected %v", tc.nums, numsCopy, tc.expected)
			}
		})
	}
}
