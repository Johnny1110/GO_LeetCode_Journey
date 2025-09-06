package permutation_sequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	ans := getPermutation(3, 3)
	assert.Equal(t, "213", ans)
}

func Test_2(t *testing.T) {
	ans := getPermutation(4, 9)
	assert.Equal(t, "2314", ans)
}

func Test_3(t *testing.T) {
	ans := getPermutation(3, 1)
	assert.Equal(t, "123", ans)
}

func Test_factorial(t *testing.T) {
	assert.Equal(t, 120, factorial(5))
	assert.Equal(t, 6, factorial(3))
	assert.Equal(t, 1, factorial(1))
	assert.Equal(t, 0, factorial(0))
}

func Test_chooseGpIdx(t *testing.T) {
	assert.Equal(t, 1, chooseGpIdx(2, 2))
	assert.Equal(t, 0, chooseGpIdx(1, 2))
	assert.Equal(t, 1, chooseGpIdx(3, 2))
	assert.Equal(t, 2, chooseGpIdx(4, 2))
	assert.Equal(t, 2, chooseGpIdx(5, 2))
}

func Test_upgradeKIdx(t *testing.T) {
	assert.Equal(t, 0, upgradeKIdx(0, 2))
	assert.Equal(t, 1, upgradeKIdx(1, 2))
	assert.Equal(t, 0, upgradeKIdx(2, 2))
	assert.Equal(t, 1, upgradeKIdx(3, 2))
	assert.Equal(t, 0, upgradeKIdx(4, 2))
	assert.Equal(t, 1, upgradeKIdx(5, 2))
}

func Test_findIdx(t *testing.T) {
	used := []bool{false, false, true, false}
	assert.Equal(t, 1, findIdxAndMarkUsed(used, 0))
	used = []bool{false, false, true, false}
	assert.Equal(t, 3, findIdxAndMarkUsed(used, 1))
}
