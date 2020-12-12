package recipes

import (
	"flag"
	"fmt"
	"strconv"
)

// go build cli-app-with-flags.go
// go run cli-app-with-flags.go --help
func TestAppCLIFlags() {
	var message = flag.String(
		"message",
		"Hello, World!",
		"The message you'd like to print to the terminal",
	)

	var number = flag.Int(
		"number",
		1,
		"The number you'd like to add to your message",
	)

	flag.Parse()

	fmt.Println("This is the message you want to display: " + *message + " with number " + strconv.Itoa(*number))
}
