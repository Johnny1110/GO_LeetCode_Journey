package top_k_frequent_elements

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "LeetCode Example 3",
			nums:     []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "NeetCode",
			nums:     []int{1, 2, 2, 3, 3, 3},
			k:        2,
			expected: []int{2, 3},
		},
		{
			name:     "Duplicate elements",
			nums:     []int{7, 7},
			k:        1,
			expected: []int{7},
		},
		{
			name:     "All elements same frequency",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{1, 2, 3}, // Any 3 elements are valid
		},
		{
			name:     "Single",
			nums:     []int{4, 1, -1, 2, -1, 2, 3},
			k:        2,
			expected: []int{-1, 2},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -2, -2, -2, -3},
			k:        2,
			expected: []int{-2, -1},
		},
		{
			name:     "Large numbers",
			nums:     []int{10000, 10000, 10001, 10001, 10001},
			k:        1,
			expected: []int{10001},
		},
		{
			name:     "Mixed frequencies",
			nums:     []int{1, 1, 1, 2, 2, 3, 3, 3, 3},
			k:        2,
			expected: []int{3, 1},
		},
		{
			name:     "K equals unique elements",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Zero",
			nums:     []int{0, 0, 1, 1, 1, 2},
			k:        2,
			expected: []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)

			// For cases where multiple valid answers exist (same frequency)
			if tt.name == "All elements same frequency" || tt.name == "K equals unique elements" {
				if len(result) != len(tt.expected) {
					t.Errorf("topKFrequent(%v, %d) = %v; want length %d",
						tt.nums, tt.k, result, len(tt.expected))
					return
				}
				// Verify all returned elements are from the original array
				if !areValidElements(result, tt.nums) {
					t.Errorf("topKFrequent(%v, %d) = %v; contains invalid elements",
						tt.nums, tt.k, result)
				}
				return
			}

			// Sort both slices for comparison since order may vary
			sort.Ints(result)
			expectedSorted := make([]int, len(tt.expected))
			copy(expectedSorted, tt.expected)
			sort.Ints(expectedSorted)

			if !reflect.DeepEqual(result, expectedSorted) {
				t.Errorf("topKFrequent(%v, %d) = %v; want %v",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func TestTopKFrequentEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Single element array",
			nums:     []int{42},
			k:        1,
			expected: []int{42},
		},
		{
			name:     "Two elements equal frequency",
			nums:     []int{1, 2},
			k:        1,
			expected: []int{1}, // Any one element is valid
		},
		{
			name:     "Large array single frequency",
			nums:     []int{5, 5, 5, 5, 5, 5, 5},
			k:        1,
			expected: []int{5},
		},
		{
			name:     "Descending frequency pattern",
			nums:     []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4},
			k:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "All different elements",
			nums:     []int{10, 20, 30, 40, 50},
			k:        2,
			expected: []int{10, 20}, // Any 2 elements are valid
		},
		{
			name:     "Large k",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
			k:        4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Complex frequency distribution",
			nums:     []int{1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3},
			k:        2,
			expected: []int{3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)

			// For cases where multiple valid answers exist
			if tt.name == "Two elements equal frequency" || tt.name == "All different elements" {
				if len(result) != len(tt.expected) {
					t.Errorf("topKFrequent(%v, %d) = %v; want length %d",
						tt.nums, tt.k, result, len(tt.expected))
					return
				}
				if !areValidElements(result, tt.nums) {
					t.Errorf("topKFrequent(%v, %d) = %v; contains invalid elements",
						tt.nums, tt.k, result)
				}
				return
			}

			// Sort both slices for comparison
			sort.Ints(result)
			expectedSorted := make([]int, len(tt.expected))
			copy(expectedSorted, tt.expected)
			sort.Ints(expectedSorted)

			if !reflect.DeepEqual(result, expectedSorted) {
				t.Errorf("topKFrequent(%v, %d) = %v; want %v",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func TestTopKFrequentProperties(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
	}{
		{"Property test 1", []int{1, 1, 1, 2, 2, 3}, 2},
		{"Property test 2", []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 3},
		{"Property test 3", []int{4, 1, -1, 2, -1, 2, 3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)

			// Test properties that should always hold

			// 1. Result should have exactly k elements
			if len(result) != tt.k {
				t.Errorf("Expected %d elements, got %d", tt.k, len(result))
			}

			// 2. All returned elements should exist in original array
			if !areValidElements(result, tt.nums) {
				t.Errorf("Result contains elements not in original array")
			}

			// 3. No duplicates in result
			if hasDuplicates(result) {
				t.Errorf("Result contains duplicate elements: %v", result)
			}

			// 4. Verify frequency ordering (most frequent elements should be returned)
			if !verifyFrequencyOrder(result, tt.nums, tt.k) {
				t.Errorf("Result does not contain the most frequent elements")
			}
		})
	}
}

// Helper function to check if all elements in result exist in original array
func areValidElements(result, original []int) bool {
	originalSet := make(map[int]bool)
	for _, num := range original {
		originalSet[num] = true
	}

	for _, num := range result {
		if !originalSet[num] {
			return false
		}
	}

	return true
}

// Helper function to check for duplicates in result
func hasDuplicates(result []int) bool {
	seen := make(map[int]bool)
	for _, num := range result {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

// Helper function to verify the result contains the most frequent elements
func verifyFrequencyOrder(result, nums []int, k int) bool {
	// Count frequencies
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Get all unique elements sorted by frequency (descending)
	type freqPair struct {
		num  int
		freq int
	}

	var pairs []freqPair
	for num, freq := range freqMap {
		pairs = append(pairs, freqPair{num, freq})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	// Get the top k frequencies
	topKFreqs := make([]int, 0, k)
	for i := 0; i < k && i < len(pairs); i++ {
		topKFreqs = append(topKFreqs, pairs[i].freq)
	}

	// Verify each result element has a frequency in top k
	for _, num := range result {
		freq := freqMap[num]
		found := false
		for _, topFreq := range topKFreqs {
			if freq >= topFreq {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func BenchmarkTopKFrequent(b *testing.B) {
	nums := []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 5, 5, 6, 7, 8, 9, 10}
	k := 3

	for i := 0; i < b.N; i++ {
		topKFrequent(nums, k)
	}
}

func BenchmarkTopKFrequentLarge(b *testing.B) {
	// Create a larger test case
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i % 100 // Create frequency patterns
	}
	k := 10

	for i := 0; i < b.N; i++ {
		topKFrequent(nums, k)
	}
}

func BenchmarkTopKFrequentWorstCase(b *testing.B) {
	// All elements have same frequency
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	k := 5

	for i := 0; i < b.N; i++ {
		topKFrequent(nums, k)
	}
}
