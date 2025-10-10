package linked_list_cycle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	// Input: head = [3,2,0,-4], pos = 1
	// Output: true
	// Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

	node1 := &ListNode{3, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{0, nil}
	node4 := &ListNode{-4, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	assert.True(t, hasCycle(node1))
}

func Test_2(t *testing.T) {
	node1 := &ListNode{3, nil}
	node2 := &ListNode{2, nil}

	node1.Next = node2
	node2.Next = node1

	assert.True(t, hasCycle(node1))
}

func Test_3(t *testing.T) {
	node1 := &ListNode{3, nil}
	node2 := &ListNode{2, nil}

	node1.Next = node2

	assert.False(t, hasCycle(node1))
}

func Test_4(t *testing.T) {
	node1 := &ListNode{3, nil}

	assert.False(t, hasCycle(node1))
}
