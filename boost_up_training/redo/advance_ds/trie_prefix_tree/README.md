# 208. Trie/Prefix Tree

```go
type Trie struct {
	head *Node
}

func Constructor() Trie {
	return Trie{
		head: NewNode(),
	}
}

func (this *Trie) Insert(word string) {
	if word == "" {
		this.head.isEnd = true
	}

	this.head.insert(0, word)
}

func (this *Trie) Search(word string) bool {
	if word == "" {
		return this.head.isEnd
	}

	return this.head.search(0, word)
}

func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}

	return this.head.startWith(0, prefix)
}

// ------------------------------------------

type Node struct {
	branch [26]*Node
	isEnd  bool
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) insert(idx int, word string) {
	if idx == len(word) {
		n.isEnd = true
		return
	}

	char := word[idx]
	node := n.branch[getIdx(char)]

	if node == nil {
		node = NewNode()
		n.branch[getIdx(char)] = node
	}

	node.insert(idx+1, word)
}

func (n *Node) search(idx int, word string) bool {
	if idx == len(word) {
		return n.isEnd
	}

	char := word[idx]
	node := n.branch[getIdx(char)]

	if node == nil {
		return false
	} else {
		return node.search(idx+1, word)
	}
}

func (n *Node) startWith(idx int, word string) bool {
	if idx == len(word) {
		return true
	}

	char := word[idx]
	node := n.branch[getIdx(char)]

	if node == nil {
		return false
	} else {
		return node.startWith(idx+1, word)
	}
}

func getIdx(char uint8) int {
	return int(char) - 97
}
```