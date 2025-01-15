package next_greater_element

import (
	"go_leetcode_journey/common"
)

type Stack struct {
	container []int
}

func (s *Stack) Push(v int) {
	s.container = append(s.container, v)
}

func (s *Stack) Pop() int {
	if len(s.container) == 0 {
		return -1
	}
	value := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]
	return value
}

func (s *Stack) Peek() int {
	if len(s.container) != 0 {
		return s.container[len(s.container)-1]
	} else {
		return -1
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.container) == 0
}

func (s *Stack) Size() int {
	return len(s.container)
}

// ----------------------------------------------------------------------------------

func nextGreaterElement(input []int) []int {

	monoStack := Stack{}
	answer := make([]int, len(input))

	for i := len(input) - 1; i >= 0; i-- {

		for !monoStack.IsEmpty() && input[i] >= monoStack.Peek() {
			monoStack.Pop()
		}

		if monoStack.IsEmpty() {
			answer[i] = -1
		} else {
			answer[i] = monoStack.Peek()
		}

		monoStack.Push(input[i])
	}

	return answer
}

// Go call this func in main.go
func Go() {
	question := []int{2, 1, 2, 4, 3}
	answer := nextGreaterElement(question)
	common.Assert_answer(answer, []int{4, 2, 4, -1, -1})
}
