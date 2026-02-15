package missing_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumN(t *testing.T) {

	res := sumN(4)
	assert.Equal(t, 10, res)

	res = sumN(3)
	assert.Equal(t, 6, res)
}
