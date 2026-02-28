package longest_increasing_subsequence

import (
	"fmt"
	"math"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Example 1",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4, // [2,3,7,18] or [2,3,7,101]
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4, // [0,1,2,3]
		},
		{
			name:     "Example 3",
			nums:     []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1, // [7] - all elements are same
		},
		{
			name:     "strictly increasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5, // entire array
		},
		{
			name:     "strictly decreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1, // any single element
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "two elements ascending",
			nums:     []int{1, 3},
			expected: 2, // [1, 3]
		},
		{
			name:     "two elements descending",
			nums:     []int{3, 1},
			expected: 1, // [3] or [1]
		},
		{
			name:     "alternating pattern",
			nums:     []int{1, 3, 2, 4, 3, 5, 4, 6},
			expected: 4, // [1,2,3,4] or [1,2,4,5] or [1,3,4,6]
		},
		{
			name:     "with negative numbers",
			nums:     []int{-10, -3, 0, 5, 9},
			expected: 5, // entire array
		},
		{
			name:     "mixed negative and positive",
			nums:     []int{-1, 3, -3, 4, 0, 6},
			expected: 4, // [-1,0,4,6] or [-3,0,4,6]
		},
		{
			name:     "all negative decreasing",
			nums:     []int{-1, -2, -3, -4},
			expected: 1, // any single element
		},
		{
			name:     "all negative increasing",
			nums:     []int{-4, -3, -2, -1},
			expected: 4, // entire array
		},
		{
			name:     "large numbers",
			nums:     []int{1000000, 2000000, 1500000, 3000000},
			expected: 3, // [1000000, 2000000, 3000000]
		},
		{
			name:     "duplicate numbers scattered",
			nums:     []int{1, 2, 2, 3, 3, 4},
			expected: 4, // [1,2,3,4]
		},
		{
			name:     "plateau pattern",
			nums:     []int{1, 2, 2, 2, 3},
			expected: 3, // [1,2,3]
		},
		{
			name:     "zigzag pattern",
			nums:     []int{1, 5, 2, 6, 3, 7, 4, 8},
			expected: 5, // [1,2,3,4,8] or other combinations
		},
		{
			name:     "reverse mountain",
			nums:     []int{1, 2, 3, 2, 1},
			expected: 3, // [1,2,3]
		},
		{
			name:     "mountain pattern",
			nums:     []int{3, 1, 2, 4, 3},
			expected: 3, // [1,2,4] or [1,2,3]
		},
		{
			name:     "long random sequence",
			nums:     []int{10, 22, 9, 33, 21, 50, 41, 60, 80},
			expected: 5, // [10,22,33,50,60] or [9,21,41,60,80]
		},
		{
			name:     "with zeros",
			nums:     []int{0, 0, 1, 1, 2, 2},
			expected: 3, // [0,1,2]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLIS(tt.nums)
			if result != tt.expected {
				t.Errorf("lengthOfLIS(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestLengthOfLISEdgeCases(t *testing.T) {
	t.Run("minimum value constraint", func(t *testing.T) {
		nums := []int{math.MinInt32}
		expected := 1
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("maximum value constraint", func(t *testing.T) {
		nums := []int{math.MaxInt32}
		expected := 1
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("extreme values mixed", func(t *testing.T) {
		nums := []int{math.MinInt32, 0, math.MaxInt32}
		expected := 3
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("all same elements", func(t *testing.T) {
		nums := make([]int, 100)
		for i := range nums {
			nums[i] = 42
		}
		expected := 1
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(array of 100 same elements) = %d, expected %d", result, expected)
		}
	})
}

func TestLengthOfLISComplexPatterns(t *testing.T) {
	t.Run("fibonacci-like sequence", func(t *testing.T) {
		nums := []int{1, 1, 2, 3, 5, 8, 13}
		expected := 6 // [1,2,3,5,8,13]
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("powers of 2", func(t *testing.T) {
		nums := []int{1, 2, 4, 8, 16, 32}
		expected := 6 // entire sequence
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("reverse powers of 2", func(t *testing.T) {
		nums := []int{32, 16, 8, 4, 2, 1}
		expected := 1 // any single element
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("mixed powers", func(t *testing.T) {
		nums := []int{1, 4, 2, 8, 3, 16}
		expected := 4 // [1,2,3,16] or [1,2,8,16]
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("arithmetic progression", func(t *testing.T) {
		nums := []int{2, 4, 6, 8, 10}
		expected := 5 // entire sequence
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("interleaved sequences", func(t *testing.T) {
		nums := []int{1, 10, 2, 20, 3, 30, 4, 40}
		expected := 4 // [1,2,3,4] or [10,20,30,40]
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})
}

func TestLengthOfLISLargeInputs(t *testing.T) {
	t.Run("large ascending array", func(t *testing.T) {
		nums := make([]int, 1000)
		for i := range nums {
			nums[i] = i + 1
		}
		expected := 1000
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(1000 ascending elements) = %d, expected %d", result, expected)
		}
	})

	t.Run("large descending array", func(t *testing.T) {
		nums := make([]int, 1000)
		for i := range nums {
			nums[i] = 1000 - i
		}
		expected := 1
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(1000 descending elements) = %d, expected %d", result, expected)
		}
	})

	t.Run("large random pattern", func(t *testing.T) {
		// Create a pattern where every 3rd element is in increasing order
		nums := make([]int, 300)
		lisLength := 0
		for i := 0; i < 100; i++ {
			nums[i*3] = i*10 + 100   // these form the LIS
			nums[i*3+1] = i*10 + 50  // smaller than LIS element
			nums[i*3+2] = i*10 + 150 // larger but non-consecutive
			lisLength++
		}
		expected := 100 // every 3rd element
		result := lengthOfLIS(nums)
		if result < expected {
			t.Errorf("lengthOfLIS(300 elements with pattern) = %d, expected at least %d", result, expected)
		}
	})
}

func TestLengthOfLISSpecialSequences(t *testing.T) {
	t.Run("prime numbers", func(t *testing.T) {
		nums := []int{2, 3, 5, 7, 11, 13, 17}
		expected := 7 // all primes in increasing order
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("squares", func(t *testing.T) {
		nums := []int{1, 4, 9, 16, 25}
		expected := 5 // all squares in increasing order
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("mixed squares", func(t *testing.T) {
		nums := []int{1, 9, 4, 16, 25}
		expected := 4 // [1,4,16,25] or [1,9,16,25]
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})

	t.Run("binary representation inspired", func(t *testing.T) {
		nums := []int{1, 10, 11, 100, 101, 110, 111} // binary-like
		expected := 7                                // entire sequence is increasing
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("lengthOfLIS(%v) = %d, expected %d", nums, result, expected)
		}
	})
}

func BenchmarkLengthOfLIS(b *testing.B) {
	nums := []int{10, 22, 9, 33, 21, 50, 41, 60, 80, 1, 5, 15, 25, 35, 45}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLIS(nums)
	}
}

func BenchmarkLengthOfLISWorstCase(b *testing.B) {
	// Worst case: decreasing sequence
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = 100 - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLIS(nums)
	}
}

func BenchmarkLengthOfLISBestCase(b *testing.B) {
	// Best case: increasing sequence
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLIS(nums)
	}
}

func BenchmarkLengthOfLISLarge(b *testing.B) {
	// Large input with mixed pattern
	nums := make([]int, 1000)
	for i := range nums {
		if i%2 == 0 {
			nums[i] = i/2 + 1
		} else {
			nums[i] = 1000 - i/2
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLIS(nums)
	}
}

func TestBinarySearch(t *testing.T) {
	tails := []int{2, 5}
	ans := binarySearch(tails, 3)
	fmt.Printf("---ans: %v\n", ans)
}
