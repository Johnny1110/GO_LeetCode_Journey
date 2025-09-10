package composite_and_iterator

import "testing"

func Test_MenuComponentIterator(t *testing.T) {
	m := NewMenuComponentIterator(mockMenuComponentSlice())

	for m.hasNext() {
		tar := m.next()
		menuComp := tar.(MenuComponent)
		menuComp.print()
	}
}

func mockMenuComponentSlice() []MenuComponent {

	item1 := NewMenuItem("7 oz ribeye steak", "from AU", false, 10.99)
	item2 := NewMenuItem("7 oz filet mignon", "from AU", false, 13.99)
	item3 := NewMenuItem("7 oz beef shoulder steak", "from US", false, 9.99)
	item4 := NewMenuItem("7 oz Marbled steak", "from AU", false, 8.99)

	mcs := make([]MenuComponent, 0)
	mcs = append(mcs, item1, item2, item3, item4)

	return mcs
}
