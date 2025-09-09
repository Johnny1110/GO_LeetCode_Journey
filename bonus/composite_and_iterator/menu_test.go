package composite_and_iterator

import "testing"

func Test_NewMenu(t *testing.T) {
	menu := NewMenu("King's steak", "Affordable steakhouse")

	item1 := NewMenuItem("7 oz ribeye steak", "from AU", false, 10.99)
	item2 := NewMenuItem("7 oz filet mignon", "from AU", false, 13.99)
	item3 := NewMenuItem("7 oz beef shoulder steak", "from US", false, 9.99)
	item4 := NewMenuItem("7 oz Marbled steak", "from AU", false, 8.99)

	menu.add(item1)
	menu.add(item2)
	menu.add(item3)
	menu.add(item4)

	menu.print()
}
