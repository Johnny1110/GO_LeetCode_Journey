package test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_1(t *testing.T) {
	target := []int{4, 2, 7, 6, 8}
	sort.Slice(target, func(i, j int) bool { return target[i] > target[j] })
	fmt.Println(target)
}
