package permutation_sequence

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	used := make([]bool, n+1)
	result := &[]int{}

	// 1. convert k to k_idx:
	kIdx := k - 1

	collect(used, kIdx, n, result)

	var sb strings.Builder
	for _, v := range *result {
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}

func collect(used []bool, kIdx int, n int, result *[]int) {
	if n == 0 {
		return
	}
	// 2. choose group
	eachGpCount := factorial(n - 1)
	targetGp := chooseGpIdx(kIdx, eachGpCount)

	// 3. put int into result.
	target := findIdxAndMarkUsed(used, targetGp)
	//fmt.Println("========================================")
	//fmt.Println("n:", n, "eachGpCount:", eachGpCount, "kIdx:", kIdx)
	//fmt.Println("used:", used, "targetGp:", targetGp)
	//fmt.Println("find target:", target)
	*result = append(*result, target)

	// 4. upgrade kIdx and n
	kIdx = upgradeKIdx(kIdx, eachGpCount)

	//fmt.Println("current result:", result)

	collect(used, kIdx, n-1, result)
}

func factorial(n int) int {
	accumulator := n
	for i := n - 1; i > 0; i-- {
		accumulator *= i
	}
	return accumulator
}

func chooseGpIdx(kIdx, eachGpCount int) int {
	if eachGpCount == 0 {
		return 0
	}
	// formula: gpIdx = kIdx / eachGpCount
	return kIdx / eachGpCount
}

func upgradeKIdx(kIdx, eachGpCount int) int {
	if eachGpCount == 0 {
		return 0
	}
	return kIdx % eachGpCount
}

func findIdxAndMarkUsed(used []bool, idx int) int {
	count := 0
	// used[0] is not include.
	for i := 1; i < len(used); i++ {
		if !used[i] {
			count++
			if idx == count-1 {
				used[i] = true
				return i
			}
		}
	}
	return -1
}
