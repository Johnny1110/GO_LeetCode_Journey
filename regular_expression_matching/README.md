# 10. Regular Expression Matching

<br>

## desc:

```bash
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

 

Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
 

Constraints:

1 <= s.length <= 20
1 <= p.length <= 20
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
```

<br>

## Topic

* String
* Dynamic Programming
* Recursion

<br>

## Approach 1. Recursion

## Approach 2. Dynamic Programming

<br>

my idea -1:

這個算法是沒有搞懂題目下的產物，*的用意代表前一個字元可以重複出現或者一次也不出現，我以為 *可以拿來匹配任意字元。
搞錯了...

"a*" _> 代表 可以: "" || "aaaaaaa" || "a" || "aa" 

```go
package main 

func myIdea(s, p string) bool {
	if p == "*" {
		return true
	}
	// source string idx + pattern index
	sourceInx, patternInx := 0, 0

	for {
		if sourceInx == len(s) && patternInx == len(p) {
			// when both reach to the end, return true
			return true
		}

		if sourceInx == len(s) || patternInx == len(p) {
			// one of then reach to then end first, return false
			return false
		}

		// if s = p or p = '.' skip this
		if s[sourceInx] == p[patternInx] || p[patternInx] == '.' {
			sourceInx++
			patternInx++
			continue
		} else if p[patternInx] == '*' {
			// 看下一個 PATTERN 自元是否符合 source

			if sourceInx+1 == len(s) { // end of source
				if patternInx+1 == len(p) { // if pattern also reached to the end
					return true
				} else {
					return false
				}
			}

			if s[sourceInx+1] == p[patternInx+1] || p[patternInx] == '.' {
				// both next char are match then both ++
				sourceInx++
				patternInx++
			} else {
				// if not, source ++ pattern stay cool
				sourceInx++
			}
			continue
		} else {
			return false
		}

	}
}
```

<br>

## about Dynamic Programing

<br>

what's the dynamic programming's trick? I feel like it always create a 2D array for contains all the result.
first of all, init base result, and then new result always based on existing calculate result.