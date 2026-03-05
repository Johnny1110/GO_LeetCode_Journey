package alien_dictionary

func foreignDictionary(words []string) string {
	if len(words) == 0 {
		return ""
	}

	dirgraph := make(map[uint8]CharSet)
	inDegree := make(map[uint8]int)

	// init inDegree + dirgraph
	for _, word := range words {
		for _, char := range word {
			inDegree[uint8(char)] = 0
			dirgraph[uint8(char)] = NewCharSet()
		}
	}

	for idx := 1; idx < len(words); idx++ {
		wordA, wordB := words[idx-1], words[idx]
		length := min(len(wordA), len(wordB))
		foundDiff := false

		for x := range length {
			charA, charB := wordA[x], wordB[x]
			if charA != charB {
				foundDiff = true

				if _, exist := dirgraph[charA][charB]; !exist {
					dirgraph[charA].Add(charB)
					inDegree[charB]++
				}

				break
			}
		}

		if !foundDiff && len(wordA) > len(wordB) {
			return ""
		}
	}

	// put all in degree-0 into queue as init
	Q := CharQueue([]uint8{})
	for char, ide := range inDegree {
		if ide == 0 {
			Q.Push(char)
		}
	}

	result := []uint8{}

	for Q.Len() != 0 {

		levelLen := Q.Len()

		for range levelLen {
			char := Q.Pop()

			result = append(result, char)

			// clean in-degree
			for intoChar, _ := range dirgraph[char] {
				inDegree[intoChar]--
				if inDegree[intoChar] == 0 {
					Q.Push(intoChar)
				}
			}
		}

	}

	resStr := string(result)

	if len(resStr) != len(inDegree) {
		return ""
	} else {
		return resStr
	}
}

type CharSet map[uint8]struct{}

func (s CharSet) Add(c uint8) {
	s[c] = struct{}{}
}

func NewCharSet() CharSet {
	return CharSet(make(map[uint8]struct{}))
}

type CharQueue []uint8

func (q CharQueue) Len() int {
	return len(q)
}

func (q *CharQueue) Push(c uint8) {
	*q = append(*q, c)
}

func (q *CharQueue) Pop() uint8 {
	ret := (*q)[0]
	(*q) = (*q)[1:]
	return ret
}
