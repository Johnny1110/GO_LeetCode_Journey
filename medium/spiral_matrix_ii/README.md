# Spiral Matrix II

<br>

---

<br>

link: https://leetcode.com/problems/spiral-matrix-ii/description/

<br>

## Thinking

### Guess Topic

* simulate

<br>

Just like I did in previous 54. Spiral Matrix.

fill the matrix step by step, we need a cursor pointing at the next position(x, y) and put next number into this position.
something like that I guess.

Direction + Visited Check.

## Coding

```go
type direction = uint8

const (
	RIGHT direction = iota
	DOWN
	LEFT
	UP
)

type cursor struct {
	direct     direction
	currentRow int
	currentCol int

	maxStep     int
	currentStep int

	topBoundary    int
	rightBoundary  int
	bottomBoundary int
	leftBoundary   int
}

// move cursor index to next position and return true if has next.
func (c *cursor) next() bool {
	if c.currentStep == c.maxStep-1 {
		return false
	}

	switch c.direct {
	case RIGHT:
		// if reached boundary.
		if c.currentCol >= c.rightBoundary {
			c.changeDirection()
			return c.next()
		} else { // not reached boundary
			c.currentCol++
			c.currentStep++
			return true
		}
	case DOWN:
		// if reached boundary.
		if c.currentRow >= c.bottomBoundary {
			c.changeDirection()
			return c.next()
		} else { // not reached boundary
			c.currentRow++
			c.currentStep++
			return true
		}
	case LEFT:
		// if reached boundary.
		if c.currentCol <= c.leftBoundary {
			c.changeDirection()
			return c.next()
		} else { // not reached boundary
			c.currentCol--
			c.currentStep++
			return true
		}
	case UP:
		// if reached boundary.
		if c.currentRow <= c.topBoundary {
			c.changeDirection()
			return c.next()
		} else { // not reached boundary
			c.currentRow--
			c.currentStep++
			return true
		}
	default:
		panic("unreachable")
	}
}

func (c *cursor) changeDirection() {
	switch c.direct {
	case RIGHT:
		c.direct = DOWN
		c.rightBoundary--
		break
	case DOWN:
		c.direct = LEFT
		c.bottomBoundary--
		break
	case LEFT:
		c.direct = UP
		c.leftBoundary++
		break
	case UP:
		c.direct = RIGHT
		c.topBoundary++
		break
	}
}

func createMatrixCursor(layer int) *cursor {
	return &cursor{
		maxStep:        layer * layer,
		currentRow:     0,
		currentCol:     0,
		direct:         RIGHT,
		topBoundary:    1,
		leftBoundary:   0,
		rightBoundary:  layer - 1,
		bottomBoundary: layer - 1,
	}
}

// main func
func generateMatrix(n int) [][]int {
	number := 1

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	c := createMatrixCursor(n)

	// position (0, 0)
	matrix[c.currentRow][c.currentCol] = number
	number += 1

	for c.next() {
		matrix[c.currentRow][c.currentCol] = number
		number += 1
	}

	return matrix
}
```

<br>

Result

![1.png](imgs/1.png)