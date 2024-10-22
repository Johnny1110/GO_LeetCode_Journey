# Longest Valid Parentheses

<br>

---

<br>

## Desc

Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses
substring.

<br>


Example 1:

```
Input: s = "(()"
Output: 2
```

Explanation: The longest valid parentheses substring is "()".

<br>

Example 2:

```
Input: s = ")()())"
Output: 4
```

Explanation: The longest valid parentheses substring is "()()".

<br>

Example 3:

```
Input: s = ""
Output: 0
```

<br>

Constraints:

```
0 <= s.length <= 3 * 104
s[i] is '(', or ')'.
```

<br>
<br>

## Topic

* String
* Dynamic Programming
* Stack

<br>

## Thinking

According to the topic, I start to thinking about how should I use Stack and Dynamic Programming.

for Stack, I guess it's a temporary storage for single parentheses (not matched yet).

first of all, init a int result.

when we iterate through the string:

* if it is a "(", put it into the stack.
* if it is a ")", pop the top of the stack,
    * if the popped element is a "(", it's a matched up. tempResult += 1;
    * if the popped element is a ")", means not match, clean the stack, and `result = max(result, tempResult)`
* return the result.

it's seems like no need to use Dynamic Programming.

<br>

thus, I wrote the following code:

```go
type Uint8Stack struct {
	stack []uint8
}

func (stack *Uint8Stack) push(s uint8) {
	stack.stack = append(stack.stack, s)
}

func (stack *Uint8Stack) pop() uint8 {
	if len(stack.stack) == 0 {
		return 0
	}
	node := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return node
}

func (stack *Uint8Stack) reset() {
	stack.stack = stack.stack[:0]
}

func (stack *Uint8Stack) len() int {
	return len(stack.stack)
}

// ---------------------------------

func longestValidParentheses(s string) int {
	stack := &Uint8Stack{}
	longest := 0
	accumulator := 0

	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '(' {
			stack.push(v)
		} else {
			if stack.pop() == '(' {
				accumulator += 2
			} else {
				longest = max(longest, accumulator)
				accumulator = 0
				stack.reset()
			}
		}
	}

	return max(longest, accumulator)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

<br>

that can not solve the problem like : 

```
"()(()"
"()(((((()"
"()(((((())"
"()(((((()()"
```

<br>

I start to think about using DP (Dynamic Programming).

<br>

init a 2D array. (x, y) -> how many parentheses pairs from idx `x` to idx `y`.

<br>

I asked CHAT-GPT for a hint, and I got this:

```

You're on the right track with your understanding of the problem, but both the stack and dynamic programming approaches you've written have some inefficiencies. Let me guide you with some hints to refine your approach.

Stack Approach:
The main issue with your current stack-based approach is that you are using an accumulator to track the length of valid parentheses and resetting it whenever you encounter a mismatch. This doesn't handle cases where a mismatch doesn't completely invalidate previous valid substrings (like in the example "()(()").

Instead of resetting the stack or accumulator, you could use the indices of the parentheses to compute the length of valid substrings. Here's the revised approach:

Use a stack to store indices:
Push the index of each '('.
On encountering ')', pop the stack.
If the stack is empty after popping, push the current index (this is a boundary for future valid substrings).
If the stack is not empty, compute the length of the valid substring by subtracting the current index from the top of the stack.
This ensures you keep track of the boundaries of valid substrings without resetting.

DP Approach:
The DP approach you wrote is creating a 2D array, which results in a space complexity of O(n^2). This is the reason for the memory limit error on larger inputs.

A better way to use DP would be to maintain a 1D DP array, where:

dp[i] represents the length of the longest valid parentheses ending at index i.
If s[i] == ')', check if s[i-1] == '(' or if the substring before a matching ( forms a valid substring, then update dp[i] accordingly.
Optimized DP Solution:
go
Copy code
func longestValidParentheses(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }
    
    dp := make([]int, n)
    maxLen := 0
    
    for i := 1; i < n; i++ {
        if s[i] == ')' {
            if s[i-1] == '(' {
                // Case: "()", we can extend 2
                dp[i] = 2
                if i-2 >= 0 {
                    dp[i] += dp[i-2]
                }
            } else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
                // Case: ")"
                dp[i] = dp[i-1] + 2
                if i-dp[i-1]-2 >= 0 {
                    dp[i] += dp[i-dp[i-1]-2]
                }
            }
        }
        maxLen = max(maxLen, dp[i])
    }
    
    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
Explanation:
If the current character is ')' and the previous one is '(', a valid pair () is found, and you extend the length by 2.
If the previous character is also ')', you look for a matching '(' before the valid substring formed by dp[i-1] and extend accordingly.
This 1D DP approach has a time complexity of O(n) and space complexity of O(n), making it much more efficient than your current solution.

This should handle all the edge cases you've mentioned and avoid memory issues!
```




