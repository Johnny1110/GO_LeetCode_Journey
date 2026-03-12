# 146. LRU Cache

<br>

---

<br>

```go
type LRUCache struct {
	capacity int
	size     int
	head     *Node
	tail     *Node
	data     map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		size:     0,
		head:     head,
		tail:     tail,
		data:     make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.data[key]; exists {

		// update node list
		if this.head.next != node {
			// remove target node from node lists.
			prev, next := node.prev, node.next
			prev.next = next
			next.prev = prev

			// move to front
			tmp := this.head.next
			this.head.next = node
			node.prev = this.head
			tmp.prev = node
			node.next = tmp
		}

		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.data[key]; exists {
		// updaate Nodes
		node.val = value
		prev, next := node.prev, node.next
		// remove target node from node lists.
		prev.next = next
		next.prev = prev

		// move to front
		tmp := this.head.next
		this.head.next = node
		node.prev = this.head
		tmp.prev = node
		node.next = tmp

		this.data[key] = node
	} else {
		// insert Node

		node := &Node{
			key: key,
			val: value,
		}

		// add new node to front
		tmp := this.head.next
		this.head.next = node
		node.prev = this.head
		tmp.prev = node
		node.next = tmp

		this.data[key] = node
		this.size++

		if this.size > this.capacity { // remove expired node
			// remove last one
			deleteNode := this.tail.prev
			prev := deleteNode.prev
			prev.next = this.tail
			this.tail.prev = prev
			// delete data
			delete(this.data, deleteNode.key)
			this.size--
		}
	}

}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}
```
