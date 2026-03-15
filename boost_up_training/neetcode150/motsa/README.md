# 4. Median of Two Sorted Arrays

<br>

---

<br>

## Failed Try

```go

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenNum1, lenNum2 := len(nums1), len(nums2)
	mergedLen := (lenNum1 + lenNum2)
	leftLen := mergedLen / 2

	n1Cnt, n2Cnt := 0, 0
	for i := 0; i <= lenNum1 && i <= leftLen; i++ {

		n1Cnt, n2Cnt = i, leftLen-i

		if n2Cnt == 0 {
			// all num1
			n1LastPut := nums1[n1Cnt-1]
			n2FirstLeft := nums2[n2Cnt]
			if n2FirstLeft > n1LastPut {
				continue
			}
		} else {
			n2LastPut := nums2[n2Cnt-1] // 2
			n1FirstLeft := nums1[n1Cnt] // 1
			if n2LastPut > n1FirstLeft {
				continue
			}
		}

		break
	}

	if mergedLen%2 == 0 {
		// even

		leftLast, rightFirst := 0, 0
		if n2Cnt == 0 {
			leftLast = nums1[n1Cnt-1]
		} else {
			leftLast = nums2[n2Cnt-1]
		}

		if n1Cnt == len(nums1) {
			rightFirst = nums2[n2Cnt]
		} else if n2Cnt == len(nums2) {
			rightFirst = nums2[n2Cnt]
		} else {
			rightFirst = min(nums1[n1Cnt], nums2[n2Cnt])
		}

		return float64(leftLast+rightFirst) / 2

	} else {
		// odd
		rightFirst := 0
		if n1Cnt == len(nums1) {
			rightFirst = nums2[n2Cnt]
		} else if n2Cnt == len(nums2) {
			rightFirst = nums2[n2Cnt]
		} else {
			rightFirst = min(nums1[n1Cnt], nums2[n2Cnt])
		}
		return float64(rightFirst)
	}

}
```

<br>
<br>

## The Mental Model

Imagine cutting both arrays with a knife:

```
A:  [1, 3, | 5, 7]    ← i = 2 (took 2 from A)
B:  [2, | 4, 6]        ← j = 1 (took 1 from B)

Left half:  {1, 3, 2}  → max is 3
Right half: {5, 7, 4, 6} → min is 4
```

The "cut position" i means: A[0..i-1] goes left, A[i..] goes right. So A[i-1] is the last element on the left from A, and A[i] is the first element on the right from A. Same for B with j.

The tricky part is when i=0 (took nothing from A) or i=len(A) (took everything). In those cases, A[i-1] or A[i] doesn't exist. We handle that with -∞ and +∞ sentinels — because "I took nothing from A" means A contributes nothing to the left, so its left-max should be as small as possible (won't affect the real max).


<br>

```go
import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1Len, nums2Len := len(nums1), len(nums2)
	if nums1Len > nums2Len {
		return findMedianSortedArrays(nums2, nums1)
	}

	mergedLen := nums1Len + nums2Len
	leftHalfLen := (mergedLen + 1) / 2
	// + 1 -> simplify odd sitiation = max val in LefHalf

	lo, hi := 0, nums1Len
	i, j := 0, 0 // i is nums1 take, j is nums2 take.

	for lo <= hi {

		// assign i, j
		i = (lo + hi) / 2
		j = leftHalfLen - i

		// boundary
		aLeft := math.MinInt64
		aRight := math.MaxInt64
		bLeft := math.MinInt64
		bRight := math.MaxInt64

		if i > 0 {
			aLeft = nums1[i-1]
		}
		if i < nums1Len {
			aRight = nums1[i]
		}

		if j > 0 {
			bLeft = nums2[j-1]
		}
		if j < nums2Len {
			bRight = nums2[j]
		}

		// check validate or not
		if aLeft <= bRight && bLeft <= aRight { // found
			if mergedLen%2 == 0 {
				// even: max left + min right
				return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2.0
			} else {
				// odd: max left
				return float64(max(aLeft, bLeft))
			}
		}

		// not found -> adjust lo and hi
		if aLeft > bRight {
			// we need cut down a left count
			hi--
		} else if bLeft > aRight {
			// we need add more a left count
			lo++
		}

	}

	return 0
}
```