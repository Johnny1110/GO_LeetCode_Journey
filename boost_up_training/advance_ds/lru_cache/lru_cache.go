package lru_cache

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
