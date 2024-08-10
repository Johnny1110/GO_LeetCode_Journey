package valid_parentheses

import (
	"go_leetcode_journey/common"
)

var bracketMap = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

// Define the stack type
type stack []rune

// Push adds an element to the top of the stack
func (s *stack) Push(v rune) {
	*s = append(*s, v)
}

// Pop removes and returns the top element of the stack
func (s *stack) Pop() (rune, bool) {
	l := len(*s)
	if l == 0 {
		return 0, false // Return zero value and false if the stack is empty
	}

	// Get the last element
	top := (*s)[l-1]

	// Update the stack to remove the last element
	*s = (*s)[:l-1]

	return top, true
}

func isValid(input string) bool {
	length := len(input)
	if length%2 == 1 || length == 0 {
		// if len is even, then there must have 1 bracket without another one.
		return false
	}
	runeStack := stack{}

	for _, c := range input {
		if c == '(' || c == '[' || c == '{' {
			runeStack.Push(c)
		} else { // c == ) ] }
			open, b := runeStack.Pop()
			theOpenPair := bracketMap[c]
			if b {
				if open != theOpenPair {
					return false
				}
			} else {
				return false
			}
		}
	}

	_, check := runeStack.Pop()
	if check {
		return false
	}

	return true
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(isValid("{}()[}"), false)
	common.Assert_answer(isValid("{}()}"), false)
	common.Assert_answer(isValid("{}()[]"), true)
	common.Assert_answer(isValid(""), false)
	common.Assert_answer(isValid("{[]}"), true)
	common.Assert_answer(isValid("{{"), false)
}
