# 136. Single Number

<br>

---

<br>

## Coding

```go
func singleNumber(nums []int) int {
	ans := 0
	for _, n := range nums {
		ans ^= n
	}
	return ans
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: N = len(nums)

Time: O(N)
Space: O(1)
```
