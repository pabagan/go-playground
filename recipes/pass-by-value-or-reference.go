package recipes

import "fmt"

// int
func modifyInt(n int) int {
	return n + 5
}
func myMap(v map[int]int) {
	v = make(map[int]int) // make() declares and initializes v to 0
}

// func myInt(v []int) {
// 	v = make([]int) // make() declares and initializes v to 0
// }
type Profile struct {
	Age          int
	Name         string
	Salary       float32
	TechInterest bool
}

// # pointer
func modifyStruct(p Profile) Profile {
	p.Age = 85
	p.Name = "Balqees"
	p.Salary = 500.45
	p.TechInterest = true

	return p
}

func ModifyBasicTypes(
	name *string, age *int, cash *float64,
	techInterest *bool, countries *[3]string, myProfile *Profile,
) {
	*name = "Golang"
	*age = 90
	*cash = 50.45
	*techInterest = !(*techInterest)
	*countries = [3]string{"sudanese", "belgium", "zambia"}
	*myProfile = Profile{
		Age:          100,
		Name:         "GOOGLE",
		Salary:       40.54,
		TechInterest: true,
	}
}
func TestPassByValue() {
	age := 30
	fmt.Println("Before function call: ", age) // 30 fmt.Println("Function call:", modifyInt(age))           // 35
	fmt.Println("After function call: ", age)  // 30

	myProfile := Profile{
		Age:          15,
		Name:         "Adeshina",
		Salary:       300,
		TechInterest: false,
	}
	fmt.Println("Before function call: ", myProfile)
	// {15 Adeshina 300 false}
	fmt.Println("Function call:", modifyStruct(myProfile))
	// {85 Balqees 500.45 true}
	fmt.Println("After function call: ", myProfile)
	// {15 Adeshina 300 false}

}
