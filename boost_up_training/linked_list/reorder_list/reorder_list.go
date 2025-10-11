package reorder_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeStack struct {
	container []*ListNode
}

func NewListNodeStack() *ListNodeStack {
	return &ListNodeStack{
		container: make([]*ListNode, 0),
	}
}

func (s *ListNodeStack) Push(node *ListNode) {
	s.container = append(s.container, node)
}

func (s *ListNodeStack) Pop() (*ListNode, error) {
	length := len(s.container)
	if length == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	node := s.container[length-1]
	s.container = s.container[:length-1]
	return node, nil
}

func (s *ListNodeStack) isEmpty() bool {
	return len(s.container) == 0
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// put all node(except head) into stack
	stack := NewListNodeStack()
	tmpPointer := head.Next
	for tmpPointer != nil {
		stack.Push(tmpPointer)
		tmpPointer = tmpPointer.Next
	}

	pointerA := head
	pointerB, _ := stack.Pop() // tail

	for (pointerA != pointerB) && (pointerA.Next != pointerB) {
		if stack.isEmpty() {
			break
		}

		CPA := pointerA
		CPB := pointerB

		pointerB, _ = stack.Pop()
		pointerA = pointerA.Next

		CPA.Next = CPB
		CPB.Next = pointerA
	}

	pointerB.Next = nil
}
