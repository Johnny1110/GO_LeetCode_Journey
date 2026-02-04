# 127. Word Ladder

<br>

---

<br>

link: https://leetcode.com/problems/word-ladder/description/

<br>
<br>

## Thinking

### Topics:

* Hash Table
* String
* Breadth-First Search

<br>

### Constraints

```
1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord, endWord, and wordList[i] consist of lowercase English letters.
beginWord != endWord
All the words in wordList are unique.
```

<br>

### Deconstruct the Problem

First, the `wordList` order does not matter at all.

I think we can try to find the first word that can be transformed to.

like:

```
beginWord = "hit"
endWord = "cog"
wordList = ["cog", "dog", "dot", "hot", "hog", "lot", "log"]
```

#### first round:

* target word can only be "hot" or "hog" (cuz it only changed 1 letter)
* so we have 2 options `["hot", "hog"]`
* now we put em into queue (BFS approach)

#### sec round:

* pop first element "hot" from BFS queue, 
* now we know "o" is done, we need to change "h" or "t" in this round. 
* we can find `["dot", "hog"]` in this round, but "hog" already in used(in first round) so we put "dot" into queue.
* now the queue is like `["hog", "dot"]`.

#### third round:

* pop first element "hog" from BFS queue, 
* now we know "og" is done, we need to only change "h" or "c" in this round.
* we already found the right answer in this round.

<br>

so I guess maybe I can use that approach to solve it. maybe wrap transform count into queue element like this:

```go
type WordElement struct {
	word string,
	previousStep int
}
```

each time put a new word into queue, wrap the previous step count into it also.


<br>
<br>

### Claude AI feedback

#### What You Got Right

* BFS is the correct approach — Because BFS explores level by level, the first time you reach endWord, you're guaranteed to have found the shortest path.
* Tracking visited words — You correctly identified that we shouldn't revisit words (like "hog" in your example). This prevents infinite loops and redundant work.
* Tracking the step count — You need to know "how many transformations" to return the answer.

#### A Small Refinement to Consider

Your WordElement struct works, but think about this property of BFS:

>In BFS, all elements at the same "level" are processed before moving to the next level.

```
Level 0:  hit
Level 1:  hot, hog         ← all words here have step count = 1
Level 2:  dot, dog, log    ← all words here have step count = 2
Level 3:  cog              ← step count = 3
```

Question: Given this property, do you need to store the step count in each queue element? Or is there another way to track which "level" you're currently processing?

Both approaches are valid, but thinking about this will deepen your understanding of BFS.

#### Next Challenge

You've got the high-level algorithm. Now think about implementation details:

* The expensive operation: Given a word like "hot", how do you efficiently find all valid transformations in wordList?

* Naive approach: 
  * Compare "hot" against every word in wordList → O(N × L) per word, where L is word length
  * Can you think of a smarter way?

<br>
<br>

### Thinking Enhancement

In each round, if I popped all element until queue is empty, which is means one level is clear. so I don't have to make that WorkElement struct.

Instead of tracking depth per element, processing one entire level before moving on.

```go
level := 0
for len(queue) > 0 {
    levelSize := len(queue)  // capture current level size
    
    for i := 0; i < levelSize; i++ {
        word := queue[0]
        queue = queue[1:]  // pop
        
        // ... find neighbors, add to queue
    }
    
    level++  // finished one level, increment
}
```

<br>

#### Given a word like "hot", how do you efficiently find all valid transformations in wordList?

How about that, we can make a set as `wordSet` to put all words in to it ar first. and also, we may need a data structure like this:

```go
// key: letter index, value: a rune list for all possible letter char.
var indexLettersMap map[int]map[rune]bool
```

<br>

after we init this indexLettersMap, it will be like:

```
key    -    value
0           [c, d, h, l]
1           [o]
2           [g, t]
```

<br>

when it comes to first word: "hit", we can try to change index-0, index-1, index-2 of it.

* change index-0 could be: cit, oit, git

  none of them exists in wordSet so skip index-0

* change index-1 could only be: hot

  hot exists in wordSet so we hot could be considered into it.

* change index-2 could be: hig, hit

  none of them exists in wordSet so skip index-2

Am I on the right track?

<br>

#### Claude AI feedback

* Question 1: What's the worst case for your indexLettersMap?

  If wordList is large and diverse, each position might have... how many unique letters?

* Question 2: What's the overhead?
  
  * You need O(N × L) time to build indexLettersMap
  * You need extra space for the map

* Question 3: Compare the two approaches asymptotically:

  * Your approach: L × (unique letters at each position)
  * Simple approach: L × 26

When wordList is large, does your optimization save much?


My Suggestion

Your optimization is valid and shows good thinking! But for this problem, the simpler approach (try all 26 letters) is often preferred because:

* 26 is a small constant
* Simpler code, fewer bugs
* The asymptotic complexity is the same: O(L × 26) = O(L)

The real key insight you already have: Use a wordSet (hash set) for O(1) lookup!

<br>
<br>

## Coding

```go
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
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	wordLen := len(beginWord)
	visited := make(map[string]bool) // visited words should not put back into queue again

	level := 0 // BFS level.
	wordQueue := make(WordQueue, 0)

	wordQueue.Push(beginWord)
	visited[beginWord] = true

	for wordQueue.Len() > 0 {
		currentLevelLen := wordQueue.Len()
		level++ // start doing current BFS level, level + 1.
		
		// here we process layer one by one.
		for range currentLevelLen {

			word, _ := wordQueue.Pop()
			for idx := range wordLen {
				// ex: hit -> cog try to change idx0,idx1, idx2
				for _, alp := range alphabet {
					bytes := []byte(word)
					bytes[idx] = alp
					newWord := string(bytes)

					if wordSet[newWord] && !visited[newWord] { // is a valid word in list and not been visited.

						if newWord == endWord {
							return level + 1 // why + 1, cuz new word is belong to next level.
						}

						visited[newWord] = true
						wordQueue.Push(newWord)
					}
				}
			}
		}
	}

	return 0 // not found any path to transform.
}
```

Result:

![1.png](imgs/1.png)

<br>
<br>

## Coding Enhancement

1. if the endWord not exists in wordList, we should just return 0.
2. I creating a new byte slice on every iteration, this should be avoided.
3. Your visited map and wordSet are separate.

   Think about this: Could you use wordSet alone to track both "valid words" AND "visited"? How would that change your code?

<br>

```go
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
```

<br>

Result:

![2.png](imgs/2.png)