package longest_common_prefix

import (
	"go_leetcode_journey/common"
	"strings"
)

// TrieNode Define the TrieNode struct
type TrieNode struct {
	subLeaves map[rune]*TrieNode
}

// InitializeTrieNode creates and returns a new TrieNode
func InitializeTrieNode() *TrieNode {
	return &TrieNode{
		subLeaves: make(map[rune]*TrieNode),
	}
}

// Insert function to insert a word into the Trie
func Insert(root *TrieNode, word string) {
	currentNode := root
	for _, char := range word {
		if _, found := currentNode.subLeaves[char]; !found {
			currentNode.subLeaves[char] = InitializeTrieNode()
		}
		currentNode = currentNode.subLeaves[char]
	}
}

func getResult(root *TrieNode) string {
	currentNode := root
	result := strings.Builder{}
	for {
		if len(currentNode.subLeaves) == 1 {
			for char, leaf := range currentNode.subLeaves {
				result.WriteString(string(char))
				currentNode = leaf
			}
		} else {
			break
		}
	}
	return result.String()
}

func longestCommonPrefix(strs []string) string {
	root := InitializeTrieNode()
	for _, val := range strs {
		if val == "" {
			return val
		}
		Insert(root, val)
	}
	return getResult(root)
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl")
}
