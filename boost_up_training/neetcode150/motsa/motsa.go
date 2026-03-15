package motsa

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
