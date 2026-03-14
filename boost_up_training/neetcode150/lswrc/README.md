# 3. Longest Substring Without Repeating Characters

```go
func lengthOfLongestSubstring(s string) int {
	result, tmp := 0, 0

	charMap := make(map[uint8]int)
	left, right := 0, 0

	for right < len(s) {
		char := s[right]

		if idx, exist := charMap[char]; exist {
			for left <= idx {
				delete(charMap, s[left])
				left++
			}
		}

		charMap[char] = right
		right++

		tmp = right - left
		result = max(result, tmp)
	}

	return result
}
```


```go
func lengthOfLongestSubstring(s string) int {
	result := 0
	charMap := make(map[uint8]int)

	left := 0
	for right := 0; right < len(s); right++ {
		char := s[right]

		if idx, exists := charMap[char]; exists && idx >= left {
			left = idx + 1
		}

		charMap[char] = right
		result = max(result, right-left+1)

	}

	return result
}
```