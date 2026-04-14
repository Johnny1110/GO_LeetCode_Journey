# 50. Pow(x, n)

<br>

---

<br>

## Coding

```go
func myPow(x float64, n int) float64 {
	sign := n > 0
	n = abs(n)

	x = pow(x, n, 1)
	if !sign {
		return 1 / x
	}
	return x
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func pow(x float64, n int, left float64) float64 {
	if n == 0 {
		return 1 * left
	} else if n == 1 {
		return x * left
	} else if n == 2 {
		return x * x * left
	}

	diff := n / 2
	remaining := n % 2

	if remaining == 1 {
		left *= x
	}

	x *= x

	return pow(x, diff, left)
}
```

<br>
<br>

## Time & Space Complexity

```
Time: O(Log n)

Space: O(Log n) -> because I'm using recursive call. max stack is Log n.
```

<br>
<br>

## Optimized Iterative Version ($O(1)$ Space)

Using for loop instead of call stack.

```go
func myPow(x float64, n int) float64 {
	sign := n > 0
	n = abs(n)

	x = pow(x, n)
	if !sign {
		return 1 / x
	}
	return x
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	left := 1.0
	for n > 1 {
		if n%2 == 1 {
			left *= x // left = 4
		}

		n = n / 2
		x *= x
	}

	return x * left
}
```