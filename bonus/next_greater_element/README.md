# Next Greater Element

<br>

---

<br>

# Desc 

Given a input int array, return an array that contains the first greater element on the right side for each element in the input array. 

<br>

For example:

input: `[2,1,2,4,3]`

answer: `4,2,4,-1,-1]`

<br>

please using Monotonic Stack/Queue to solve this problem.

<br>

## Topic

Monotonic Stack/Queue

<br>

## Thinking

### what is Monotonic Stack/Queue ?

The Monotonic Stack/Queue is a type of stack or queue with an additional restriction: it maintains elements in either increasing or decreasing order.

Like a Increasing Monotonic Stack only allows pushing elements that is bigger then the last pushed element, and vice versa for a decreasing Monotonic Stack.

And it allows to pop elements to maintain the order. (e.g., popping smaller element to keep that monotonic restriction)

<br>

## What kind of Monotonic Stack/Queue is that I needed ?

I need a first in last out data structure. so it must be a stack.
let's make a golang stack first.

```golang

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
```

<br>

### How to use that stack ?

* Iterate through the input array in reverse order.

* If there are no elements in the monotonic stack (which is mean no next greater element exists), insert `-1` into answer array at the same index position.

* If there are more the 1 element in the monotonic stack, compare iterate number with stack.Peek() in a while loop.
  * If peek number is bigger than iterate number, insert peek number into answer array at the same index position, and put iterate number into mono-stack, then break the while loop.
  * If peek number is lower than iterate number, pop the mono-stack, and loop it again.

continuous unit iterate is done.