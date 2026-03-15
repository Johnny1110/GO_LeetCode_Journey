# 5. Longest Palindromic

<br>

---

<br>

## Coding

```go

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	bestLeft, bestRight := 0, 0
	bestLen := 0

	for i := 0; i < len(s); i++ {

		leftP, rightP := i-1, i+1

		// explore:
		for leftP >= 0 && rightP < len(s) && s[leftP] == s[rightP] {
			currentLen := rightP - leftP + 1
			if bestLen < currentLen {
				bestLeft, bestRight, bestLen = leftP, rightP, currentLen
			}
			leftP--
			rightP++
		}

		leftP, rightP = i-1, i

		for leftP >= 0 && rightP < len(s) && s[leftP] == s[rightP] {
			currentLen := rightP - leftP + 1
			if bestLen < currentLen {
				bestLeft, bestRight, bestLen = leftP, rightP, currentLen
			}
			leftP--
			rightP++
		}

	}

	return s[bestLeft : bestRight+1]
}
```