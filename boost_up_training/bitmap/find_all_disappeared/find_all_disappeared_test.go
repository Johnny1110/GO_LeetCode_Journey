package find_all_disappeared

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	res := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Printf("res: %v \n", res)
}
