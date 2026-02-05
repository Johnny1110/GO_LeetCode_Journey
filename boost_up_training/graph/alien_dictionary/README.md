# 269. Alien Dictionary

<br>

---

<br>

link: https://neetcode.io/problems/foreign-dictionary/question

<br>

## Example

```
Input: ["hrn","hrf","er","enn","rfnn"]

Output: "hernf"
```

Explanation:

* from "hrn" and "hrf", we know 'n' < 'f'
* from "hrf" and "er", we know 'h' < 'e'
* from "er" and "enn", we know get 'r' < 'n'
* from "enn" and "rfnn" we know 'e'<'r'

so one possibile solution is `"hernf"`

<br>
<br>

## Topic 

* Array
* String
* DFS, BFS
* Graph
* [Topological Sort](https://medium.com/@mswukris/%E6%8B%93%E6%92%B2%E6%8E%92%E5%BA%8F%E6%B3%95-topological-sorting-4a727d851c62)

<br>
<br>

## Thinking

### Deconstruct Problem

I know this problem is about to using topological sort to solve. For example:

```
Input: ["hrn","hrf","er","enn","rfnn"]
```

I guess I have to create a kind of node graph, and logic is like:

1. load "hrn"

    ```go
    h -> r
    r -> n
    ```
   
2. load "hrf"

    ```
    r -> f
    ```
   
3. load "er"

   ```
    e -> r
   ```

4. load "enn"

    ```
   e -> n
   ```
   
5. load "rfnn"

   ```
   f -> n
   ```

<br>

Now we created a **Directed Graph**, head node is "h" and "e" because both of em are not be pointed by any letter.

we can using a queue (BFS) and a list to collect the result answer.

* round-1:
  * queue: ["h", "e"], result: []
  * pop "h" and push "h"'s pointer node into queue, then push "h" into result
  * queue: ["e", "r"], result: ["h"]
* round-2:
  * queue: ["e", "r"], result: ["h"]
  * pop "e" and push "e"'s pointer node into queue, then push "e" into result
  * queue: ["r", "n"], result: ["h", "e"]
* round-3:
  * queue: ["r", "n"], result: ["h", "e"]
  * pop "r" and push "r"'s pointer node into queue, then push "r" into result
  * queue: ["n", "f"], result: ["h", "e", "r"]
* round-4:
  * queue: ["n", "f"], result: ["h", "e", "r"]
  * pop "n" and push "n"'s pointer node into queue, then push "n" into result
  * queue: ["f"], result: ["h", "e", "r", "n"]
* round-5:
  * queue: ["f"], result: ["h", "e", "r", "n"]
  * pop "f" and push "f"'s pointer node into queue, then push "f" into result
  * queue: [], result: ["h", "e", "r", "n", "f"]

result list to string = "hernf"

<br>

### Claude AI feedback

You're definitely on the right track with using topological sort! But I see a **critical misunderstanding in how you're building the graph**. 
Let me guide you through this.

<br>

#### The Key Insight You're Missing

You're extracting relationships from within each word (h→r→n from "hrn"). But that's not what the problem tells us.
Think about it this way: In a dictionary, words are sorted. What can we learn from two adjacent words being in sorted order?

For example, in English:

* "apple" comes before "banana"
* What does that tell us? Only that 'a' < 'b'
* We learn nothing about p, l, e, n within those words

<br>

#### The Correct Way to Extract Relationships

Compare adjacent words in the list, not characters within a single word.

Given your input: `["hrn", "hrf", "er", "enn", "rfnn"]`

1. Compare "hrn" vs "hrf"
   ```
   h == h  (same, keep going)
   r == r  (same, keep going)
   n != f  → we learn: n < f (n comes before f in alien alphabet)
   ```
   
2. Compare "hrf" vs "er"
   ```
   h != e  → we learn: h < e
   ```

3. Compare "er" vs "enn"
   ```
   e == e  (same, keep going)
   r != n  → we learn: r < n
   ```
   
4. Compare "enn" vs "rfnn"
   ```
   e != r  → we learn: e < r
   ```

<br>

#### Your Task

Based on this corrected understanding:

1. What edges should your directed graph actually have?
2. What would the in-degree of each node be?
3. There's also an edge case to consider: What if a longer word appears before its prefix? (e.g., ["abc", "ab"]) — What should happen?

<br>

#### Sum up

First pass: collect all unique letters from all words
Second pass: compare adjacent words to build edges

We're using Kahn's algorithm (BFS Approach) to solve this problem.

So the first step is creating that **Directed Graph**

We need a `LetterNode` to represent each letter and their `inDegree` count and out-degree `*LetterNode` list.

We also need a `map[uint8]LetterNode` to let us have a letterNodesIndex.

<br>

Iterate through the input `["hrn","hrf","er","enn","rfnn"]`:

1. Compare "hrn","hrf":

   * create 'h' node, it will be the first in-degree 0 node
   * create 'n' and 'f' nodes, 'n' pointing to 'f' and f's inDegree++

2. Compare "hrf", "er":

    * create 'e' node, 'h' pointing to 'e', e's inDegree++

3. and so on...

<br>

Finally, we can get a Directed Graph, iterate through letterNodesIndex to find all 0 inDegree (h) and put em into BFS Queue. (maybe this could be improved)

<br>

#### BFS process

round-1:

```
queue: [h]
tSort: []

pop 'h' add it into tSort, h pointing to e, e's inDegree--, and e's inDegree is 0, put it into queue
```

round-2:

```go
queue: [e]
tSort: [h]

pop 'e' add it into tSort, e pointing to r, r's inDegree--, and r's inDegree is 0, put it into queue
```

round-3:

```go
queue: [r]
tSort: [h, e]

pop 'r' add it into tSort, r pointing to n, n's inDegree--, and n's inDegree is 0, put it into queue
```

round-4:

```go
queue: [n]
tSort: [h, e, r]

pop 'n' add it into tSort, n pointing to f, f's inDegree--, and f's inDegree is 0, put it into queue
```

round-4:

```go
queue: [f]
tSort: [h, e, r, n]

pop 'f' add it into tSort, f pointing to nothing

FINAL:
queue: []
tSort: [h, e, r, n, f]
```

<br>

---

<br>

## Coding

```go

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
```

<br>

This is a leetcode premium level problem, so there is no Result. 