package recipes

import (
	"fmt"
)

func MapHasProperty(t map[int]string, property int) bool {
	if _, ok := t[property]; ok {
		return true
	}

	return false
}

func TestTMapCheckProperty() {
	mapNumbers := map[int]string{
		5: "5th",
		1: "1st",
		4: "4th",
		3: "3rd",
		2: "2nd",
	}

	fmt.Println(MapHasProperty(mapNumbers, 5))
}
