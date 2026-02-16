package find_duclicate

import "testing"

func TestSandBox(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 4},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "Example 3",
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("findDuplicate(%v) = %d; expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			name:     "Duplicate at beginning",
			nums:     []int{1, 1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Duplicate at end",
			nums:     []int{1, 2, 3, 4, 5, 5},
			expected: 5,
		},
		{
			name:     "Multiple duplicates of same number",
			nums:     []int{2, 2, 2, 2, 3},
			expected: 2,
		},
		{
			name:     "Large array with duplicate in middle",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 5, 10},
			expected: 5,
		},
		{
			name:     "Minimum size array",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "Sequential with one duplicate",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 7},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("findDuplicate(%v) = %d; expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkFindDuplicate(t *testing.B) {
	nums := []int{1, 3, 4, 2, 2}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		findDuplicate(nums)
	}
}
