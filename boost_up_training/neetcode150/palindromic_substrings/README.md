# 647. Palindromic Substrings

<br>

---

<br>

## EX

```
Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
```

```
Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
```

<br>
<br>

## Coding - Expanding from Center

```go
func countSubstrings(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {

		pointerA, pointerB := i, i
		for pointerA >= 0 && pointerB < len(s) {
			if s[pointerA] == s[pointerB] {
				res++ // found a palidromic
			} else {
				break
			}
			pointerA--
			pointerB++
		}

		pointerA, pointerB = i, i+1
		for pointerA >= 0 && pointerB < len(s) {
			if s[pointerA] == s[pointerB] {
				res++ // found a palidromic
			} else {
				break
			}
			pointerA--
			pointerB++
		}
	}

	return res
}
```

### Refine Code Structure

```go
func countSubstrings(s string) int {
	res := 0

	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}

	return res
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: N = len(s)

Time: O(N^2)
Space: O(1)
```

<br>
<br>

## Coding - DP

```go
func countSubstrings(s string) int {
	res := 0

	// init DP
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	n := len(s)
	// iterate -> "aaa"
	for length := 1; length <= n; length++ { // 只看當前長度為 length 的 sub-string
		for i := 0; i+length-1 < n; i++ { // i 從 0 開始，i+j 必須等於 length-1, i+length-1 是 j 當的位置，j 沒頂到頭就可以繼續
			j := i + length - 1

			if s[i] == s[j] && (length <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
		}

	}

	return res
}
```


<br>
<br>

## Time & Space Complexity

```
Assume: N = len(s)

Time: O(N^2)
Space: O(N^2)
```