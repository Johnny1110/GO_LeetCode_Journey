package edit_distance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: word1 = "horse", word2 = "ros"
	//Output: 3
	//Explanation:
	//	horse -> rorse (replace 'h' with 'r')
	//	rorse -> rose (remove 'r')
	//	rose -> ros (remove 'e')
	assert.Equal(t, 3, minDistance("horse", "ros"))
}

func Test_2(t *testing.T) {
	assert.Equal(t, 5, minDistance("intention", "execution"))
}
