package insert_interval

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_insert_1(t *testing.T) {
	intervals_1 := [][]int{{1, 3}, {6, 9}}
	res_1 := insert(intervals_1, []int{2, 5})
	fmt.Println("res_1:", res_1)

	expect := [][]int{{1, 5}, {6, 9}}
	assert.Equal(t, expect, res_1)
}

func Test_insert_2(t *testing.T) {
	intervals_2 := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	res_2 := insert(intervals_2, []int{4, 8})
	fmt.Println("res_2:", res_2)

	expect := [][]int{{1, 2}, {3, 10}, {12, 16}}
	assert.Equal(t, expect, res_2)
}

func Test_insert_3(t *testing.T) {
	intervals_2 := [][]int{}
	res_2 := insert(intervals_2, []int{5, 7})
	fmt.Println("res_2:", res_2)

	expect := [][]int{{5, 7}}
	assert.Equal(t, expect, res_2)
}

func Test_insert_4(t *testing.T) {
	intervals_2 := [][]int{{1, 5}}
	res_2 := insert(intervals_2, []int{2, 3})
	fmt.Println("res_2:", res_2)

	expect := [][]int{{1, 5}}
	assert.Equal(t, expect, res_2)
}

func Test_insert_5(t *testing.T) {
	intervals_2 := [][]int{{1, 5}}
	res_2 := insert(intervals_2, []int{2, 7})
	fmt.Println("res_2:", res_2)

	expect := [][]int{{1, 7}}
	assert.Equal(t, expect, res_2)
}

func Test_insert_6(t *testing.T) {
	intervals_2 := [][]int{{1, 5}}
	res_2 := insert(intervals_2, []int{2, 3})
	fmt.Println("res_2:", res_2)

	expect := [][]int{{1, 5}}
	assert.Equal(t, expect, res_2)
}
