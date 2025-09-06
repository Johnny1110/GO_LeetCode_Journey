# 7. Reverse Integer

<br>

---

<br>

## Thinking

...

<br>
<br>

## Coding

<br>

```go
const (
	MAX_INT = 2<<30 - 1
	MIN_INT = -2 << 30
)

func reverse(x int) int {

	isNegative := x < 0
	result := 0
	x = abs(x)

	for {
		remainder := x % 10
		x = x / 10
		result = result*10 + remainder
		if x == 0 {
			break
		}
	}

	if isNegative {
		result = result * -1
	}

	if result > MAX_INT || result < MIN_INT {
		return 0
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

<br>
<br>