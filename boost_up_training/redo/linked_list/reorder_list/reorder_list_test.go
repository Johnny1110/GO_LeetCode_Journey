package reorder_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func listToSlice(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func Test_1(t *testing.T) {
	// Input: head = [1,2,3,4]
	// Output: [1,4,2,3]
	head := buildList([]int{1, 2, 3, 4})
	reorderList(head)
	assert.Equal(t, []int{1, 4, 2, 3}, listToSlice(head))
}

func Test_2(t *testing.T) {
	// Input: head = [1,2,3,4,5]
	// Output: [1,5,2,4,3]
	head := buildList([]int{1, 2, 3, 4, 5})
	reorderList(head)
	assert.Equal(t, []int{1, 5, 2, 4, 3}, listToSlice(head))
}

func Test_3(t *testing.T) {
	// Input: head = [1,2,3,4,5,6]
	// Output: [1,6,2,5,3,4]
	head := buildList([]int{1, 2, 3, 4, 5, 6})
	reorderList(head)
	assert.Equal(t, []int{1, 6, 2, 5, 3, 4}, listToSlice(head))
}

func Test_4(t *testing.T) {
	// Input: head = [1] (single node, no change)
	// Output: [1]
	head := buildList([]int{1})
	reorderList(head)
	assert.Equal(t, []int{1}, listToSlice(head))
}
