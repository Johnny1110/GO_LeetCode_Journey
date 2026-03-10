package heap

import (
	"fmt"
	"testing"
)

func Test_Heap(t *testing.T) {
	H := NewHeap[int](func(a, b int) bool { return a < b })

	H.Push(1)
	fmt.Printf("Heap after PUSH(%v): %v \n", 3, H)
	H.Push(4)
	fmt.Printf("Heap after PUSH(%v): %v \n", 2, H)
	H.Push(2)
	fmt.Printf("Heap after PUSH(%v): %v \n", 1, H)
	H.Push(3)
	fmt.Printf("Heap after PUSH(%v): %v \n", 3, H)
	H.Push(9)
	fmt.Printf("Heap after PUSH(%v): %v \n", 9, H)

	fmt.Println("========= ans =========")
	for H.Len() > 0 {
		fmt.Println(H.Pop())
		fmt.Printf("Heap after Pop(): %v \n", H)
	}
}
