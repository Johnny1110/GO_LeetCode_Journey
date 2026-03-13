package addandsearchwords

type WordDictionary struct {
	head *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		head: &Node{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	if word == "" {
		this.head.isEnd = true
		return
	}

	this.head.add(0, word)
}

func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		return this.head.isEnd
	}

	return this.head.search(0, word)
}

// ------------------------------------------
func getBranchIdx(char uint8) int {
	return int(char) - 97
}

type Node struct {
	branch [26]*Node
	isEnd  bool
}

func (n *Node) add(idx int, word string) {
	if idx == len(word) {
		n.isEnd = true
		return
	}

	char := word[idx]
	branchIdx := getBranchIdx(char)
	node := n.branch[branchIdx]

	if node == nil {
		node = &Node{}
		n.branch[branchIdx] = node
	}

	node.add(idx+1, word)
}

func (n *Node) search(idx int, word string) bool {
	if idx == len(word) {
		return n.isEnd
	}

	char := word[idx]

	if char == '.' {
		// wildcard
		for _, node := range n.branch {
			if node != nil {
				if node.search(idx+1, word) {
					return true
				}
			}
		}

		return false
	} else {
		branchIdx := getBranchIdx(char)
		node := n.branch[branchIdx]
		if node == nil {
			return false
		} else {
			return node.search(idx+1, word)
		}
	}
}
