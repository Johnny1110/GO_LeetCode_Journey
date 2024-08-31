package generate_parentheses

import (
	"fmt"
	"go_leetcode_journey/common"
)

func generateParenthesis(n int) []string {
	// init left parentheses & right parentheses.
	left := n
	right := n
	ans := make([]string, 0)

	backTracking(&ans, "", left, right)

	for _, an := range ans {
		fmt.Println(an)
	}

	return ans
}

func backTracking(ans *[]string, str string, left, right int) {
	if right < left || left < 0 {
		return
	}

	if left == 0 && right == 0 {
		*ans = append(*ans, str)
		return
	}

	// 2 分法
	backTracking(ans, str+"(", left-1, right)
	backTracking(ans, str+")", left, right-1)

}

// Go call this func in main.go
func Go() {
	common.Assert_answer(generateParenthesis(3), nil)
}
