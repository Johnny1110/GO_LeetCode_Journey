package combination_sum_II

import "fmt"

type NumQueue struct {
}

func combinationSum2(candidates []int, target int) [][]int {
	return nil
}

// Go call this func in main.go
func Go() {
	candidates := []int{2, 3, 5}
	res := combinationSum2(candidates, 8)
	debugResult(res)
	// [[2,2,2,2],[2,3,3],[3,5]]
}

func debugResult(res [][]int) {
	fmt.Println(res)
}
