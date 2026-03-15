package lswrc

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
