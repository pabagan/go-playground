package routinesAndChannels

import (
	"fmt"
	"time"
)

func TestRoutinesChannelsReceiveMultiple() {
	c := make(chan string)

	go countChannelMultiple("sheep", c)

	msg := <-c

	fmt.Println(msg)
}

func countChannelMultiple(thing string, c chan string) {
	for i := 1; true; i++ {
		time.Sleep(time.Second * 2)
		c <- thing
	}
}
