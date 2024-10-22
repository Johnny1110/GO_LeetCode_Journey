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



