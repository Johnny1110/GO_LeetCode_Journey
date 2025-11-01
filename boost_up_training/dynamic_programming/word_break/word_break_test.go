package word_break

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: s = "leetcode", wordDict = ["leet","code"]
	//Output: true
	//Explanation: Return true because "leetcode" can be segmented as "leet code".
	wordDict := []string{"leet", "code"}
	assert.True(t, wordBreak("leetcode", wordDict))
}

func Test_2(t *testing.T) {
	wordDict := []string{"apple", "pen"}
	assert.True(t, wordBreak("applepenapple", wordDict))
}

func Test_3(t *testing.T) {
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	assert.False(t, wordBreak("catsandog", wordDict))
}
