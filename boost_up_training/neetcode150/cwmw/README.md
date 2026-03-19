# 11. Container With Most Water

<br>

---

<br>

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

<br>
<br>

## Coding

```go
func maxArea(height []int) int {
	bestAns := 0

	pointerA, pointerB := 0, len(height)-1

	for pointerA < pointerB {
		A, B := height[pointerA], height[pointerB]
		w, h := pointerB-pointerA, min(A, B)
		bestAns = max(bestAns, w*h)

		if A > B {
			pointerB--
		} else if A < B {
			pointerA++
		} else {
			pointerB--
			pointerA++
		}
	}

	return bestAns

}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = length of height array

Time: O(n) ~ O(n/2)

Space: O(1)
```