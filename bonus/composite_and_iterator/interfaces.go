package composite_and_iterator

type MenuComponent interface {
	add(component MenuComponent)
	remove(component MenuComponent)
	getChild(idx int) (MenuComponent, error)
	getName() string
	getDesc() string
	getPrice() float32
	isVegetarian() bool
	print()

	createIterator() Iterator
}

type Iterator interface {
	hasNext() bool
	next() interface{}
	remove()
}
