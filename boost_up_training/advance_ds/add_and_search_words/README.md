# 211. Design Add and Search Words

<br>

---

<br>

link: https://leetcode.com/problems/design-add-and-search-words-data-structure/description/

<br>

Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

* `WordDictionary()` Initializes the object.
* `void addWord(word)` Adds word to the data structure, it can be matched later.
* `bool search(word)` Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.


Example:

```
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]

[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]

Output
[null,null,null,null,false,true,true,true]
```

Explanation
```
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True
```

<br>
<br>

## Thinking

I think this problem is very alike prefix tree problem (# 208), the different part is this `search()` support wildcard `'.'` now.

The critical problem is in the recursive call `search` when the next letter comes to `'.'`, we suppose to using DFS to explore all not nil branches (max to 26).   

<br>
<br>

## Coding

```go
type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: NewNode(0),
	}
}

func (this *WordDictionary) AddWord(word string) {
	if word == "" {
		this.root.isEnd = true
	}

	this.root.insert(word, 0)
}

func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		return this.root.isEnd
	}

	return this.root.search(word, 0)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

 type Node struct {
	letter byte
	branch [26]*Node
	isEnd bool
 }

 func NewNode(letter byte) *Node {
	return &Node {
		letter: letter,
	}
 }

 func (n *Node) insert(word string, wordIdx int) {
	ltr := word[wordIdx]
	branchIdx := ltr-97


	if n.branch[branchIdx] == nil {
		n.branch[branchIdx] = NewNode(ltr)
	}

	if wordIdx+1 == len(word) {
		n.branch[branchIdx].isEnd = true
	} else {
		n.branch[branchIdx].insert(word, wordIdx+1)
	}
 }

 func (n *Node) search(word string, wordIdx int) bool {
	ltr := word[wordIdx]
	isLastWordIdx := wordIdx+1 == len(word) 


	if ltr != '.' {
		// normal search
		branchIdx := ltr-97


		if n.branch[branchIdx] != nil { // found in branch


			if isLastWordIdx {
				return n.branch[branchIdx].isEnd
			} else {
				return n.branch[branchIdx].search(word, wordIdx+1)
			}
		} else {
			return false // no more letter in trie
		}


	} else {
		// explore all branch

		if isLastWordIdx {
			for _, bn := range n.branch {
				if bn != nil && bn.isEnd {
					return true
				}
			}
			return false
		}

		for _, bn := range n.branch {
			
			if bn != nil {// this branch has val
				if bn.search(word, wordIdx+1) {
					return true
				}
			}
		}

		return false
	}
 }
```

<br>

Result:

![1](imgs/1.png)