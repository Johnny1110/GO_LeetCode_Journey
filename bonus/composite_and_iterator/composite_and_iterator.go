package composite_and_iterator

type MenuComponentIterator struct {
	mcs      []MenuComponent
	position int
}

func (m *MenuComponentIterator) hasNext() bool {
	return m.position < len(m.mcs)
}

func (m *MenuComponentIterator) next() interface{} {
	mc := m.mcs[m.position]
	m.position++
	return mc
}

func (m *MenuComponentIterator) remove() {
	panic("implement me")
}

func NewMenuComponentIterator(mcs []MenuComponent) Iterator {
	return &MenuComponentIterator{
		mcs:      mcs,
		position: 0,
	}
}

// ===============================================

type IteratorStack struct {
	iterators []Iterator
}

func (is *IteratorStack) pop() Iterator {
	if len(is.iterators) == 0 {
		return nil
	}
	iterator := is.iterators[len(is.iterators)-1]
	is.iterators = is.iterators[:len(is.iterators)-1]
	return iterator
}

func (is *IteratorStack) push(iterator Iterator) {
	is.iterators = append(is.iterators, iterator)
}

func (is *IteratorStack) peek() Iterator {
	if len(is.iterators) == 0 {
		return nil
	}
	iterator := is.iterators[len(is.iterators)-1]
	return iterator
}

func (is *IteratorStack) isEmpty() bool {
	return len(is.iterators) == 0
}

func NewIteratorStack() *IteratorStack {
	return &IteratorStack{
		iterators: make([]Iterator, 0),
	}
}

type CompositeIterator struct {
	iteratorStack *IteratorStack
}

func (c *CompositeIterator) hasNext() bool {
	if c.iteratorStack.isEmpty() {
		return false
	}

	// 彈出 1 個 iterator
	it := c.iteratorStack.peek()

	if !it.hasNext() { // 如果 iterator 空了
		c.iteratorStack.pop() // 移除當前 iterator
		return c.hasNext()
	} else {
		return true
	}
}

func (c *CompositeIterator) next() interface{} {
	if c.hasNext() {
		// 取出一個 ITERATOR
		it := c.iteratorStack.peek()
		// 從該 ITERATOR 取出一個 MenuComponent
		next := it.next().(MenuComponent)
		childIt := next.createIterator()
		if childIt != nil {
			c.iteratorStack.push(childIt)
		}
		return next
	}
	return nil
}

func (c *CompositeIterator) remove() {
	panic("implement me")
}

func NewCompositeIterator(iterator Iterator) Iterator {
	c := &CompositeIterator{
		iteratorStack: NewIteratorStack(),
	}

	c.iteratorStack.push(iterator)
	return c
}
