package subsets

import (
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "example 2",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			name:     "empty input",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "two elements",
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, 0, 1},
			expected: [][]int{{}, {-1}, {0}, {-1, 0}, {1}, {-1, 1}, {0, 1}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsets(tt.nums)
			if len(result) != len(tt.expected) {
				t.Errorf("subsets(%v) returned %d subsets; want %d", tt.nums, len(result), len(tt.expected))
				return
			}
			if !equalSubsets(result, tt.expected) {
				t.Errorf("subsets(%v) = %v; want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

// equalSubsets checks if two slices of slices contain the same subsets (order-independent).
func equalSubsets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	normalize := func(sets [][]int) [][]int {
		copied := make([][]int, len(sets))
		for i, s := range sets {
			c := make([]int, len(s))
			copy(c, s)
			sort.Ints(c)
			copied[i] = c
		}
		sort.Slice(copied, func(i, j int) bool {
			if len(copied[i]) != len(copied[j]) {
				return len(copied[i]) < len(copied[j])
			}
			for k := 0; k < len(copied[i]); k++ {
				if copied[i][k] != copied[j][k] {
					return copied[i][k] < copied[j][k]
				}
			}
			return false
		})
		return copied
	}
	na, nb := normalize(a), normalize(b)
	for i := range na {
		if len(na[i]) != len(nb[i]) {
			return false
		}
		for j := range na[i] {
			if na[i][j] != nb[i][j] {
				return false
			}
		}
	}
	return true
}
