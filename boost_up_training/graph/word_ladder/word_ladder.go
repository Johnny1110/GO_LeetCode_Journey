package word_ladder

var alphabet = []uint8{
	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
}

type WordQueue []string

func (q WordQueue) Len() int {
	return len(q)
}

func (q *WordQueue) Pop() (string, bool) {
	if len(*q) == 0 {
		return "", false
	}
	res := (*q)[0]
	*q = (*q)[1:]
	return res, true
}

func (q *WordQueue) Push(word string) {
	*q = append(*q, word)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]int)
	for _, word := range wordList {
		wordSet[word] = 1 // 1: valid, 2 : valid & visited
	}

	if wordSet[endWord] != 1 {
		return 0 // endWord not exist at all.
	}

	wordLen := len(beginWord)

	level := 0 // BFS level.
	wordQueue := make(WordQueue, 0)

	wordQueue.Push(beginWord)
	wordSet[beginWord] = 2

	for wordQueue.Len() > 0 {
		currentLevelLen := wordQueue.Len()
		level++ // start doing current BFS level, level + 1.

		// here we process layer one by one.
		for range currentLevelLen {

			word, _ := wordQueue.Pop()
			bytes := []byte(word)

			for idx := range wordLen {

				// remember idx origin alp
				originAlp := bytes[idx]

				// ex: hit -> cog try to change idx0,idx1, idx2
				for _, alp := range alphabet {
					bytes[idx] = alp
					newWord := string(bytes)

					status, _ := wordSet[newWord]
					if status == 1 { // is a valid not been visited

						if newWord == endWord {
							return level + 1 // why + 1, cuz new word is belong to next level.
						}

						wordSet[newWord] = 2 // visited
						wordQueue.Push(newWord)
					}
				}

				// recover origin alp to word
				bytes[idx] = originAlp
			}
		}
	}

	return 0 // not found any path to transform.
}
