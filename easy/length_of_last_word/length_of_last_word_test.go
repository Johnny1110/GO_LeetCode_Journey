package length_of_last_word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: s = "Hello World"
	//Output: 5
	//Explanation: The last word is "World" with length 5.
	s := "Hello World"
	r := lengthOfLastWord(s)
	assert.Equal(t, 5, r)
}

func Test_2(t *testing.T) {
	//Input: s = "   fly me   to   the moon  "
	//Output: 4
	//Explanation: The last word is "moon" with length 4.
	s := "   fly me   to   the moon  "
	r := lengthOfLastWord(s)
	assert.Equal(t, 4, r)
}

func Test_3(t *testing.T) {
	//Input: s = "luffy is still joyboy"
	//Output: 6
	//Explanation: The last word is "joyboy" with length 6.
	s := "luffy is still joyboy"
	r := lengthOfLastWord(s)
	assert.Equal(t, 6, r)
}

func Test_4(t *testing.T) {
	s := "luffy"
	r := lengthOfLastWord(s)
	assert.Equal(t, 5, r)
}
