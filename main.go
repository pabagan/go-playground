package main

import (
	"log"

	"github.com/pabagan/go-playground/recipes"
)

func main() {
	recipes.Log("hola")
	recipes.Log("hola")
	time := recipes.getTime()

	log.Println(time)
}
