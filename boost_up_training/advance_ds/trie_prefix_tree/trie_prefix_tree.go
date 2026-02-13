package trie_prefix_tree

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

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
