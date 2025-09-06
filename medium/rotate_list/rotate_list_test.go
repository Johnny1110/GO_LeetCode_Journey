package rotate_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	// prepare 1~5
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}

	result := rotateRight(node1, 2)
	assertEQ(t, []int{4, 5, 1, 2, 3}, result)
}

func Test_2(t *testing.T) {
	// prepare 1~5
	node3 := &ListNode{2, nil}
	node2 := &ListNode{1, node3}
	node1 := &ListNode{0, node2}

	result := rotateRight(node1, 4)
	assertEQ(t, []int{2, 0, 1}, result)
}

func Test_3(t *testing.T) {
	// prepare 1~5
	node3 := &ListNode{2, nil}
	node2 := &ListNode{1, node3}
	node1 := &ListNode{0, node2}

	result := rotateRight(node1, 0)
	assertEQ(t, []int{0, 1, 2}, result)
}

func Test_4(t *testing.T) {

	node1 := &ListNode{10, nil}

	result := rotateRight(node1, 0)
	assertEQ(t, []int{10}, result)
}

func Test_unit_test_tool(t *testing.T) {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}

	assertEQ(t, []int{1, 2, 3, 4, 5}, node1)
}

func assertEQ(t *testing.T, expect []int, result *ListNode) {
	resultSlice := []int{}

	head := result
	for head != nil {
		resultSlice = append(resultSlice, head.Val)
		head = head.Next
	}

	assert.Equal(t, expect, resultSlice)
}
func Test_locateNode(t *testing.T) {
	assert.Equal(t, 4, locateTailNodeIdx(5, 0))
	assert.Equal(t, 3, locateTailNodeIdx(5, 1))
	assert.Equal(t, 2, locateTailNodeIdx(5, 2))
	assert.Equal(t, 1, locateTailNodeIdx(5, 3))
	assert.Equal(t, 0, locateTailNodeIdx(5, 4))
	assert.Equal(t, 4, locateTailNodeIdx(5, 5))
	assert.Equal(t, 3, locateTailNodeIdx(5, 6))
	assert.Equal(t, 2, locateTailNodeIdx(5, 7))
	assert.Equal(t, 1, locateTailNodeIdx(5, 8))
	assert.Equal(t, 0, locateTailNodeIdx(5, 9))
}
