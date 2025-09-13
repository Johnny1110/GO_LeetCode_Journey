package valid_number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	assert.True(t, isNumber("0"))
}

func Test_2(t *testing.T) {
	assert.False(t, isNumber("e"))
}

func Test_3(t *testing.T) {
	assert.False(t, isNumber("."))
}

func Test_common(t *testing.T) {
	topic := ".+-eE"
	for _, a := range topic {
		fmt.Println(a)
	}
	// . : 46
	// + : 43
	// - : 45
	// e : 101
	// E : 69
}
