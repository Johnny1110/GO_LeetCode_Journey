package reverse_integer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
}

func Test_2(t *testing.T) {
	assert.Equal(t, -123, reverse(-321))
	assert.Equal(t, 21, reverse(120))
}

func Test_3(t *testing.T) {
	assert.Equal(t, 0, reverse(1534236469))
}

func Test_multiple(t *testing.T) {
	fmt.Println(321 * -1)
}

func Test_Shifting(t *testing.T) {
	res := leftShifting(2, 6)
	fmt.Println("res: ", res)
	res = leftShifting(2, 5)
	fmt.Println("res: ", res)
	res = leftShifting(2, 4)
	fmt.Println("res: ", res)
	res = leftShifting(2, 3)
	fmt.Println("res: ", res)
	res = leftShifting(2, 2)
	fmt.Println("res: ", res)
	res = leftShifting(2, 1)
	fmt.Println("res: ", res)
	res = leftShifting(2, 0)
	fmt.Println("res: ", res)
}
