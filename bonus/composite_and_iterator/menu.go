package composite_and_iterator

import "fmt"

type Menu struct {
	menuComponents []MenuComponent
	name           string
	description    string
}

func (m *Menu) createIterator() Iterator {
	return NewCompositeIterator(NewMenuComponentIterator(m.menuComponents))
}

func (m *Menu) add(component MenuComponent) {
	m.menuComponents = append(m.menuComponents, component)
}

func (m *Menu) remove(component MenuComponent) {
	for i, c := range m.menuComponents {
		if c.getName() == component.getName() {
			m.menuComponents = append(m.menuComponents[:i], m.menuComponents[i+1:]...)
			return
		}
	}
}

func (m *Menu) getChild(idx int) (MenuComponent, error) {
	if idx < 0 || idx >= len(m.menuComponents) {
		return nil, fmt.Errorf("child index out of range")
	}
	return m.menuComponents[idx], nil
}

func (m *Menu) getName() string {
	return m.name
}

func (m *Menu) getDesc() string {
	return m.description
}

func (m *Menu) getPrice() float32 {
	panic("unsupported operation")
}

func (m *Menu) isVegetarian() bool {
	return false
}

func (m *Menu) print() {
	fmt.Printf("\n Menu: <%s>, %s\n", m.name, m.description)
	fmt.Println("-------------------------------------------")
	for _, item := range m.menuComponents {
		item.print()
	}
}

func NewMenu(name string, description string) MenuComponent {
	return &Menu{
		name:           name,
		description:    description,
		menuComponents: make([]MenuComponent, 0),
	}
}
