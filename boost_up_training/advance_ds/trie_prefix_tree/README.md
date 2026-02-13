# 208. Trie/Prefix Tree

<br>

---

<br>

link: https://leetcode.com/problems/lru-cache/description/

<br>
<br>

A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. 
There are various applications of this data structure, such as autocomplete and spellchecker.

<br>

Implement the Trie class:

* `Trie()` Initializes the trie object.
* `void insert(String word)` Inserts the string word into the trie.
* `boolean search(String word)` Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
* `boolean startsWith(String prefix)` Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.


Example 1:

```
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]

Output
[null, null, true, false, true, null, true]
```

Explanation:

```
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True
```

<br>
<br>


## Thinking

I used to made a MPT like early of 2025, I know this problem is all about patrica tree.

I will do it like what I used to do.

<br>

## Trie Design

Since the problem not asking for delete, it will be easy.

This tree will be a bunch of nodes like this:

```go
type Node struct {
val    byte    // store half letter
branch []*Node // len is fixed 16 (idx: 0~15)
}
```

<br>

When we want to insert a word into this trie, we need split the word into a slice, each letter will split to 2 bytes also.

like a word is `[0xA1, 0xA2]` will be split to `[0x0A, 0x01, 0x0A, 0x02]`

<br>

Define a `dummyHead` node as trie head, it will not contains any val, it just keep a 16 length `branch`.

When we want to put `[0x0A, 0x01, 0x0A, 0x02]` into that trie.

* try to find `0x0A` -> 10th position in `dummyHead.branch` which is `idx = 9`
  * found it: pass remains to it `[0x01, 0x0A, 0x02]` (recursive insert call)
  * not found: create node with `0x0A` as `val`, put it into `dummyHead.branch[0x0A]`, also pass remain to it.

* recursive insert logic until no remain data left.

<br>

Search is very simple like insert recursive search until no bytes remaining.

<br>

### Claude feedback

#### This Works, But Consider the Problem Constraints

Leetcode 208 states: inputs consist of only lowercase English letters (a-z). That's just 26 possible characters.

Your approach splits each byte into two nibbles and uses a 16-branch node. This works, but think about the tradeoff:

* Your approach: each character requires traversing 2 nodes (two nibbles), branch size 16
* Standard trie: each character requires traversing 1 node, branch size 26

For this specific problem, the standard approach is simpler and more direct:

```go
type TrieNode struct {
    children [26]*TrieNode
    isEnd    bool
}
```

#### Your Patricia Trie Approach

It's not wrong — it's actually more space-efficient for general byte data. 
But for 26 lowercase letters, it adds unnecessary complexity and doubles the tree depth.

<br>

### Refine

I will adopt Claude's suggestion, and using 26 length branch, and isEnd mark.

<br>
<br>

## Coding

```go
type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			isEnd: false,
		},
	}
}

func (this *Trie) Insert(word string) {
	this.root.insert(word, 0)
}

func (this *Trie) Search(word string) bool {
	return this.root.search(word, 0)
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.root.startWith(prefix, 0)
}

type Node struct {
	val    byte      // store letter
	branch [26]*Node // len is fixed 26 (idx: 0~25)
	isEnd  bool
}

// insert a=97 z=122
func (n *Node) insert(word string, idx int) {
	letter := word[idx]
	letterIdx := letter - 97

	if n.branch[letterIdx] == nil {
		node := &Node{
			val: letter,
		}
		n.branch[letterIdx] = node
	}

	if idx+1 == len(word) {
		// this is the end.
		n.branch[letterIdx].isEnd = true
		return
	}

	n.branch[letterIdx].insert(word, idx+1)
}

func (n *Node) startWith(prefix string, idx int) bool {
	letter := prefix[idx]
	letterIdx := letter - 97

	if n.branch[letterIdx] == nil {
		return false
	}

	if idx+1 == len(prefix) {
		return true
	}

	return n.branch[letterIdx].startWith(prefix, idx+1)
}

func (n *Node) search(word string, idx int) bool {
	letter := word[idx]
	letterIdx := letter - 97

	if n.branch[letterIdx] == nil {
		return false
	}

	if idx+1 < len(word) {
		return n.branch[letterIdx].search(word, idx+1)
	}

	if idx+1 == len(word) && n.branch[letterIdx].isEnd {
		return true
	}

	return false
}
```

Result:

![1.png](imgs/1.png)