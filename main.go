package main

import "fmt"

func main() {
	n := 4

	// x, y := toMatrix(n, 2, 2)
	// fmt.Printf("x=%v, y = %v \n", x, y)

	row, col := toArrayIdx(n, 2, 2)
	fmt.Printf("row=%v, col = %v \n", row, col)
}

func toMatrix(n, row, col int) (x, y int) {
	mid := n / 2
	x = col - mid
	y = -(row - mid)
	return
}

func toArrayIdx(n, x, y int) (row, col int) {
	mid := n / 2
	col = x + mid
	row = mid - y
	return
}
