package word_ladder

var allLetters []uint8 = []uint8{
	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
}

type WordQueue []string

func (q WordQueue) Len() int {
	return len(q)
}

func (q *WordQueue) Push(val string) {
	*q = append(*q, val)
}

func (q *WordQueue) Pop() string {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

// ==============================================================================

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}

	wordDict := make(map[string]bool)
	for _, word := range wordList {
		wordDict[word] = false
	}

	if _, exists := wordDict[endWord]; !exists {
		// no change at all.
		return 0
	}

	Q := WordQueue([]string{beginWord})
	currentLevel := 0

	for Q.Len() != 0 {
		currentLevel++ // move to next level
		wordCount := Q.Len()

		//fmt.Printf("currentLevel: %v, wordCount: %v \n", currentLevel, wordCount)

		for range wordCount {
			word := Q.Pop()

			// try change 1 letter
			for i := 0; i < len(word); i++ {

				for _, letter := range allLetters {

					if word[i] == letter {
						continue
					}

					tmp := []byte(word)
					tmp[i] = letter
					tmpWord := string(tmp)

					if visited, exists := wordDict[tmpWord]; !visited && exists {

						if tmpWord == endWord {
							return currentLevel + 1 // answer found
						} else {
							// put to next level
							Q.Push(tmpWord)
							wordDict[tmpWord] = true
						}
					}
				}

			}

		}
	}

	return 0
}
