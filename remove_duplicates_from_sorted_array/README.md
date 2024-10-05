# Remove Duplicates from Sorted Array

<br>

---

<br>

## Desc

Given an integer array nums sorted in non-decreasing order, remove the duplicates __[in-place](https://en.wikipedia.org/wiki/In-place_algorithm)__ such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
Custom Judge:

The judge will test your solution with the following code:
```
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
assert nums[i] == expectedNums[i];
}

```
If all assertions pass, then your solution will be accepted.

<br>
<br>

Example 1:
```
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
```
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

<br>

Example 2:
```
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
```
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

<br>

Constraints:
```
1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.
```

<br>
<br>

## Topic

1. Array
2. Two Pointers

<br>
<br>

## Thinking

First of all, Desc said I should remove the duplicates in-place, which is mean I can not utilize Set to solve this problem.
Topic said using Two Pointer, so I think it's should be like just using two pointer like:


```golang
package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	pointerA := 0

	for pointerB := 1; pointerB < len(nums); pointerB++ {
		if nums[pointerA] != nums[pointerB] {
			pointerA++
			nums[pointerA] = nums[pointerB]
		}
	}

	return pointerA + 1
}
```

