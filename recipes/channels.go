package recipes

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func Channel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

type Container []interface{}

func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

func (c *Container) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem
}

func AssertExample() {
	intContainer := &Container{}
	intContainer.Put(7)
	intContainer.Put(42)
	elem, ok := intContainer.Get().(int) // assert that the actual type is int
	if !ok {
		fmt.Println("Unable to read an int from intContainer")
	}
	fmt.Printf("assertExample: %d (%T)\n", elem, elem)
}
