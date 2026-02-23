package reverse_linked_list

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
	// Input: head = [1,2,3,4,5]
	// Output: [5,4,3,2,1]
	head := buildList([]int{1, 2, 3, 4, 5})
	result := reverseList(head)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, listToSlice(result))
}

func Test_2(t *testing.T) {
	// Input: head = [1,2]
	// Output: [2,1]
	head := buildList([]int{1, 2})
	result := reverseList(head)
	assert.Equal(t, []int{2, 1}, listToSlice(result))
}

func Test_3(t *testing.T) {
	// Input: head = [] (empty list)
	// Output: nil
	result := reverseList(nil)
	assert.Nil(t, result)
}

func Test_4(t *testing.T) {
	// Input: head = [1] (single node)
	// Output: [1]
	head := buildList([]int{1})
	result := reverseList(head)
	assert.Equal(t, []int{1}, listToSlice(result))
}
