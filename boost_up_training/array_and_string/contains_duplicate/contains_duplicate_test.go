package contains_duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	assert.True(t, containsDuplicate([]int{1, 2, 3, 1}))
}

func Test_2(t *testing.T) {
	assert.False(t, containsDuplicate([]int{1, 2, 3, 4}))
}

func Test_3(t *testing.T) {
	assert.True(t, containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
