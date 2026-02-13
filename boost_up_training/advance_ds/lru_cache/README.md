# 146. LRU Cache

<br>

---

<br>

link: https://leetcode.com/problems/lru-cache/description/

<br>

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

* LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
* int get(int key) Return the value of the key if the key exists, otherwise return -1.
* void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
* The functions get and put must each run in O(1) average time complexity.

<br>

Example 1:

```
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]
```

Explanation

```
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
```

<br>

## Thinking

<br>

We need a kind of map data structure to implement LRUCache.

The critical thing we have to know is:

> "If the number of keys exceeds the capacity from this operation, evict the least recently used key."

I gonna need a priority queue, priority by recently access time. actually I think we don't need a field to record the timestamp

I want use a doubly linked list to implement this priority queue:

```
[DUMMY HEAD] -> [1] <-> [2] <-> [3] <-> ... <-> [n]
   head                                         end
```

we need a `map[int]*Node` to index all the Linked Node and also each node should have pre-node-ref and next-node-ref that made a doubly linked list.

the LRUCache needs `head node pointer` and `end node pointer`, which is a proper O(1) approach to move a latest accessed node from any position to the front, or remove the last (least recently used).

<br>

#### Get(key) → what two things need to happen?

1. find the key in `map[int]*Node`, if key not exists return -1 as not found
2. if find node by key, link `node.prev` -> `node.next`, move node to DUMMY Node' next 

<br>

#### Put(key, value) → what cases do you need to handle? (hint: there are three)

1. if key exist, update node val.
2. if key not exist, check current capacity is max or not
   * if capacity if full, move end pointer pointing to `end.prev`, and then `end.next = nil`
3. put this key to front of the linked list.

<br>
<br>

## Coding

```go
type LRUCache struct {
	index    map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
	capacity int
	size     int
}

func (this *LRUCache) Size() int {
	return this.size
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		index:    make(map[int]*LRUNode, capacity),
		head:     &LRUNode{}, // DUMMY HEAD
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.index[key]

	if exist {
		node.val = value
		this.removeNode(node)
		this.appendNode(node)
	} else {
		if this.capacity == 0 {
			return
		}
		node = newNode(key, value)
		if this.size == this.capacity {
			// delete last one
			delete(this.index, this.tail.key)
			this.removeNode(this.tail)
		}
		this.appendNode(node)
		this.index[key] = node
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.index[key]
	if !exist {
		return -1
	}

	this.removeNode(node)
	this.appendNode(node)
	return node.val
}

func (this *LRUCache) removeNode(node *LRUNode) {
	prev := node.prev
	next := node.next

	prev.next = next
	if next != nil {
		next.prev = prev
	}

	node.prev = nil
	node.next = nil

	this.size--

	// tail check
	if this.tail == node {
		if prev != this.head {
			this.tail = prev
		} else {
			this.tail = nil
		}
	}
}

func (this *LRUCache) appendNode(node *LRUNode) {
	head := this.head

	if head.next == nil {
		head.next = node
		node.prev = head

		this.tail = node
	} else {
		tmp := head.next

		head.next = node
		node.prev = head

		node.next = tmp
		tmp.prev = node
	}

	this.size++
}

type LRUNode struct {
	key  int
	val  int
	prev *LRUNode
	next *LRUNode
}

func newNode(key, val int) *LRUNode {
	return &LRUNode{
		key: key,
		val: val,
	}
}
```

Result:

![1.png](imgs/1.png)


<br>
<br>

## Coding - Enhancement

Using Dummy Tail also like Dummy Head:

```
dummy_head <-> [real nodes] <-> dummy_tail
```

<br>

```go
type LRUCache struct {
	index    map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
	capacity int
	size     int
}

func (this *LRUCache) Size() int {
	return this.size
}

func Constructor(capacity int) LRUCache {
	dummyHead := newNode(0, 0)
	dummyTail := newNode(0, 0)

	dummyHead.next = dummyTail
	dummyTail.prev = dummyHead

	return LRUCache{
		index:    make(map[int]*LRUNode, capacity),
		head:     dummyHead, // DUMMY HEAD
		tail:     dummyTail, // DUMMY TAIL
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.index[key]

	if !exist {
		node = newNode(key, value)
		if this.size == this.capacity {
			// delete last one
			lastOne := this.tail.prev
			delete(this.index, lastOne.key)
			this.removeNode(lastOne)
		} else {
			this.size++
		}
	} else {
		node.val = value
		this.removeNode(node)

	}

	this.appendNode(node)
	this.index[key] = node
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.index[key]
	if !exist {
		return -1
	}

	this.removeNode(node)
	this.appendNode(node)
	return node.val
}

func (this *LRUCache) removeNode(node *LRUNode) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node.prev = nil
	node.next = nil
}

func (this *LRUCache) appendNode(node *LRUNode) {
	head := this.head

	tmp := head.next

	head.next = node
	node.prev = head

	node.next = tmp
	tmp.prev = node
}

// ==============================================================

type LRUNode struct {
	key  int
	val  int
	prev *LRUNode
	next *LRUNode
}

func newNode(key, val int) *LRUNode {
	return &LRUNode{
		key: key,
		val: val,
	}
}
```

Result:

![2.png](imgs/2.png)