package word_ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	visitedWords := make(map[string]bool)
	for _, word := range wordList {
		visitedWords[word] = false
	}

	if _, exists := visitedWords[endWord]; !exists {
		return 0
	}

	Q := WordQueue([]string{beginWord})

	currentLevel := 0
	for Q.Len() != 0 {

		levelLen := Q.Len()
		currentLevel++

		for range levelLen {
			word := Q.Pop()

			for idx := 0; idx < len(word); idx++ {
				var tmpbytes []uint8

				for _, letter := range letters {
					tmpbytes = []uint8(word)
					if letter == tmpbytes[idx] {
						continue // skip duplicate
					}
					tmpbytes[idx] = letter
					tmpStr := string(tmpbytes)

					if visited, exists := visitedWords[tmpStr]; exists && !visited {
						if tmpStr == endWord {
							return currentLevel + 1
						}

						visitedWords[tmpStr] = true
						Q.Push(tmpStr)
					}
				}
			}
		}
	}

	return 0
}

// ==============================================================================

var letters = []uint8{
	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
}

type WordQueue []string

func (q WordQueue) Len() int {
	return len(q)
}

func (q *WordQueue) Push(word string) {
	*q = append(*q, word)
}

func (q *WordQueue) Pop() string {
	ret := (*q)[0]
	(*q) = (*q)[1:]
	return ret
}
