package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: m = 3, n = 7
	// O * * * * * *
	// * * * * * * *
	// * * * * * * X
	//Output: 28
	assert.Equal(t, 28, uniquePaths(3, 7))
}

func Test_2(t *testing.T) {
	//Input: m = 3, n = 2
	//Output: 3
	assert.Equal(t, 3, uniquePaths(3, 2))
}
