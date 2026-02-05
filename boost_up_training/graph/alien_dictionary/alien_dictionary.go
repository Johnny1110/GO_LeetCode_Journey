package alien_dictionary

import "fmt"

// =====================================

type Queue struct {
	letters []uint8
}

func NewQueue() *Queue {
	return &Queue{
		letters: make([]uint8, 0),
	}
}

func (s *Queue) String() string {
	return fmt.Sprintf("%q", s.letters)
}

func (s *Queue) Push(letter uint8) {
	s.letters = append(s.letters, letter)
}

func (s *Queue) Pop() (uint8, bool) {
	l := len(s.letters)
	if l == 0 {
		return 0, false
	}

	ret := s.letters[0]
	s.letters = s.letters[1:]
	return ret, true
}

func (s *Queue) IsEmpty() bool {
	return len(s.letters) == 0
}

// =====================================

func foreignDictionary(words []string) string {
	if len(words) == 0 {
		// prevent no input
		return ""
	}

	// Define - an in-degree map, key: char, value: inDegree.
	inDegree := make(map[uint8]int)
	// Define - an out-edges map, key: node, values: pointing to other nodes.
	outEdges := make(map[uint8]map[uint8]bool)
	// Define - Letter Queue
	S := NewQueue()

	// init inDegree map with all letters:
	for i := 0; i < len(words); i++ {
		word := words[i]
		for k := 0; k < len(word); k++ {
			inDegree[word[k]] = 0
		}
	}

	// create outEdges and inDegree value
	for i := 0; i < len(words)-1; i++ {
		w1 := words[i]
		w2 := words[i+1]

		// find a -> b
		a, b, ok, valid := figureOutABRelationship(w1, w2)
		if !valid {
			return ""
		}

		if ok {
			if outEdges[a] == nil {
				// init
				outEdges[a] = make(map[uint8]bool)
			}

			if !outEdges[a][b] { // never been put into a's outEdges
				inDegree[b]++
				outEdges[a][b] = true
			}
		}
	}

	//debugInDegree(inDegree)
	//debugOutEdges(outEdges)

	result := make([]uint8, 0)

	// Init S, put all 0 inDegree letters into S
	for k, v := range inDegree {
		if v == 0 {
			S.Push(k)
		}
	}

	// Start BFS
	for !S.IsEmpty() {

		pLetter, _ := S.Pop()
		result = append(result, pLetter)

		for oLetter, _ := range outEdges[pLetter] {
			inDegree[oLetter]--
			if inDegree[oLetter] == 0 {
				// on one pointing to it, upgrade to S
				S.Push(oLetter)
			}
		}

		if S.IsEmpty() && len(result) != len(inDegree) {
			// have cycle, can resolve
			return ""
		}

	}

	return string(result)
}

// figureOutABRelationship word1, word2, found, valid
func figureOutABRelationship(w1 string, w2 string) (uint8, uint8, bool, bool) {
	i := 0
	for i < len(w1) && i < len(w2) {
		if w1[i] != w2[i] {
			return w1[i], w2[i], true, true
		} else {
			i++
		}
	}

	// No difference found in common prefix
	// If w1 is longer than w2, it's invalid: ["abc", "ab"]
	if len(w1) > len(w2) {
		return 0, 0, false, false // invalid
	}

	return 0, 0, false, true
}
