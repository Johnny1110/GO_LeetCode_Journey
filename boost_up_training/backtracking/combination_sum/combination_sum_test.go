package combination_sum

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to sort combinations for comparison
func sortCombinations(combinations [][]int) {
	for _, combination := range combinations {
		sort.Ints(combination)
	}
	sort.Slice(combinations, func(i, j int) bool {
		for k := 0; k < len(combinations[i]) && k < len(combinations[j]); k++ {
			if combinations[i][k] != combinations[j][k] {
				return combinations[i][k] < combinations[j][k]
			}
		}
		return len(combinations[i]) < len(combinations[j])
	})
}

// Helper function to check if two sets of combinations are equal
func equalCombinations(actual, expected [][]int) bool {
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
	sortCombinations(actualCopy)
	sortCombinations(expectedCopy)
	
	return reflect.DeepEqual(actualCopy, expectedCopy)
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Example 1 - Multiple solutions",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Example 2 - Multiple ways with same numbers",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Example 3 - No solution",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "Single element solution",
			candidates: []int{1},
			target:     1,
			expected:   [][]int{{1}},
		},
		{
			name:       "Single element multiple uses",
			candidates: []int{1},
			target:     3,
			expected:   [][]int{{1, 1, 1}},
		},
		{
			name:       "Exact match single candidate",
			candidates: []int{5},
			target:     5,
			expected:   [][]int{{5}},
		},
		{
			name:       "Target zero",
			candidates: []int{1, 2, 3},
			target:     0,
			expected:   [][]int{{}},
		},
		{
			name:       "Large candidates small target",
			candidates: []int{10, 15, 20},
			target:     5,
			expected:   [][]int{},
		},
		{
			name:       "Multiple same values allowed",
			candidates: []int{2, 4},
			target:     6,
			expected:   [][]int{{2, 2, 2}, {2, 4}},
		},
		{
			name:       "Complex case with many solutions",
			candidates: []int{1, 2, 3},
			target:     4,
			expected:   [][]int{{1, 1, 1, 1}, {1, 1, 2}, {1, 3}, {2, 2}},
		},
		{
			name:       "Larger numbers",
			candidates: []int{3, 5, 4},
			target:     8,
			expected:   [][]int{{3, 5}, {4, 4}},
		},
		{
			name:       "Single large target",
			candidates: []int{7, 14, 21},
			target:     21,
			expected:   [][]int{{7, 7, 7}, {7, 14}, {21}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := combinationSum(tt.candidates, tt.target)
			
			if !equalCombinations(actual, tt.expected) {
				t.Errorf("combinationSum(%v, %d) = %v, expected %v", tt.candidates, tt.target, actual, tt.expected)
			}
		})
	}
}

func TestCombinationSumProperties(t *testing.T) {
	t.Run("All combinations sum to target", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 8
		result := combinationSum(candidates, target)
		
		for i, combination := range result {
			sum := 0
			for _, num := range combination {
				sum += num
			}
			if sum != target {
				t.Errorf("Combination %d (%v) sums to %d, expected %d", i, combination, sum, target)
			}
		}
	})

	t.Run("All elements in combinations are from candidates", func(t *testing.T) {
		candidates := []int{2, 3, 6, 7}
		target := 7
		result := combinationSum(candidates, target)
		
		candidateSet := make(map[int]bool)
		for _, candidate := range candidates {
			candidateSet[candidate] = true
		}
		
		for i, combination := range result {
			for j, num := range combination {
				if !candidateSet[num] {
					t.Errorf("Combination %d, element %d (%d) is not in candidates %v", i, j, num, candidates)
				}
			}
		}
	})

	t.Run("All combinations are unique", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 8
		result := combinationSum(candidates, target)
		
		seen := make(map[string]bool)
		for i, combination := range result {
			// Sort combination to create consistent key
			sorted := make([]int, len(combination))
			copy(sorted, combination)
			sort.Ints(sorted)
			
			key := ""
			for _, num := range sorted {
				key += string(rune(num + 1000)) // Convert to string for map key
			}
			
			if seen[key] {
				t.Errorf("Duplicate combination found at index %d: %v", i, combination)
			}
			seen[key] = true
		}
	})

	t.Run("Empty candidates returns empty result for positive target", func(t *testing.T) {
		candidates := []int{}
		target := 5
		result := combinationSum(candidates, target)
		
		if len(result) != 0 {
			t.Errorf("combinationSum([], %d) = %v, expected []", target, result)
		}
	})

	t.Run("Combinations are in non-descending order", func(t *testing.T) {
		candidates := []int{2, 3, 6, 7}
		target := 7
		result := combinationSum(candidates, target)
		
		for i, combination := range result {
			for j := 1; j < len(combination); j++ {
				if combination[j] < combination[j-1] {
					t.Errorf("Combination %d (%v) is not in non-descending order", i, combination)
					break
				}
			}
		}
	})
}

func TestCombinationSumEdgeCases(t *testing.T) {
	t.Run("Large target with small candidates", func(t *testing.T) {
		candidates := []int{1}
		target := 10
		result := combinationSum(candidates, target)
		expected := [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
		
		if !equalCombinations(result, expected) {
			t.Errorf("Large target test failed: got %v, expected %v", result, expected)
		}
	})

	t.Run("All candidates larger than target", func(t *testing.T) {
		candidates := []int{5, 10, 15}
		target := 3
		result := combinationSum(candidates, target)
		
		if len(result) != 0 {
			t.Errorf("Expected empty result when all candidates > target, got %v", result)
		}
	})

	t.Run("Negative target", func(t *testing.T) {
		candidates := []int{1, 2, 3}
		target := -1
		result := combinationSum(candidates, target)
		
		if len(result) != 0 {
			t.Errorf("Expected empty result for negative target, got %v", result)
		}
	})
}

func BenchmarkCombinationSum(b *testing.B) {
	benchmarks := []struct {
		name       string
		candidates []int
		target     int
	}{
		{"Small", []int{2, 3, 5}, 8},
		{"Medium", []int{2, 3, 6, 7}, 7},
		{"Large", []int{1, 2, 3, 4, 5}, 10},
		{"Single", []int{1}, 15},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				combinationSum(bm.candidates, bm.target)
			}
		})
	}
}
