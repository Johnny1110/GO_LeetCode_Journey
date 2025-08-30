package length_of_last_word

func lengthOfLastWord(s string) int {
	firstIdx := -1
	secondIdx := -1

	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		if firstIdx == -1 && char != 32 { // 32 is ' '
			firstIdx = i
			continue
		}

		if firstIdx != -1 && char == 32 {
			secondIdx = i
			break
		}
	}

	return firstIdx - secondIdx
}
