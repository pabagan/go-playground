package routinesAndChannels

import (
	"fmt"
	"time"
)

func TestRoutinesAndChannel() {
	fmt.Println("TestRoutinesAndChannel")
	go count("Paella")
	count("Cocido")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}
