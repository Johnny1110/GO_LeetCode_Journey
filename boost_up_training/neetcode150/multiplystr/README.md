# 43. Multiply Strings

<br>

---

<br>

## Coding

```go

var ct int = 48

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]int, len(num1)+len(num2))
	num1Arr := make([]int, len(num1))
	num2Arr := make([]int, len(num2))

	for i, c := range num1 {
		num1Arr[i] = int(c) - ct
	}
	for i, c := range num2 {
		num2Arr[i] = int(c) - ct
	}

	for i := len(num1Arr) - 1; i >= 0; i-- {
		for j := len(num2Arr) - 1; j >= 0; j-- {
			tmp := num1Arr[i] * num2Arr[j]

			affectIdx := i + j + 1
			tmp = result[affectIdx] + tmp

			this := tmp % 10
			next := tmp / 10

			result[affectIdx] = this
			for next > 0 {
				affectIdx -= 1
				tmp = result[affectIdx] + next
				this = tmp % 10
				next = tmp / 10
				result[affectIdx] = this
			}
		}
	}

	// find first non-zero
	startIdx := 0
	for i := 0; i < len(result); i++ {
		if result[i] > 0 {
			startIdx = i
			break
		}
	}

	ans := ""
	for i := startIdx; i < len(result); i++ {
		tmp := rune(result[i] + ct)
		ans += string(tmp)
	}

	return ans
}
```

<br>
<br>

## Time Space Complexity

```
Assume: n = len of num1, m = len of num2

Time: O(m*n) -> convert num1 num2 string to int cost O(n+m), calculate result array cost O(n*m), dump to answer string O(x) (x is result length)

Space: O(m+n)
```

<br>
<br>

## Revamp Version

```go
var ct int = 48

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {

			n1 := int(num1[i]) - ct
			n2 := int(num2[j]) - ct

			result[i+j+1] += n1 * n2
			result[i+j] += result[i+j+1] / 10
			result[i+j+1] %= 10
		}
	}

	// find first non-zero
	startIdx := 0
	for i := 0; i < len(result); i++ {
		if result[i] > 0 {
			startIdx = i
			break
		}
	}

	ans := ""
	for i := startIdx; i < len(result); i++ {
		tmp := rune(result[i] + ct)
		ans += string(tmp)
	}

	return ans
}
```