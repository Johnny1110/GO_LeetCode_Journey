package word_ladder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
	// Output: 5
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	assert.Equal(t, 5, ladderLength(beginWord, endWord, wordList))
}

func Test_2(t *testing.T) {
	// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
	// Output: 0
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log"}

	assert.Equal(t, 0, ladderLength(beginWord, endWord, wordList))
}
