package composite_and_iterator

type Waitress struct {
	allMenus []MenuComponent
}

func NewWaitress() *Waitress {
	return &Waitress{
		allMenus: make([]MenuComponent, 0),
	}
}

func (w *Waitress) AddMenu(m MenuComponent) {
	w.allMenus = append(w.allMenus, m)
}

func (w *Waitress) PrintVegetarianMenu() {
	// TODO: implement iterator
	panic("implement me")
}

func (w *Waitress) PrintMenus() {
	for _, m := range w.allMenus {
		m.print()
	}
}
