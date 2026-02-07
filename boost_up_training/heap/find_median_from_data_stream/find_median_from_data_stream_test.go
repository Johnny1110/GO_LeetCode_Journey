package find_median_from_data_stream

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMedianFinderBasic(t *testing.T) {
	mf := Constructor()

	// Test single element
	mf.AddNum(1)
	if median := mf.FindMedian(); median != 1.0 {
		t.Errorf("Expected median 1.0, got %f", median)
	}

	// Test two elements
	mf.AddNum(2)
	if median := mf.FindMedian(); median != 1.5 {
		t.Errorf("Expected median 1.5, got %f", median)
	}

	// Test three elements (odd count)
	mf.AddNum(3)
	if median := mf.FindMedian(); median != 2.0 {
		t.Errorf("Expected median 2.0, got %f", median)
	}
}

func TestMedianFinderLeetCodeExample(t *testing.T) {
	mf := Constructor()

	// LeetCode example sequence
	operations := []struct {
		operation string
		num       int
		expected  float64
	}{
		{"addNum", 1, 0},       // median not checked here
		{"addNum", 2, 0},       // median not checked here
		{"findMedian", 0, 1.5}, // (1 + 2) / 2 = 1.5
		{"addNum", 3, 0},       // median not checked here
		{"findMedian", 0, 2.0}, // median of [1, 2, 3] is 2
	}

	for i, op := range operations {
		if op.operation == "addNum" {
			mf.AddNum(op.num)
		} else if op.operation == "findMedian" {
			median := mf.FindMedian()
			if math.Abs(median-op.expected) > 1e-9 {
				t.Errorf("Step %d: Expected median %f, got %f", i, op.expected, median)
			}
		}
	}
}

func TestMedianFinderWithDuplicates(t *testing.T) {
	mf := Constructor()

	// Test with duplicate values
	mf.AddNum(5)
	mf.AddNum(5)
	if median := mf.FindMedian(); median != 5.0 {
		t.Errorf("Expected median 5.0, got %f", median)
	}

	mf.AddNum(5)
	if median := mf.FindMedian(); median != 5.0 {
		t.Errorf("Expected median 5.0, got %f", median)
	}

	mf.AddNum(1)
	if median := mf.FindMedian(); median != 5.0 {
		t.Errorf("Expected median 5.0, got %f", median)
	}
}

func TestMedianFinderWithNegativeNumbers(t *testing.T) {
	mf := Constructor()

	// Test with negative numbers
	mf.AddNum(-1)
	if median := mf.FindMedian(); median != -1.0 {
		t.Errorf("Expected median -1.0, got %f", median)
	}

	mf.AddNum(-2)
	if median := mf.FindMedian(); median != -1.5 {
		t.Errorf("Expected median -1.5, got %f", median)
	}

	mf.AddNum(-3)
	if median := mf.FindMedian(); median != -2.0 {
		t.Errorf("Expected median -2.0, got %f", median)
	}
}

func TestMedianFinderMixedPositiveNegative(t *testing.T) {
	mf := Constructor()

	// Test with mixed positive and negative numbers
	numbers := []int{-5, 0, 5, -10, 10}
	expected := []float64{-5.0, -2.5, 0.0, -2.5, 0.0}

	for i, num := range numbers {
		mf.AddNum(num)
		median := mf.FindMedian()
		if math.Abs(median-expected[i]) > 1e-9 {
			t.Errorf("Step %d: Expected median %f, got %f", i, expected[i], median)
		}
	}
}

func TestMedianFinderLargeRange(t *testing.T) {
	mf := Constructor()

	// Test with large numbers
	mf.AddNum(100000)
	mf.AddNum(-100000)
	if median := mf.FindMedian(); median != 0.0 {
		t.Errorf("Expected median 0.0, got %f", median)
	}

	mf.AddNum(50000)
	if median := mf.FindMedian(); median != 50000.0 {
		t.Errorf("Expected median 50000.0, got %f", median)
	}
}

func TestMedianFinderSequentialNumbers(t *testing.T) {
	mf := Constructor()

	// Test with sequential numbers 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		mf.AddNum(i)
		var expected float64
		if i%2 == 1 {
			// Odd count: median is middle element
			expected = float64((i + 1) / 2)
		} else {
			// Even count: median is average of two middle elements
			expected = float64(i+1) / 2.0
		}

		median := mf.FindMedian()
		if math.Abs(median-expected) > 1e-9 {
			t.Errorf("After adding %d elements: Expected median %f, got %f", i, expected, median)
		}
	}
}

func TestMedianFinderReverseOrder(t *testing.T) {
	mf := Constructor()

	// Test with numbers added in reverse order: 5, 4, 3, 2, 1
	numbers := []int{5, 4, 3, 2, 1}
	expected := []float64{5.0, 4.5, 4.0, 3.5, 3.0}

	for i, num := range numbers {
		mf.AddNum(num)
		median := mf.FindMedian()
		if math.Abs(median-expected[i]) > 1e-9 {
			t.Errorf("Step %d: Expected median %f, got %f", i, expected[i], median)
		}
	}
}

func TestMedianFinderRandomOrder(t *testing.T) {
	mf := Constructor()

	// Test with numbers in random order
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	for _, num := range numbers {
		mf.AddNum(num)
	}

	// After adding all numbers [1, 1, 2, 3, 4, 5, 6, 9], median should be (3+4)/2 = 3.5
	median := mf.FindMedian()
	expected := 3.5
	if math.Abs(median-expected) > 1e-9 {
		t.Errorf("Expected median %f, got %f", expected, median)
	}
}

func TestMedianFinderStressTest(t *testing.T) {
	mf := Constructor()

	// Add many numbers and verify median is reasonable
	for i := 1; i <= 100; i++ {
		mf.AddNum(i)
	}

	// Median of 1..100 should be 50.5
	median := mf.FindMedian()
	expected := 50.5
	if math.Abs(median-expected) > 1e-9 {
		t.Errorf("Expected median %f, got %f", expected, median)
	}
}

// Benchmark tests
func BenchmarkMedianFinderAddNum(b *testing.B) {
	mf := Constructor()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mf.AddNum(i)
	}
}

func BenchmarkMedianFinderFindMedian(b *testing.B) {
	mf := Constructor()

	// Pre-populate with some numbers
	for i := 0; i < 1000; i++ {
		mf.AddNum(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mf.FindMedian()
	}
}

func BenchmarkMedianFinderMixed(b *testing.B) {
	mf := Constructor()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mf.AddNum(i % 1000)
		if i%10 == 0 {
			mf.FindMedian()
		}
	}
}

func Test_findPositionIdx(t *testing.T) {
	mf := Constructor()
	mf.data = []int{1, 2, 3, 5, 6, 7}
	left, right := 0, len(mf.data)-1
	ans := mf.binarySearch(4, left, right)
	assert.Equal(t, 3, ans)

	ans = mf.binarySearch(6, left, right)
	assert.Equal(t, 5, ans)

	ans = mf.binarySearch(8, left, right)
	assert.Equal(t, 6, ans)

	ans = mf.binarySearch(-1, left, right)
	assert.Equal(t, 0, ans)

	ans = mf.binarySearch(2, left, right)
	assert.Equal(t, 2, ans)
}

func Test_findPositionIdx_2(t *testing.T) {
	mf := Constructor()
	mf.data = []int{1}
	left, right := 0, len(mf.data)-1
	ans := mf.binarySearch(3, left, right)
	assert.Equal(t, 1, ans)
}

func Test_Add(t *testing.T) {
	mf := Constructor()

	mf.AddNum(1)
	mf.AddNum(3)
	mf.AddNum(5)

	fmt.Printf("data: %v\n", mf.data)

	mf.AddNum(2)
	fmt.Printf("data: %v\n", mf.data)
	mf.AddNum(-1)
	fmt.Printf("data: %v\n", mf.data)

}
