# 61. Rotate List

<br>

---

<br>

link: https://leetcode.com/problems/rotate-list/description/

<br>

## Thinking

I think the critical point is how do we locate the final tail node index with input `k`. 

* case node length = 5
* 
  * if `k` == 0 -> new tail is `node[4]`
  * if `k` == 1 -> new tail is `node[3]`
  * if `k` == 2 -> new tail is `node[2]`
  * if `k` == 3 -> new tail is `node[1]`
  * if `k` == 4 -> new tail is `node[0]`
  * if `k` == 5 -> new tail is `node[4]`
  * if `k` == 6 -> new tail is `node[3]`
  * if `k` == 7 -> new tail is `node[2]`
  * if `k` == 8 -> new tail is `node[1]`
  * if `k` == 9 -> new tail is `node[0]`

the `node[n]` formula is `n = length-1 % k` 

<br>
<br>

## Coding

```go
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newTailNode := head

	length := calculateLenAndClosedLoop(head)
	newTailIndex := locateTailNodeIdx(length, k)

	// loop to new tail and cut the next pointer
	for i := 1; i <= newTailIndex; i++ {
		newTailNode = newTailNode.Next
	}

	newHeadNode := newTailNode.Next
	newTailNode.Next = nil

	return newHeadNode
}

func calculateLenAndClosedLoop(head *ListNode) int {
	if head == nil {
		return 0
	}

	tmp := head
	count := 1
	for tmp.Next != nil {
		count++
		tmp = tmp.Next
	}

	// connect tail to head
	tmp.Next = head

	return count
}

func locateTailNodeIdx(nodeLen, k int) int {
	return (nodeLen - 1) - k%nodeLen
}
```

<br>
<br>

## Result

![1.png](imgs/1.png)
