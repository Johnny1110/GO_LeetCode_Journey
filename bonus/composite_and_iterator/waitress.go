package composite_and_iterator

import "fmt"

type Waitress struct {
	allMenus MenuComponent
}

func NewWaitress() *Waitress {
	return &Waitress{
		allMenus: NewMenu("Super Waitress", "I'm a waitress."),
	}
}

func (w *Waitress) AddMenu(m MenuComponent) {
	w.allMenus.add(m)
}

func (w *Waitress) PrintVegetarianMenu() {
	it := w.allMenus.createIterator()
	fmt.Println("\n 🥦 Vegetarian Menu:")
	for it.hasNext() {
		m := it.next()

		switch v := m.(type) {
		case *MenuItem:
			if v.isVegetarian() {
				v.print()
			}
		default:
			continue
		}
	}
}

func (w *Waitress) PrintMenus() {
	w.allMenus.print()
}
