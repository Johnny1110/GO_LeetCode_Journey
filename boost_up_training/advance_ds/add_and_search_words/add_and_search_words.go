package add_and_search_words

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