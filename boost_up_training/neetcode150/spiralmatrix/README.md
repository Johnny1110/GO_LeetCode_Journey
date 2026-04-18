# 54. Spiral Matrix

<br>

---

<br>

## Coding

```go
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	if m == 1 {
		return matrix[0]
	}

	direction := RIGHT
	leftBoundary := 0
	rightBoundary := n - 1
	upBoundary := 0
	downBoundary := m - 1

	cnt := m * n
	result := make([]int, cnt)

	row, col := 0, 0
	for idx := 0; idx < cnt; idx++ {
		// collect this position
		result[idx] = matrix[row][col]
		// change direction if needed

		switch direction {
		case RIGHT:
			if col == rightBoundary {
				direction = DOWN
				upBoundary++
			}
		case LEFT:
			if col == leftBoundary {
				direction = UP
				downBoundary--
			}
		case DOWN:
			if row == downBoundary {
				direction = LEFT
				rightBoundary--
			}
		case UP:
			if row == upBoundary {
				direction = RIGHT
				leftBoundary++
			}
		}

		// update next row, col
		switch direction {
		case RIGHT:
			col++
		case LEFT:
			col--
		case UP:
			row--
		case DOWN:
			row++
		}
	}

	return result
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)
```

<br>
<br>

## Time & Space Complexity

```
Assume: N = total row*col

Time: O(N)

Space: O (1)
```