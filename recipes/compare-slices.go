package recipes

import (
	"bytes"
	"fmt"
	"reflect"
)

func TestCompareSlices() {
	a := []byte{
		'a',
		'b',
	}
	b := []byte{
		'b',
	}

	fmt.Println("Check using bytes.Equal:", bytes.Equal(a, b))
	fmt.Println("Check using reflect.DeepEqual:", reflect.DeepEqual(a, b))
}
