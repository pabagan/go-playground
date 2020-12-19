package recipes

import (
	"fmt"
)

func square(number *float64) {
	*number = *number * *number
}

func main() {
	x := 1.5
	square(&x)
	fmt.Print(x == 2.25)
}
