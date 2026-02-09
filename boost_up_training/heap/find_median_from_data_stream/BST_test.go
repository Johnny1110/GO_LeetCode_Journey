package find_median_from_data_stream

import (
	"fmt"
	"testing"
)

func Test_BST(t *testing.T) {
	root := NewNode(2)

	root.Push(1)
	root.Push(3)
	root.Push(4)
	root.Push(5)
	root.Push(6)

	fmt.Printf("ROOT size = %v \n", root.Size)
	fmt.Printf("median = %v\n", root.FindMedian())
}
