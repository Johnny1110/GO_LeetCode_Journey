package coinchangeii

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	amt := change(5, []int{1, 2, 5})
	fmt.Printf("amt: %v \n", amt)
}
