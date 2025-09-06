package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	assert.Equal(t, 28, uniquePaths(3, 7))
}

func Test_2(t *testing.T) {
	assert.Equal(t, 3, uniquePaths(3, 2))
}
