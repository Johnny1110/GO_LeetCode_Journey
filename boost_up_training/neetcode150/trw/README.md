# 42. Trapping Rain Water

<br>

---

<br>

## Coding

```go
func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	ans := 0

	prefix := make([]int, len(height))
	suffix := make([]int, len(height))

	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	for right >= 0 {

		prefix[left] = maxLeft
		suffix[right] = maxRight
		lh, rh := height[left], height[right]
		maxLeft = max(maxLeft, lh)
		maxRight = max(maxRight, rh)

		left++
		right--
	}

	for axis := 1; axis < len(height)-1; axis++ {
		sideHeight := min(prefix[axis], suffix[axis])
		water := sideHeight - height[axis]
		if water > 0 {
			ans += water
		}
	}

	return ans
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = input height length

Time: O(2*n) = O(n) -> loop 2 times one for collect prefix & suffix and one for calculate result

Space: O(2n) + O(1) -> using 2 more array to store prefix suffix and some pointer. 
```

<br>
<br>

## A better approach: Time O(n), Space O(1) - Powered by AI

```go
func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }

    left, right := 0, len(height)-1
    maxLeft, maxRight := 0, 0
    ans := 0

    for left < right {
        // The core logic: move the pointer with the smaller height 
        // because that side is the "bottleneck" for trapping water.
        if height[left] < height[right] {
            if height[left] >= maxLeft {
                // We found a new peak, no water can be trapped here
                maxLeft = height[left]
            } else {
                // This spot is lower than the left peak, 
                // and we know the right side is even taller, so water is trapped!
                ans += maxLeft - height[left]
            }
            left++
        } else {
            if height[right] >= maxRight {
                // New peak on the right
                maxRight = height[right]
            } else {
                // Right side is lower than its peak, 
                // and we know the left side is taller, so water is trapped!
                ans += maxRight - height[right]
            }
            right--
        }
    }

    return ans
}
```