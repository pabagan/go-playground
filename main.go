package main

import (
	"fmt"

	"github.com/pabagan/go-playground/recipes"
)

func main() {
	// recipes.TestLogs()
	// recipes.TestTimes()
	// recipes.TestGenerics()
	// recipes.TestSeparateHouseNumber()
	// recipes.TestTypeAssertions()
	// recipes.TestChannelRequestLevelTask()
	// recipes.TestChannelServerLevelTask()
	randomString := recipes.RandomString(4)

	fmt.Println("Random string", randomString)
}
