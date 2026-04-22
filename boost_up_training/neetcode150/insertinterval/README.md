# 57. Insert Interval

<br>

---

<br>

## Coding

```go

type State int

const (
	MERGE State = iota
	IDLE
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := [][]int{}

	// edge case: new_interval is lower than 1st & can't merge
	if !canMerge(intervals[0], newInterval) && newInterval[0] < intervals[0][0] {
		result = append(result, newInterval)
		result = append(result, intervals...)
		return result
	}

	state := IDLE
	for i := 0; i < len(intervals); i++ {

		currInterval := intervals[i]

		if state == IDLE {
			// newInterval is not merged yet.
			if canMerge(newInterval, currInterval) {
				// first merged
				newInterval = merge(newInterval, currInterval)
				result = append(result, newInterval)
				state = MERGE
			} else if currInterval[0] < newInterval[0] {
				result = append(result, currInterval)
			} else {
				result = append(result, newInterval)
				result = append(result, currInterval)
				state = MERGE
			}

		} else { // state = MERGE (newInterval already merged in last results)
			if canMerge(newInterval, currInterval) {
				newInterval = merge(newInterval, currInterval)
				result[len(result)-1] = newInterval
			} else {
				result = append(result, currInterval)
			}
		}
	}

	// edge case: new_interval is bigger than last & can't merge
	if state == IDLE {
		result = append(result, newInterval)
	}

	return result
}

func canMerge(a, b []int) bool {
	if a[1] < b[0] || b[1] < a[0] {
		return false
	}
	return true
}

func merge(a, b []int) []int {
	a[0] = min(a[0], b[0])
	a[1] = max(a[1], b[1])
	return a
}
```

## Time & Space Complexity

```
Assume: n = length of intervals

Time: O(n)

Space: O(1) -> only auxiliary space: no extra space usage
```

<br>
<br>

## Clean Code Revamp

```go
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := [][]int{}

	i, n := 0, len(intervals)

	// 1. intervals entirely before newInterval
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 2. intervals that overlap with newInterval - do merge
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	// 3. intervals entirely after newInterval
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
```