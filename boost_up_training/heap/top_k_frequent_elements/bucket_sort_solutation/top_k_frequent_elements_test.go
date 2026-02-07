package bucket_sort_solutation

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
			name:     "All same frequency",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{1, 2, 3}, // Any 3 elements are valid
		},
		{
			name:     "Multiple elements with same highest frequency",
			nums:     []int{1, 1, 2, 2, 3, 3},
			k:        2,
			expected: []int{1, 2}, // Any 2 of {1, 2, 3} are valid
		},
		{
			name:     "Single element repeated",
			nums:     []int{5, 5, 5, 5},
			k:        1,
			expected: []int{5},
		},
		{
			name:     "Large array with clear top k",
			nums:     []int{4, 1, -1, 2, -1, 2, 3, 4, 4, 4},
			k:        2,
			expected: []int{4, -1}, // 4 appears 4 times, -1 appears 2 times
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -1, -2, -2, -3},
			k:        2,
			expected: []int{-1, -2},
		},
		{
			name:     "Return all elements when k equals length",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)

			// Check if result length is correct
			if len(result) != tt.k {
				t.Errorf("Expected length %d, got %d", tt.k, len(result))
				return
			}

			// For cases where order doesn't matter or multiple valid answers exist,
			// we need to check if the result is valid rather than exact match
			if isValidTopK(tt.nums, result, tt.k) {
				return
			}

			// For deterministic cases, check exact match (sorted)
			sortedResult := make([]int, len(result))
			copy(sortedResult, result)
			sort.Ints(sortedResult)

			sortedExpected := make([]int, len(tt.expected))
			copy(sortedExpected, tt.expected)
			sort.Ints(sortedExpected)

			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Helper function to validate if the result is a valid top k frequent elements
func isValidTopK(nums []int, result []int, k int) bool {
	// Count frequencies
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Get frequencies of result elements
	resultFreqs := make([]int, len(result))
	for i, num := range result {
		resultFreqs[i] = freq[num]
	}

	// Check if all result elements have frequencies >= minimum frequency of top k
	allFreqs := make([]int, 0, len(freq))
	for _, f := range freq {
		allFreqs = append(allFreqs, f)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(allFreqs)))

	if len(allFreqs) < k {
		return false
	}

	minTopKFreq := allFreqs[k-1]

	for _, f := range resultFreqs {
		if f < minTopKFreq {
			return false
		}
	}

	return true
}

func TestTopKFrequentEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
	}{
		{
			name: "Minimum constraints",
			nums: []int{1},
			k:    1,
		},
		{
			name: "k equals array length",
			nums: []int{1, 2, 3, 4, 5},
			k:    5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)
			if len(result) != tt.k {
				t.Errorf("Expected length %d, got %d", tt.k, len(result))
			}

			if !isValidTopK(tt.nums, result, tt.k) {
				t.Errorf("Result %v is not a valid top %d frequent elements for input %v", result, tt.k, tt.nums)
			}
		})
	}
}

// Benchmark test
func BenchmarkTopKFrequent(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i % 100 // Create frequency distribution
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		topKFrequent(nums, 10)
	}
}
