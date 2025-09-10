package composite_and_iterator

import "testing"

func Test_NewWaitress(t *testing.T) {
	waitress := NewWaitress()

	steakMenu := createSteakMenu()
	waitress.AddMenu(steakMenu)

	vegeMenu := createVegetarianMenu()
	waitress.AddMenu(vegeMenu)

	waitress.PrintMenus()
}

func TestWaitress_PrintVegetarianMenu(t *testing.T) {
	waitress := NewWaitress()
	steakMenu := createSteakMenu()
	waitress.AddMenu(steakMenu)

	vegeMenu := createVegetarianMenu()
	waitress.AddMenu(vegeMenu)

	waitress.PrintVegetarianMenu()
}

func createVegetarianMenu() MenuComponent {
	menu1 := NewMenu("Queen's Vegetarian", "Expensive Vegetarian")

	item1 := NewMenuItem("Gold Soup", "Very tasty soup", true, 21.99)
	item2 := NewMenuItem("Veggie burger", "Very tasty burger", true, 50)
	item3 := NewMenuItem("Veggie noddle", "Very tasty noddle", true, 15.99)

	menu1.add(item1)
	menu1.add(item2)
	menu1.add(item3)

	dessertMenu := NewMenu("Queen's Vegetarian Dessert", "Expensive Vegetarian Dessert")
	dessertItem1 := NewMenuItem("Defi Chocolate Cake", "Very tasty Cake", true, 9.99)
	dessertItem2 := NewMenuItem("Faker vanilla Cake", "Very tasty Cake", true, 9.99)
	dessertMenu.add(dessertItem1)
	dessertMenu.add(dessertItem2)
	menu1.add(dessertMenu)

	return menu1
}

func createSteakMenu() MenuComponent {
	menu1 := NewMenu("King's steak", "Affordable steakhouse")

	item1 := NewMenuItem("7 oz ribeye steak", "from AU", false, 10.99)
	item2 := NewMenuItem("7 oz filet mignon", "from AU", false, 13.99)
	item3 := NewMenuItem("7 oz beef shoulder steak", "from US", false, 9.99)
	item4 := NewMenuItem("7 oz Marbled steak", "from AU", false, 8.99)
	item5 := NewMenuItem("5 oz Vegetarian steak", "from TW", true, 7.99)

	menu1.add(item1)
	menu1.add(item2)
	menu1.add(item3)
	menu1.add(item4)
	menu1.add(item5)
	return menu1
}
