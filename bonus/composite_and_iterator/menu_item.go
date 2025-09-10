package composite_and_iterator

import "fmt"

type MenuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float32
}

func (m *MenuItem) createIterator() Iterator {
	return NewMenuComponentIterator([]MenuComponent{})
}

func (m *MenuItem) add(component MenuComponent) {
	panic("unsupported operation")
}

func (m *MenuItem) remove(component MenuComponent) {
	panic("unsupported operation")
}

func (m *MenuItem) getChild(idx int) (MenuComponent, error) {
	panic("unsupported operation")
}

func (m *MenuItem) getName() string {
	return m.name
}

func (m *MenuItem) getDesc() string {
	return m.description
}

func (m *MenuItem) getPrice() float32 {
	return m.price
}

func (m *MenuItem) isVegetarian() bool {
	return m.vegetarian
}

func (m *MenuItem) print() {
	mark := ""
	if m.vegetarian {
		mark = "(v)"
	}
	fmt.Printf("  %s%s, $ %.2f \n", m.name, mark, m.price)
	fmt.Printf("    -- Desc: %s \n", m.description)
}

func NewMenuItem(name string, description string, vegetarian bool, price float32) MenuComponent {
	return &MenuItem{
		name:        name,
		description: description,
		vegetarian:  vegetarian,
		price:       price,
	}
}
