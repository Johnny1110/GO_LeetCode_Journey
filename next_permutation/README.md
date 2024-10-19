# Next Permutation

<br>

---

<br>

## Desc

A permutation of an array of integers is an arrangement of its members into a sequence or linear order.

For example, for arr = `[1,2,3]`, the following are all the permutations of arr: `[1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1]`.
The next permutation of an array of integers is the next lexicographically greater permutation of its integer. More formally, if all the permutations of the array are sorted in one container according to their lexicographical order,
then the next permutation of that array is the permutation that follows it in the sorted container. If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).

For example, the next permutation of arr = `[1,2,3]` is `[1,3,2]`.
Similarly, the next permutation of arr = `[2,3,1]` is `[3,1,2]`.
While the next permutation of arr = `[3,2,1]` is `[1,2,3]` because `[3,2,1]` does not have a lexicographical larger rearrangement.
Given an array of integers nums, find the next permutation of nums.

The replacement must be in place and use only constant extra memory.

<br>
<br>

Example 1:

```
Input: nums = [1,2,3]
Output: [1,3,2]
```

<br>

Example 2:

```
Input: nums = [3,2,1]
Output: [1,2,3]
```

<br>

Example 3:

```
Input: nums = [1,1,5]
Output: [1,5,1]
```

<br>

Constraints:

```
1 <= nums.length <= 100
0 <= nums[i] <= 100
```

<br>
<br>

## Topic

* Array
* Two Pointers

<br>
<br>

## Thinking

<br>

I was thinking of using two pointers, pointer A and pointer B, both starting from the end of the integer array. Pointer B moves to the left, looking for the first element smaller than pointer A, and then swaps them.

After the swap, move pointer B one step to the right, and then sort the elements from pointer B to pointer A in ascending order.

![1.png](imgs/1.png)

<br>

And ChatGPT give me a comment like below: 

```
Yes, your solution is both correct and efficient! 
The approach you're using is essentially the standard algorithm for solving the Next Permutation problem.

Here’s why it’s optimal:

Steps of the Algorithm:

1. Find the pivot: Start from the end of the array (using your pointers A and B) and find the first element that is smaller than its next element. 
    This is your pivot point (pointer A).

2. Find the element to swap: Move pointer B leftwards to find the first element that is larger than the pivot (pointer A). Swap these two elements.

3. Reverse the suffix: After the swap, reverse the portion of the array to the right of pointer A to get the next lexicographically smallest order.


Time Complexity:
Finding the pivot and finding the element to swap each take O(n) in the worst case.
Reversing the suffix (sorting in ascending order) also takes O(n).
So overall, the time complexity is O(n), where n is the length of the array, which is optimal for this problem.

Conclusion:
Your approach is a standard, efficient solution to the Next Permutation problem, and it works perfectly! Great job!
```

<br>

After I tried, I still can not make it. my code is like:


```golang
func nextPermutation(nums []int) {
	maxIdx := len(nums) - 1
	pointerA := 0
	pointerB := 0

	// find first element who is greater than last one and swap them.
	for i := maxIdx; i > 0; i-- {
		pointerA = i
		pointerB = i - 1

		for pointerB >= 0 {
			if nums[pointerA] > nums[pointerB] {
				swap(nums, pointerA, pointerB)

				//printNums(nums)

				// move 1 step toward to right side and sort
				pointerB++
				sort(nums, pointerB, maxIdx)
				return
			}
			pointerB--
		}
	}

	// if reaching this point, that means input is last max permutation. like 3, 2, 1
	sort(nums, 0, maxIdx)

}

func printNums(nums []int) {
	for i := 0; i < len(nums); i++ {
		print(nums[i])
	}
	println()
}

func sort(nums []int, i, j int) {
	for ; i < j; i++ {
		for x := i + 1; x <= j; x++ {
			if nums[i] > nums[x] {
				swap(nums, x, i)
			}
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
```

<br>

And I asked ChatGPT to give me some hints:

* Step 1: Find the first decreasing element from the end

* Step 2: Find the element just larger than nums[i] from the right

* Step 3: Swap nums[i] and nums[j]

* Step 4: Reverse the subarray from i+1 to the end to get the smallest lexicographical order

<br>

### For example: `[5, 4, 7, 5, 3, 2]`

To find the next permutation, the first step is identifying which digits need to be swapped.

In this case, it's `4`. When looking at `7, 5, 3, 2`, there’s no way to find a permutation larger than this sequence because it's in descending order. 
But when we consider `4, 7, 5, 3, 2`, we can swap to find a new, larger permutation.

Now that we’ve identified `4`, the next step is to find a number in `7, 5, 3, 2` that is as small as possible but still larger than `4`.

we can find `5`

so now, we locked onto `4` and `5`. and we swap them. After swap we got a  new nums is like:

`5, 7, 4, 3, 2`

At this point, the subarray `7, 4, 3, 2` is sorted in descending order.

Next, we reverse the subarray `7, 4, 3, 2` to sort it in ascending order.

we got `5, 2, 3, 4, 7` after reversed.

Thus, the final answer is `5, 5, 2, 3, 4, 7`.

<br>

Let's try this on a new coding area.

<br>

```go
func nextPermutation(nums []int) {

	endIdx := len(nums) - 1

	// Step 1. identify the num need to swapped. (pointerA)
	pointerA := -1

	for i := endIdx; i > 0; i-- {
		if nums[i] > nums[i-1] {
			pointerA = i - 1
			break
		}
	}

	// if can not find the pointer, that means input nums is biggest permutation, reverse it and return.
	if pointerA == -1 {
		reverse(nums, 0, endIdx)
		return
	}

	// Step 2. find the element just larger then nums[pointerA] from the right side.
	pointerB := 0

	for i := endIdx; i > pointerA; i-- {
		if nums[i] > nums[pointerA] {
			pointerB = i
			break
		}
	}

	// step 3. swap nums[pointerA] and pointer[B]
	swap(nums, pointerA, pointerB)

	// step 4. reverse the sub array from pointer+1 to the end.
	reverse(nums, pointerA+1, endIdx)
}

func reverse(nums []int, a int, b int) {
	for a < b {
		swap(nums, a, b)
		a++
		b--
	}
}
```






