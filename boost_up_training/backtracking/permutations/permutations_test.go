package permutations

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to sort permutations for comparison
func sortPermutations(perms [][]int) {
	for _, perm := range perms {
		sort.Ints(perm)
	}
	sort.Slice(perms, func(i, j int) bool {
		for k := 0; k < len(perms[i]) && k < len(perms[j]); k++ {
			if perms[i][k] != perms[j][k] {
				return perms[i][k] < perms[j][k]
			}
		}
		return len(perms[i]) < len(perms[j])
	})
}

// Helper function to check if two sets of permutations are equal
func equalPermutations(actual, expected [][]int) bool {
	if len(actual) != len(expected) {
		return false
	}
	
	// Create copies to avoid modifying original slices
	actualCopy := make([][]int, len(actual))
	expectedCopy := make([][]int, len(expected))
	
	for i := range actual {
		actualCopy[i] = make([]int, len(actual[i]))
		copy(actualCopy[i], actual[i])
	}
	
	for i := range expected {
		expectedCopy[i] = make([]int, len(expected[i]))
		copy(expectedCopy[i], expected[i])
	}
	
	// Sort both for comparison
	sortPermutations(actualCopy)
	sortPermutations(expectedCopy)
	
	return reflect.DeepEqual(actualCopy, expectedCopy)
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "Example 1 - Three elements",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
				{2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name:     "Example 2 - Two elements",
			input:    []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			name:     "Example 3 - Single element",
			input:    []int{1},
			expected: [][]int{{1}},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: [][]int{{}},
		},
		{
			name:  "Four elements",
			input: []int{1, 2, 3, 4},
			expected: [][]int{
				{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2},
				{1, 4, 2, 3}, {1, 4, 3, 2}, {2, 1, 3, 4}, {2, 1, 4, 3},
				{2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1},
				{3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1},
				{3, 4, 1, 2}, {3, 4, 2, 1}, {4, 1, 2, 3}, {4, 1, 3, 2},
				{4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1},
			},
		},
		{
			name:     "Negative numbers",
			input:    []int{-1, 0, 1},
			expected: [][]int{{-1, 0, 1}, {-1, 1, 0}, {0, -1, 1}, {0, 1, -1}, {1, -1, 0}, {1, 0, -1}},
		},
		{
			name:     "Large numbers",
			input:    []int{10, 20},
			expected: [][]int{{10, 20}, {20, 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := permute(tt.input)
			
			if !equalPermutations(actual, tt.expected) {
				t.Errorf("permute(%v) = %v, expected %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestPermuteProperties(t *testing.T) {
	t.Run("Factorial count check", func(t *testing.T) {
		testCases := []struct {
			input    []int
			expected int
		}{
			{[]int{}, 1},         // 0! = 1
			{[]int{1}, 1},        // 1! = 1
			{[]int{1, 2}, 2},     // 2! = 2
			{[]int{1, 2, 3}, 6},  // 3! = 6
			{[]int{1, 2, 3, 4}, 24}, // 4! = 24
		}

		for _, tc := range testCases {
			result := permute(tc.input)
			if len(result) != tc.expected {
				t.Errorf("permute(%v) returned %d permutations, expected %d", tc.input, len(result), tc.expected)
			}
		}
	})

	t.Run("All permutations are unique", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := permute(input)
		
		seen := make(map[string]bool)
		for _, perm := range result {
			key := ""
			for _, num := range perm {
				key += string(rune(num + 1000)) // Convert to string for map key
			}
			if seen[key] {
				t.Errorf("Duplicate permutation found: %v", perm)
			}
			seen[key] = true
		}
	})

	t.Run("Each permutation has correct length", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		result := permute(input)
		
		for i, perm := range result {
			if len(perm) != len(input) {
				t.Errorf("Permutation %d has length %d, expected %d", i, len(perm), len(input))
			}
		}
	})

	t.Run("Each permutation contains all original elements", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := permute(input)
		
		for i, perm := range result {
			// Check if permutation contains all elements from input
			inputCopy := make([]int, len(input))
			copy(inputCopy, input)
			sort.Ints(inputCopy)
			
			permCopy := make([]int, len(perm))
			copy(permCopy, perm)
			sort.Ints(permCopy)
			
			if !reflect.DeepEqual(inputCopy, permCopy) {
				t.Errorf("Permutation %d (%v) doesn't contain all original elements (%v)", i, perm, input)
			}
		}
	})
}

func BenchmarkPermute(b *testing.B) {
	benchmarks := []struct {
		name  string
		input []int
	}{
		{"Small-3", []int{1, 2, 3}},
		{"Medium-4", []int{1, 2, 3, 4}},
		{"Large-5", []int{1, 2, 3, 4, 5}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				permute(bm.input)
			}
		})
	}
}
