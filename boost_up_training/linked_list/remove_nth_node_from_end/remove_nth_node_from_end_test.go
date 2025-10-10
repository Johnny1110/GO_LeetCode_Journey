package remove_nth_node_from_end

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: head = [1,2,3,4,5], n = 2
	//Output: [1,2,3,5]

	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	resultHead := removeNthFromEnd(node1, 2)

	resultValues := []int{}

	for resultHead != nil {
		resultValues = append(resultValues, resultHead.Val)
		resultHead = resultHead.Next
	}

	assert.Equal(t, []int{1, 2, 3, 5}, resultValues)
}

func Test_2(t *testing.T) {
	node1 := &ListNode{1, nil}
	resultHead := removeNthFromEnd(node1, 1)
	assert.Nil(t, resultHead)
}

func Test_3(t *testing.T) {

	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}

	node1.Next = node2

	resultHead := removeNthFromEnd(node1, 1)

	resultValues := []int{}

	for resultHead != nil {
		resultValues = append(resultValues, resultHead.Val)
		resultHead = resultHead.Next
	}

	assert.Equal(t, []int{1}, resultValues)
}

func Test_4(t *testing.T) {
	//Input: head = [1,2,3,4,5], n = 2
	//Output: [1,2,3,5]

	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	resultHead := removeNthFromEnd(node1, 5)

	resultValues := []int{}

	for resultHead != nil {
		resultValues = append(resultValues, resultHead.Val)
		resultHead = resultHead.Next
	}

	assert.Equal(t, []int{2, 3, 4, 5}, resultValues)
}
