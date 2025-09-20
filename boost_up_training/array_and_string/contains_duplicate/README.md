# 217. Contains Duplicate

<br>

---

<br>

link: https://leetcode.com/problems/contains-duplicate/description/

<br>
<br>

## Thinking

That one is pretty ez, just put all element into a set. them compare the size between 
input nums and set. or using map instead.

<br>

## Coding-1

```go
func containsDuplicate(nums []int) bool {
	checkMap := map[int]bool{}

	for _, num := range nums {
		if _, exists := checkMap[num]; exists {
			return true
		} else {
			checkMap[num] = true
		}
	}

	return false
}
```

<br>

![1.png](imgs/1.png)

I'll take it, good.
