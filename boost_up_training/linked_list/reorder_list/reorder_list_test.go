package reorder_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	// Input: head = [1,2,3,4]
	// Output: [1,4,2,3]

	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	reorderList(node1)

	resultHead := node1
	resultValues := []int{}

	for resultHead != nil {
		resultValues = append(resultValues, resultHead.Val)
		resultHead = resultHead.Next
	}

	assert.Equal(t, []int{1, 4, 2, 3}, resultValues)
}

func Test_2(t *testing.T) {
	// Input: head = [1,2,3,4,5]
	// Output: [1,5,2,4,3]

	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	reorderList(node1)

	resultHead := node1
	resultValues := []int{}

	for resultHead != nil {
		resultValues = append(resultValues, resultHead.Val)
		resultHead = resultHead.Next
	}

	assert.Equal(t, []int{1, 5, 2, 4, 3}, resultValues)
}

func Test_3(t *testing.T) {
	// Input: head = [1,2,3,4,5]
	// Output: [1,5,2,4,3]

	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}
	node6 := &ListNode{6, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	reorderList(node1)

	resultHead := node1
	resultValues := []int{}

	for resultHead != nil {
		resultValues = append(resultValues, resultHead.Val)
		resultHead = resultHead.Next
	}

	assert.Equal(t, []int{1, 6, 2, 5, 3, 4}, resultValues)
}

func Test_Stack(t *testing.T) {
	stack := NewListNodeStack()

	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}

	stack.Push(node1)
	stack.Push(node2)
	stack.Push(node3)
	stack.Push(node4)
	stack.Push(node5)

	for !stack.isEmpty() {
		val, _ := stack.Pop()
		fmt.Println("val:", val)
	}
}
