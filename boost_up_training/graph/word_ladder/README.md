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

