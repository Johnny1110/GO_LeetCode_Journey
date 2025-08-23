# Zigzag Conversion

<br>

---

<br>

## link

https://leetcode.com/problems/zigzag-conversion/description/

<br>

## Brain storm

skip

<br>

## Coding

```go
// zigzagFormula y = 2x-2 -> return y
func zigzagFormula(x int) int {
	return 2*x - 2
}

func diff(i, j int) int {
	if i > j {
		return j - i
	} else {
		return j - i
	}
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	// 0. init return string
	finalResult := make([]uint8, 0, len(s))

	// 1. process box list
	formulaY := zigzagFormula(numRows)

	// 0 ~ 3
	for i := 0; i < numRows; i++ {

		// 2-1 add header
		if i >= len(s) {
			break
		}
		finalResult = append(finalResult, s[i])

		// 2-2 adapt formula number with input s.
		numA := formulaY - 2*i
		numB := diff(numA, formulaY)

		// 2-3 iterate input s
		tempIdx := i
		for tempIdx < len(s) {
			// 抓

			if numA > 0 {
				tempIdx += numA
				if tempIdx >= len(s) {
					break
				} else {
					finalResult = append(finalResult, s[tempIdx])
				}
			}

			// 剩
			if numB > 0 {
				tempIdx += numB
				if tempIdx >= len(s) {
					break
				} else {
					finalResult = append(finalResult, s[tempIdx])
				}
			}
		}

	}

	return string(finalResult)
}
```