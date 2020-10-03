package recipes

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemCapsule struct {
	s []Item
}

func NewItemCapsule() *ItemCapsule {
	return &ItemCapsule{s: []Item{}}
}

func (c *ItemCapsule) Put(val Item) {
	c.s = append(c.s, val)
}

func (c *ItemCapsule) Get(element int) Item {
	if element >= len(c.s) {
		return nil
	}

	return c.s[element]
}

func (c *ItemCapsule) String() string {
	return fmt.Sprint(c.s)
}

func (c *ItemCapsule) Delete(element int) {
	outPut := NewItemCapsule()

	for i := 0; i < len(c.s); i++ {
		if i != element {
			outPut.s = append(outPut.s, c.s[i])
		}
	}

	c.s = outPut.s
}

func TestGenerics() {
	i := NewItemCapsule()
	i.Put(0)
	i.Put(1)
	i.Put("two")
	i.Put(3)
	i.Put(true)

	fmt.Println(i.Get(0))
	fmt.Println("Delete index 1")
	i.Delete(1)
	fmt.Println(i.Get(0))

	fmt.Println(i.String())
}
