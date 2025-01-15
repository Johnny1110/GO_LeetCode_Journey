package trapping_rain_water

import "fmt"

type Stack struct {
	container []int
}

func (stack *Stack) Push(val int) {
	stack.container = append(stack.container, val)
}

func (stack *Stack) Pop() int {
	if len(stack.container) == 0 {
		return -99
	}
	val := stack.container[len(stack.container)-1]
	stack.container = stack.container[:len(stack.container)-1]
	return val
}

func (stack *Stack) Len() int {
	return len(stack.container)
}

func (stack *Stack) Peek() int {
	return stack.container[len(stack.container)-1]
}

// ----------------------------------------------------------------------------------

func MonotonicStackExample() {
	question := []int{2, 1, 2, 4, 3}
	answer := nextGreaterElement(question)
	fmt.Println("answer: ", answer)
}

// ----------------------------------------------------------------------------------

func nextGreaterElement(input []int) []int {
	monoStack := Stack{}
	greaterArray := make([]int, len(input))

	inputLen := len(input)

	for i := inputLen - 1; i >= 0; i-- {

		for monoStack.Len() > 0 && input[i] > monoStack.Peek() {
			// remove monoStack element if input number is bigger.
			monoStack.Pop()
		}

		if monoStack.Len() > 0 {
			greaterArray[i] = monoStack.Peek()
		} else {
			greaterArray[i] = -1
		}

		monoStack.Push(input[i])
	}

	return greaterArray
}
