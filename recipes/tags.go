package recipes

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `mytag:"MyName"`
	Email string `mytag:"MyEmail"`
}

func TestTags() {
	u := User{"Bob", "bob@mycompany.com"}
	t := reflect.TypeOf(u)

	fmt.Println("Type of u is", t)

	for _, fieldName := range []string{"Name", "Email"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}
		fmt.Printf("\nField: User.%s\n", fieldName)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
	}
}
