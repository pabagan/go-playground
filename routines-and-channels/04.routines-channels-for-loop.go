package routinesAndChannels

import (
	"fmt"
	"time"
)

func TestRoutinesChannels() {
	c := make(chan string, 2)

	go countChannel("sheep", c)

	// Full operation
	// for {
	// 	msg, open := <-c

	// 	if !open {
	// 		break
	// 	}

	// 	fmt.Println(msg)
	// }

	// Or syntactic sugar
	for msg := range c {
		fmt.Println(msg)
	}

}

func countChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
