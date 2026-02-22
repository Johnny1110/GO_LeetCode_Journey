# 11. Container With Most Water

<br>

---

<br>

link: https://leetcode.com/problems/container-with-most-water/description

<br>
<br>

## Coding-1

```go
func maxArea(height []int) int {
    
	pointerA, pointerB := 0, len(height)-1
	tmp := 0

	for pointerA < pointerB {
		w := pointerB-pointerA
		h := min(height[pointerA], height[pointerB]) // using smaller one
		tmp = max(tmp, w*h)

		// moving smaller one
		if height[pointerA] > height[pointerB] {
			pointerB--
		} else {
			pointerA++
		}
	}

	return tmp
}
```