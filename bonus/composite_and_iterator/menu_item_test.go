package composite_and_iterator

import "testing"

func Test_New(t *testing.T) {
	m := NewMenuItem("Gold Soup", "A very tasty soup", true, 2.1)
	m.print()
	m = NewMenuItem("Super Rice", "A very tasty rice", true, 0.4)
	m.print()
}
