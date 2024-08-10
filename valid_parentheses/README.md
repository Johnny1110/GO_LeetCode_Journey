# Valid Parentheses

<br>

---

<br>

## Desc

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.



Example 1:
```
Input: s = "()"
Output: true
```

Example 2:
```
Input: s = "()[]{}"
Output: true
```
Example 3:
```
Input: s = "(]"
Output: false
```

<br>
<br>

Constraints:

    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.

<br>

## Topic

* String
* Stack


<br>

## Thinking

LeetCode Hint say using stack top resolve it.

Iterate through all the char in input string, if encounter a open bracket like `( { [`, push em
to the top of the stack, if encounter a close bracket like `) ] }` then check the top of the stack if it is
a mapping open bracket.

but I think I don't want to do it like that, what don't iterate though the input string then check i and i+1 ?

let's try it out.

<br>

after like 10 mins later, I did it like this:

```golang
var bracketMap = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	//for _, c := range s {
	//	fmt.Println(string(c))
	//}

	length := len(s)
	if length%2 == 1 || length == 0 {
		// if len is even, then there must have 1 bracket without another one.
		return false
	}
	for i := 0; i < length; i++ {
		c1 := rune(s[i])
		c2 := rune(s[i+1])
		if c2 != bracketMap[c1] {
			return false
		}
		i = i + 1
	}

	return true
}
```

<br>

and there's a error:

```
Input
s = "{[]}"

Output: false
Expected: true
```

<br>

Oh... that's why leetcode give me hint using stack.

<br>

let's do it again.






