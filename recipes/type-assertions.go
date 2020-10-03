package recipes

import (
	"fmt"
)

/
func IsInt(t interface{}) bool {
	_, ok := t.(int) // assert that the actual type is int

	if ok {
		return true
	}

	return false
}

func TestTypeAssertions() {
	i := 1
	b := false
	c := "I'm a string"

	fmt.Printf("Is integer value \"%t\"? %t\n", b, IsInt(b))
	fmt.Printf("Is integer value \"%d\"? %t\n", i, IsInt(i))
	fmt.Printf("Is integer value \"%s\"? %t\n", c, IsInt(b))
}
