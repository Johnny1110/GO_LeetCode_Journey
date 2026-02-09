package merge_k_sorted_lists

import (
	"reflect"
	"testing"
)

// Helper function to create a linked list from slice
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	
	head := &ListNode{Val: vals[0]}
	current := head
	
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	
	return head
}

// Helper function to convert linked list to slice for comparison
func listToSlice(head *ListNode) []int {
	result := []int{}
	current := head
	
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	
	return result
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "Example 1 - Three sorted lists",
			input:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "Empty lists array",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name:     "Single empty list",
			input:    [][]int{{}},
			expected: []int{},
		},
		{
			name:     "Single non-empty list",
			input:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Two lists - one empty",
			input:    [][]int{{1, 3, 5}, {}},
			expected: []int{1, 3, 5},
		},
		{
			name:     "All empty lists",
			input:    [][]int{{}, {}, {}},
			expected: []int{},
		},
		{
			name:     "Lists with duplicate values",
			input:    [][]int{{1, 1, 2}, {1, 2, 3}},
			expected: []int{1, 1, 1, 2, 2, 3},
		},
		{
			name:     "Lists with negative numbers",
			input:    [][]int{{-2, 0, 2}, {-1, 1, 3}},
			expected: []int{-2, -1, 0, 1, 2, 3},
		},
		{
			name:     "Single element lists",
			input:    [][]int{{5}, {2}, {8}, {1}},
			expected: []int{1, 2, 5, 8},
		},
		{
			name:     "Large numbers",
			input:    [][]int{{100, 200, 300}, {150, 250, 350}},
			expected: []int{100, 150, 200, 250, 300, 350},
		},
		{
			name:     "Different length lists",
			input:    [][]int{{1}, {2, 3, 4, 5, 6}, {7, 8}},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert input slices to linked lists
			lists := make([]*ListNode, len(tt.input))
			for i, vals := range tt.input {
				lists[i] = createList(vals)
			}

			// Call the function
			result := mergeKLists(lists)
			
			// Convert result to slice for comparison
			actual := listToSlice(result)

			// Compare with expected
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("mergeKLists() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestMergeKListsEdgeCases(t *testing.T) {
	t.Run("Nil input", func(t *testing.T) {
		result := mergeKLists(nil)
		if result != nil {
			t.Errorf("mergeKLists(nil) = %v, expected nil", result)
		}
	})

	t.Run("Very large input", func(t *testing.T) {
		// Create 1000 single-element lists
		lists := make([]*ListNode, 1000)
		expected := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			lists[i] = createList([]int{i})
			expected[i] = i
		}

		result := mergeKLists(lists)
		actual := listToSlice(result)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Large input test failed")
		}
	})
}
