# 269. Alien Dictionary

<br>

---

<br>

link: https://leetcode.com/problems/alien-dictionary

<br>

A new alien language uses the English alphabet, but the order of letters is unknown. You are given a list of `words[]` from the alien language’s dictionary, where the words are claimed to be sorted lexicographically according to the language’s rules.

Your task is to determine the correct order of letters in this alien language based on the given words. If the order is valid, return a string containing the unique letters in lexicographically increasing order as per the new language's rules. If there are multiple valid orders, return any one of them.

However, if the given arrangement of words is inconsistent with any possible letter ordering, return an empty string ("").

A string `a` is lexicographically smaller than a string b if, at the first position where they differ, the character in a appears earlier in the alien language than the corresponding character in b. If all characters in the shorter word match the beginning of the longer word, the shorter word is considered smaller.

Note: Your implementation will be tested using a driver code. It will print true if your returned order correctly follows the alien language’s lexicographic rules; otherwise, it will print false.

<br>

Example

```
Input: ["hrn","hrf","er","enn","rfnn"]

Output: "hernf"
```

Explanation:

from "hrn" and "hrf", we know 'n' < 'f'
from "hrf" and "er", we know 'h' < 'e'
from "er" and "enn", we know get 'r' < 'n'
from "enn" and "rfnn" we know 'e'<'r'
so one possibile solution is "hernf"

<br>

## Coding - 1

```go
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
```