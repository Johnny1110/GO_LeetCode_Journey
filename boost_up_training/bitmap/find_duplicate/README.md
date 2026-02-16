# 287. Find the Duplicate Number (HARD as MDFK 這題太難了，直接死記硬背答案)

<br>

---

<br>

link: https://leetcode.com/problems/find-the-duplicate-number

## Floyd's algorithm

>> 用快慢針找出 linked-list 無窮迴圈的入口，做法是在快慢針相遇的 position，重置其中一個指針回到 0，然後兩個指針重新開始以每倫前進一步的步數移動，下次二者相遇的位置就是無窮迴圈的入口。

<br>
<br>

## Thinking

I have no clue about this, I just asking for some hint.

### Claude AI Hint:

Defining a function `f(i) = nums[i]`. 

Since values are in [1, n] and indices are [0, n], each value "points to" another index.

For example with `[1, 3, 4, 2, 2]`:

```
index 0 → value 1 → index 1
index 1 → value 3 → index 3
index 3 → value 2 → index 2
index 2 → value 4 → index 4
index 4 → value 2 → index 2  ← cycle!
```

<br>

### Inference

The problem is seems like a nature linked list with cycle.

Trace through the first index-0, it will lead us go to index() 

<br>

**But wait! what if value just equals to index?**

For example with value-2 sit on index-2  `[1, 4, 2, 3, 3]`:

```

... 
index 2 → value 2 → index 2
-> infinity loop on index-2
```

Then how do first timewe get in to index 2?

index? -> value 2 -> index 2

Is it possible? yes, but it will lead us to cicle result. I think it doesn't matter.

<br>

I did linked-list detection algo before. Let's adapt it in this problem.

<br>

`[3, 1, 3, 4, 2]`

```
trace: 3 -> 4 -> 2 -> 3 -> 4 -> 2 -> 3 -> 4......
```

Using **fast/slow pointer** we can find detect cicle in linked-list, but both pointer will meet in somewhere inside of cicle. not necessarily to be the duplicate.

But what we know is duplicate num must exists in cicle `[2, 3, 4]`.

<br>

### Floyd's algorithm 

F = distance from start to cycle entrance
a = distance from entrance to where slow and fast meet
C = total cycle length

<br>

What we know when they meet:

* Slow traveled: `F + a` steps
* Fast traveled: `F + a + (some full cycles) = F + a + kC`
* Fast moves twice as fast as slow, so:

```
2(F + a) = F + a + kC
F + a = kC
F = kC - a
```

#### The key insight:

```
F = kC - a
```

Starting from the **meeting point**, if you walk `F` more steps equals to you go `kC - a` steps. (F = kC- a)
 
That's the same as going backwards a steps in the cycle(then completing `k` full loops). 



* One pointer starts at 0, walks F steps → arrives at entrance
* Other pointer starts at meeting point, walks F steps → also arrives at entrance

**They meet at the entrance. And the entrance is your duplicate number.**

> 幹，這個真的超難，直接背公式吧：
>> Floyd's algorithm: 用快慢針找出 linked-list 無窮迴圈的入口，做法是在快慢針相遇的 position，重置其中一個指針回到 0，然後兩個指針重新開始以每倫前進一步的步數移動，下次二者相遇的位置就是無窮迴圈的入口。


<br>
<br>

## Coding


```go

func findDuplicate(nums []int) int {
	slowP := 0
	fastP := 0

	// phase-1: find cicle, and locate 2 pointers at same position.
	for {
		slowP = nums[slowP]
		fastP = nums[nums[fastP]]

		if slowP == fastP {
			break
		}
	}

	// phase-2: locate cicle entrance
	fastP = 0 // reset fastP to 0
	for {
		if fastP == slowP {
			return fastP
		}

		// both move 1 step
		slowP = nums[slowP]
		fastP = nums[fastP]
	}
}
```

<br>

Result:

![1](imgs/1.png)


<br>

## Other Approach

### Binary Search on value range 

This is a clever one. 

Instead of searching the array, you binary search on the answer space `[1, n]`. 

For a mid value, count how many numbers in the array are ≤ mid. If the count exceeds mid, the duplicate must be in `[1, mid]`. Otherwise it's in `[mid+1, n]`. This is `O(n log n)` time, `O(1)` space.

The key insight: **in a perfect array with no duplicates, exactly mid numbers would be `≤ mid`. The duplicate pushes that count higher on one side.**

<br>

### Bit manipulation

For each bit position, count how many times that bit is set across all values in nums, and compare with how many times it's set across `[1, n]`. 

If the count in nums is higher, that bit is set in the duplicate. Build the answer bit by bit. `O(n log n)` time, `O(1)` space.