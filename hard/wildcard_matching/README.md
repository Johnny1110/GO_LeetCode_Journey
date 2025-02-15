# Wildcard Matching

<br>

---

<br>

## Desc

Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

<br>

Example 1:

```
Input: s = "aa", p = "a"
Output: false
```

Explanation: "a" does not match the entire string "aa".


Example 2:

```
Input: s = "aa", p = "*"
Output: true
```


Explanation: '*' matches any sequence.


Example 3:

```
Input: s = "cb", p = "?a"
Output: false
```

Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.

<br>

Constraints:

```
0 <= s.length, p.length <= 2000
s contains only lowercase English letters.
p contains only lowercase English letters, '?' or '*'.
```

<br>

## Topic

* String
* Dynamic Programming
* Greedy [Video](https://www.bilibili.com/video/BV18V411b74c/?spm_id_from=333.337.search-card.all.click&vd_source=9780a181ac9f1fee5f680f255ee5bc73), [Article](https://medium.com/@ralph-tech/%E6%BC%94%E7%AE%97%E6%B3%95%E5%AD%B8%E7%BF%92%E7%AD%86%E8%A8%98-%E8%B2%AA%E5%A9%AA%E6%BC%94%E7%AE%97%E6%B3%95-greedy-algorithm-e2666b93d05f)
* Recursion

<br>

## Thinking

| Approach  | Pros  | Cons  |
|---|---|---|
| Recursion  | Simple, intuitive  | 	Slow (exponential) without memoization |
| Dynamic Programming  | Efficient, optimal	  | Uses extra space O(nm)|
| Greedy  | Fastest, O(n + m)	  |  Harder to implement correctly|

I gonna try with DP first, then I will go with Greedy.

<br>

### Dynamic Programming 

* Define `dp[i][j]` as whether `s[0:i]` matches `p[0:j]`.

* Transition:
  * If `p[j-1] == s[i-1]` or `p[j-1] == '?'`, inherit `dp[i-1][j-1]`.
  * If `p[j-1] == '*'`, we have two choices:
    * `dp[i][j-1]`: Treat '*' as an empty sequence.
    * `dp[i-1][j]`: Treat '*' as matching `s[i-1]` (extend the match).

* Base cases:

  * `dp[0][0] = True` (empty string matches empty pattern).
  * `dp[i][0] = False` (non-empty string can't match empty pattern).
  * `dp[0][j]` depends on whether `p[0:j]` consists only of '*'.


<br>

### Greedy (Optimized for '*')

If '*' appears, it can match any number of characters, so:

1. Remember the last position where '*' was used.
2. If we reach a mismatch, backtrack to the last '*' and try matching more characters. 

This approach avoids unnecessary recursion and speeds up execution.

How to Implement:

* Use two pointers: i for s and j for p.
* Track starIndex (last seen '*') and matchIndex (where matching started).
* If `p[j] == '*'`, store starIndex and move j forward.
* If mismatch occurs and we saw a '*', backtrack and retry from `matchIndex + 1`.