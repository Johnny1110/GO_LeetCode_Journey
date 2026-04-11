# 48. Rotate Image

<br>

---

<br>

## Think

In Liner Algebra, basis vector is `[[1, 0] [0, 1]]`

Right Rotate M = `[[0, 1], [-1, 0]]`

any vector `[x, y]` we want to calculate right rotate new vector:

```
// vector * M
[x, y] * [[0, 1], [-1, 0]]
= [y, -x]
```

<br>

So we know if any vector `[x, y]` want to right rotate 90 dgree is `[y, -x]`

<br>

* First element in 2D array is (0, 0)
* convert it to vector is `[-2, 2]`
* right rotate `vector * M = [2, 2]`
* `[2, 2]` convert to 2D array index = (0, 2)

<br>
<br>

### Assume N = 4

listing some answer:

```
(row, col)  -> (row*, col*)
(0, 0)         (0, 3)
(0, 1)         (1, 3)
(0, 2)         (2, 3)
(0, 3)         (3, 3)
(1, 0)         (0, 2)
(1, 1)         (0, 1)
```

<br>

formula:

```
row* = col
col* = n-1-row
```

<br>
<br>

### This conversion between matrix index and math vector is the key insight.

```
(row, col) → (x, y) = ( ? , ? )
```

<br>
<br>

## Coding

```go
func rotate(matrix [][]int) {
	n := len(matrix) // n = 4
	if n == 1 {
		return
	}

	for row := 0; row < n/2; row++ {
		for col := row; col < n-1-row; col++ {
			temp := matrix[row][col]

			// P1 = [row][col]
			// P2 = [col][n-1-row]
			// P3 = [n-1-row][n-1-col]
			// P4 = [n-1-col][row]. -> why row? cuz (n-1-(n-1-row)) = row

			matrix[row][col] = matrix[n-1-col][row]         // P1 = P4
			matrix[n-1-col][row] = matrix[n-1-row][n-1-col] // P4 = P3
			matrix[n-1-row][n-1-col] = matrix[col][n-1-row] // P3 = P2
			matrix[col][n-1-row] = temp                     //P2 = temp
		}
	}
}
```

<br>
<br>

## Time & Space Complexity

```
Time: O(n square) -> 2 nested loop

Space: O(1) -> no extra space needed
```