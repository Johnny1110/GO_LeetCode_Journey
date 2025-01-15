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