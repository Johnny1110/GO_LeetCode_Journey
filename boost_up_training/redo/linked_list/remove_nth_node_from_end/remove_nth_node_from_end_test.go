package remove_nth_node_from_end

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
	// Input: head = [1,2,3,4,5], n = 2
	// Output: [1,2,3,5]
	head := buildList([]int{1, 2, 3, 4, 5})
	result := removeNthFromEnd(head, 2)
	assert.Equal(t, []int{1, 2, 3, 5}, listToSlice(result))
}

func Test_2(t *testing.T) {
	// Input: head = [1], n = 1
	// Output: []
	head := buildList([]int{1})
	result := removeNthFromEnd(head, 1)
	assert.Nil(t, result)
}

func Test_3(t *testing.T) {
	// Input: head = [1,2], n = 1
	// Output: [1]
	head := buildList([]int{1, 2})
	result := removeNthFromEnd(head, 1)
	assert.Equal(t, []int{1}, listToSlice(result))
}

func Test_4(t *testing.T) {
	// Input: head = [1,2,3,4,5], n = 5 (remove first node)
	// Output: [2,3,4,5]
	head := buildList([]int{1, 2, 3, 4, 5})
	result := removeNthFromEnd(head, 5)
	assert.Equal(t, []int{2, 3, 4, 5}, listToSlice(result))
}
