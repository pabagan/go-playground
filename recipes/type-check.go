package recipes

import (
	"fmt"
	"reflect"
)

// Every method has a different best use case:
//   string formatting - short and low footprint (not necessary to import reflect package)
//   reflect package - when need more details about the type we have access to the full reflection capabilities
//   type assertions - allows grouping types, for example recognize all int32, int64, uint32, uint64 types as "int"
func IsInt(t interface{}) bool {
	_, ok := t.(int) // assert that the actual type is int

	if ok {
		return true
	}

	return false
}

func typeofInterface(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// Using reflect package
func typeofReject(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// Using type assertions
func typeofSwitch(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	//... etc
	default:
		return "unknown"
	}
}

func TestTypeAssertions() {
	i := 1
	b := false
	c := "I'm a string"

	fmt.Printf("Is integer value \"%t\"? %t\n", b, IsInt(b))
	fmt.Printf("Is integer value \"%d\"? %t\n", i, IsInt(i))
	fmt.Printf("Is integer value \"%s\"? %t\n", c, IsInt(b))
}
