# 56. Merge Intervals

<br>

---

<br>

## Coding

```go
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	result := [][]int{}

	// 1 order intervals (order by first number)
	for i := 0; i < len(intervals); i++ {
		for j := i; j < len(intervals); j++ {
			intA, intB := intervals[i][0], intervals[j][0]
			//fmt.Printf("i:%v, j:%v,  intA: %v, intB:%v, swap=%v \n", i, j, intA, intB, intA > intB)
			if intA > intB {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	//fmt.Printf("intervals: %v\n", intervals)
	tmp := intervals[0]
	for idx := 1; idx < len(intervals); idx++ {
		if canMerge(tmp, intervals[idx]) {
			tmp[1] = max(tmp[1], intervals[idx][1])
		} else {
			// append result
			result = append(result, []int{tmp[0], tmp[1]})
			tmp = intervals[idx]
		}
	}

	result = append(result, []int{tmp[0], tmp[1]})

	return result
}

func canMerge(intsA, intsB []int) bool {
	if intsB[0] >= intsA[0] && intsB[0] <= intsA[1] {
		return true
	}
	return false
}
```

<br>
<br>

## Space & Time Complexity

```
Assume: n = input length of intervals

Time: O(n square) -> order take n square times, merge take n times.

Space: O(1) -> except result slice, I only use a tmp(2 len slice).
```

<br>
<br>

## Revamp O(n log N) sort

```go
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort the intervals by first int element time: O(n log2 n), space: O(log n)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// merge O(n)
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]

		if last[1] >= intervals[i][0] { // can merge
			// merge
			last[1] = max(last[1], intervals[i][1])
		} else {
			// add new result
			result = append(result, intervals[i])
		}
	}

	return result
}
```

```
Assume: n = input length of intervals

Time: O(N log N) -> order take N * log2 N, merge take n times.

Space: O(log N) -> result not count in result slice, sort.Slice uses an introsort variant which take O(log N)
```